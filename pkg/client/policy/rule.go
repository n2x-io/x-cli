package policy

/*
import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/node"
	"n2x.dev/x-cli/pkg/client/subnet"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-cli/pkg/vars"
	"n2x.dev/x-lib/pkg/ipnet"
	"n2x.dev/x-lib/pkg/utils/msg"
)

type ruleObjectType string

const (
	endpointObject  ruleObjectType = "Endpoint Object"
	ipNetCIDRObject ruleObjectType = "IPNet CIDR"
)

type ruleEndpoint struct {
	objectType ruleObjectType
	af         ipnet.AddressFamily
	addr       string
}

func setNetworkPolicyRule(s *topology.Subnet) *topology.SetNetworkPolicyRequest {
	np := s.NetworkPolicy

	i, nf := getNetworkFilter(np, true)
	if nf != nil && i != -1 { // editing existing resource
		output.Choice("Edit Object")
	} else { // <new> resource
		output.Choice("New Object")

		nf = &topology.Filter{}
	}

	nf.Index = getPolicyIndex(nf.Index)

	nf.Description = input.GetInput("Description (optional):", "", nf.Description, nil)

	af := getAddressFamily()

	src := getRuleEndpoint(false, nf.SrcIPNet, af)
	nf.SrcIPNet = src.addr
	dst := getRuleEndpoint(true, nf.DstIPNet, af)
	nf.DstIPNet = dst.addr

	var protocols []string
	switch af {
	case ipnet.AddressFamilyIPv4:
		protocols = []string{"TCP", "UDP", "ICMPv4", "ANY"}
	case ipnet.AddressFamilyIPv6:
		protocols = []string{"TCP", "UDP", "ICMPv6", "ANY"}
	}

	nf.Proto = input.GetSelect("Protocol:", "", protocols, survey.Required)

	if nf.Proto == "TCP" || nf.Proto == "UDP" {
		helpText := "Destination TCP or UDP port"
		defaultPort := fmt.Sprintf("%d", nf.DstPort)

		dstP := input.GetInput("Destination Port:", helpText, defaultPort, input.ValidPort)

		p, err := strconv.Atoi(dstP)
		if err != nil {
			status.Error(err, "Invalid port")
		}
		nf.DstPort = int32(p)
	}

	nf.Policy = subnet.GetSecurityPolicy("Security Policy:")

	if np.NetworkFilters == nil {
		np.NetworkFilters = make([]*topology.Filter, 0)
	}

	if i == -1 { // new element
		np.NetworkFilters = append(np.NetworkFilters, nf)
	} else {
		np.NetworkFilters[i] = nf
	}

	return &topology.SetNetworkPolicyRequest{
		AccountID:     s.AccountID,
		TenantID:      s.TenantID,
		NetID:         s.NetID,
		SubnetID:      s.SubnetID,
		NetworkPolicy: np,
	}
}

func unsetNetworkPolicyRule(s *topology.Subnet) *topology.SetNetworkPolicyRequest {
	np := s.NetworkPolicy

	i, _ := getNetworkFilter(np, false)

	if np.NetworkFilters == nil {
		np.NetworkFilters = make([]*topology.Filter, 0)
	}

	// Remove the element at index i from np.NetworkFilters
	copy(np.NetworkFilters[i:], np.NetworkFilters[i+1:])             // Shift a[i+1:] left one index.
	np.NetworkFilters[len(np.NetworkFilters)-1] = nil                // Erase last element (write nil value).
	np.NetworkFilters = np.NetworkFilters[:len(np.NetworkFilters)-1] // Truncate slice.

	return &topology.SetNetworkPolicyRequest{
		AccountID:     s.AccountID,
		TenantID:      s.TenantID,
		NetID:         s.NetID,
		SubnetID:      s.SubnetID,
		NetworkPolicy: np,
	}
}

func getNetworkFilter(np *topology.Policy, edit bool) (int, *topology.Filter) {
	var objects []string

	nfilters := make(map[string]int)

	for i, nf := range np.NetworkFilters {
		obj := fmt.Sprintf("%d: src %s | dst %s | %d/%s | %s", nf.Index, nf.SrcIPNet, nf.DstIPNet, nf.DstPort, nf.Proto, nf.Policy)
		objects = append(objects, obj)
		nfilters[obj] = i
	}

	sort.Strings(objects)
	if edit {
		objects = append(objects, input.NewResource)
	}

	if len(objects) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	objID := input.GetSelect("Network Policy", "", objects, survey.Required)

	if objID == input.NewResource {
		return -1, nil
	}

	i := nfilters[objID]

	return i, np.NetworkFilters[i]
}

func ipNetHelp() string {
	text := `CIDR notation and prefix lenght, like '192.168.1.14/32', '10.32.4.0/24',
'fd77:f:a45:2d2:1::/128' or '2001:db8::/32' as defined in RFC 4632 and RFC 4291`

	return text
}

func getPolicyIndex(defIdx uint32) uint32 {
	index := input.GetInput("Policy Index:", "", fmt.Sprintf("%d", defIdx), input.ValidUint)

	i, err := strconv.Atoi(index)
	if err != nil {
		status.Error(err, "Invalid index")
	}

	return uint32(i)
}

func getRuleEndpoint(dst bool, def string, af ipnet.AddressFamily) *ruleEndpoint {
	var ruleEndp *ruleEndpoint

	vars.NodeID = ""

	endpMsg := "Source"
	if dst {
		endpMsg = "Destination"
	}

	var endpObj string
	endpOpts := []string{string(endpointObject), string(ipNetCIDRObject)}

	inputText := fmt.Sprintf("Select %s:", endpMsg)

	endpObj = input.GetSelect(inputText, "", endpOpts, survey.Required)

	if endpObj == string(endpointObject) {
		n := node.GetNode(false)
		e := node.GetEndpoint(n)
		ruleEndp = &ruleEndpoint{
			objectType: endpointObject,
			af:         af,
		}
		switch af {
		case ipnet.AddressFamilyIPv4:
			ruleEndp.addr = fmt.Sprintf("%s/32", e.IPv4)
		case ipnet.AddressFamilyIPv6:
			ruleEndp.addr = fmt.Sprintf("%s/128", e.IPv6)
		}
	} else { // endpObj == ipNetCIDRObject
		var ok bool
		var addr string

		inputMsg := fmt.Sprintf("%s %s CIDR:", endpMsg, af.String())

		for !ok {
			addr = input.GetInput(inputMsg, ipNetHelp(), def, input.ValidIPNetCIDR)

			switch af {
			case ipnet.AddressFamilyIPv4:
				if strings.Contains(addr, ":") {
					msg.Error("Invalid IPv4 address")
					continue
				}
			case ipnet.AddressFamilyIPv6:
				if strings.Contains(addr, ".") {
					msg.Error("Invalid IPv6 address")
					continue
				}
			}
			ok = true
		}

		ruleEndp = &ruleEndpoint{
			objectType: ipNetCIDRObject,
			addr:       addr,
			af:         af,
		}
	}

	return ruleEndp
}

func getAddressFamily() ipnet.AddressFamily {
	opts := []string{
		ipnet.AddressFamilyIPv4.String(),
		ipnet.AddressFamilyIPv6.String(),
	}

	afOpt := input.GetSelect("Protocol:", "", opts, survey.Required)

	switch afOpt {
	case ipnet.AddressFamilyIPv4.String():
		return ipnet.AddressFamilyIPv4
	case ipnet.AddressFamilyIPv6.String():
		return ipnet.AddressFamilyIPv6
	}

	return ipnet.AddressFamilyUnspec
}
*/
