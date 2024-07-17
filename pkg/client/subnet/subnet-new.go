package subnet

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/network"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) New() {
	n := network.GetNetwork(false)

	nsr := &topology.NewSubnetRequest{
		AccountID:   n.AccountID,
		TenantID:    n.TenantID,
		NetID:       n.NetID,
		NetworkCIDR: n.NetworkCIDR,
	}

	// global var needed by the validation function
	networkCIDR = n.NetworkCIDR

	helpText, err := subnetHelp(networkCIDR)
	if err != nil {
		status.Error(err, "Unable to parse network CIDR")
	}

	nsr.SubnetCIDR = input.GetInput("Subnet CIDR:", helpText, "", validSubnet)

	nsr.Description = input.GetInput("Subnet Description:", "", "", survey.Required)

	nsr.DefaultPolicy = GetSecurityPolicy("Default Security Policy:")

	ss := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	s, err := nxc.CreateSubnet(context.TODO(), nsr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to create subnet")
	}

	ss.Stop()

	// output.Show(s)
	Output().Show(s)
}
