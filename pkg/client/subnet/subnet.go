package subnet

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/resource"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/network"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-cli/pkg/vars"
	"n2x.dev/x-lib/pkg/utils/msg"
)

var networkCIDR string
var subnetsMap map[string]*topology.Subnet = nil

func GetSubnet(edit bool) *topology.Subnet {
	sl := subnets()

	if len(sl) == 0 {
		msg.Info("No subnets found")
		os.Exit(1)
	}

	var subnetOptID string
	subnetsOpts := make([]string, 0)
	subnets := make(map[string]*topology.Subnet)

	for _, s := range sl {
		if s.IPAM != nil {
			subnetOptID = fmt.Sprintf("[%s] %s", s.SubnetID, s.Description)
		} else {
			subnetOptID = s.SubnetID
		}
		subnetsOpts = append(subnetsOpts, subnetOptID)
		subnets[subnetOptID] = s
	}

	sort.Strings(subnetsOpts)

	if edit {
		subnetsOpts = append(subnetsOpts, input.NewResource)
	}

	subnetOptID = input.GetSelect("Subnet:", "", subnetsOpts, survey.Required)

	if subnetOptID == input.NewResource {
		return nil
	}

	vars.SubnetID = subnets[subnetOptID].SubnetID

	return subnets[subnetOptID]
}

func subnets() map[string]*topology.Subnet {
	n := network.GetNetwork(false)

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	lr := &topology.ListSubnetsRequest{
		Meta: &resource.ListRequest{},
		Network: &topology.NetworkReq{
			AccountID: n.AccountID,
			TenantID:  n.TenantID,
			NetID:     n.NetID,
		},
	}

	subnets := make(map[string]*topology.Subnet) // map[subnetCIDR]*topology.Subnet

	for {
		sl, err := nxc.ListSubnets(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list subnets")
		}

		for _, s := range sl.Subnets {
			subnets[s.IPAM.SubnetCIDR] = s
		}

		if len(sl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = sl.Meta.NextPageToken
		} else {
			break
		}
	}

	return subnets
}

func validSubnet(val interface{}) error {
	_, ipv4Network, err := net.ParseCIDR(networkCIDR)
	if err != nil {
		return err
	}

	ipv4Addr, ipv4Subnet, err := net.ParseCIDR(val.(string))
	if err != nil {
		return err
	}

	if !ipv4Addr.Equal(ipv4Subnet.IP) {
		return fmt.Errorf("invalid subnet address %v", ipv4Addr)
	}

	if !ipv4Network.Contains(ipv4Subnet.IP) {
		return fmt.Errorf("subnet %s is not included in network %s", val.(string), networkCIDR)
	}

	cidrMask, _ := ipv4Subnet.Mask.Size()

	if cidrMask != 24 {
		return errors.New("only /24 subnets are supported at the moment")
	}

	subnetCIDR := val.(string)

	if subnetsMap == nil {
		subnetsMap = subnets()
	}

	if _, ok := subnetsMap[subnetCIDR]; ok {
		return fmt.Errorf("subnet %s already exist", subnetCIDR)
	}

	return nil
}

/*
func getSubnetID(subnetCIDR string) (string, error) {
	_, ipv4Net, err := net.ParseCIDR(subnetCIDR)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("subnet-%03d", ipv4Net.IP[2]), nil
}
*/

func subnetHelp(networkCIDR string) (string, error) {
	_, ipv4Net, err := net.ParseCIDR(networkCIDR)
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("%d.%d.n.0/24", ipv4Net.IP[0], ipv4Net.IP[1])

	text := fmt.Sprintf("Network %s: A valid /24 subnet with format '%s' is required", networkCIDR, s)

	return text, nil
}

func GetSecurityPolicy(text string) topology.SecurityPolicy {
	secPolicies := map[string]topology.SecurityPolicy{
		topology.SecurityPolicy_ACCEPT.String(): topology.SecurityPolicy_ACCEPT,
		topology.SecurityPolicy_DROP.String():   topology.SecurityPolicy_DROP,
	}

	policies := []string{
		topology.SecurityPolicy_ACCEPT.String(),
		topology.SecurityPolicy_DROP.String(),
	}

	p := input.GetSelect(text, "", policies, survey.Required)

	return secPolicies[p]
}
