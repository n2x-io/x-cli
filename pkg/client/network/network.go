package network

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/resource"
	tenant_pb "n2x.dev/x-api-go/grpc/resources/tenant"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/tenant"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-cli/pkg/vars"
	"n2x.dev/x-lib/pkg/utils/msg"
)

var networksMap map[string]*topology.Network = nil
var selectedNetwork *topology.Network = nil

func GetNetwork(edit bool) *topology.Network {
	if selectedNetwork != nil {
		return selectedNetwork
	}

	nl := networks()

	if len(nl) == 0 {
		msg.Info("No networks found")
		os.Exit(1)
	}

	var networkOptID string
	networksOpts := make([]string, 0)
	networks := make(map[string]*topology.Network)

	for _, n := range nl {
		networkOptID = fmt.Sprintf("[%s] %s", n.NetID, n.Description)
		networksOpts = append(networksOpts, networkOptID)
		networks[networkOptID] = n
	}

	sort.Strings(networksOpts)

	if edit {
		networksOpts = append(networksOpts, input.NewResource)
	}

	networkOptID = input.GetSelect("Network:", "", networksOpts, survey.Required)

	if networkOptID == input.NewResource {
		return nil
	}

	vars.NetID = networks[networkOptID].NetID
	selectedNetwork = networks[networkOptID]

	return networks[networkOptID]
}

func networks() map[string]*topology.Network {
	if networksMap != nil {
		return networksMap
	}

	t := tenant.GetTenant()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	lr := &topology.ListNetworksRequest{
		Meta: &resource.ListRequest{},
		Tenant: &tenant_pb.TenantReq{
			AccountID: t.AccountID,
			TenantID:  t.TenantID,
		},
	}

	networks := make(map[string]*topology.Network)

	for {
		nl, err := nxc.ListNetworks(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list networks")
		}

		for _, n := range nl.Networks {
			networks[n.NetID] = n
		}

		if len(nl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = nl.Meta.NextPageToken
		} else {
			break
		}
	}

	networksMap = networks

	return networks
}
