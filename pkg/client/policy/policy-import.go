package policy

/*
import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/account"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/utils"
)

func (api *API) Import(yamlFile string) {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	s := output.Spinner()

	snpr := &topology.SetNetworkPolicyRequest{}

	if err := utils.FileParser(yamlFile, snpr); err != nil {
		s.Stop()
		status.Error(err, "Unable to parse YAML file")
	}

	snpr.AccountID = a.AccountID

	np, err := nxc.SetNetworkPolicy(context.TODO(), snpr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set network policy")
	}

	s.Stop()

	// output.Show(np)
	Output().Show(snpr.TenantID, snpr.NetID, snpr.SubnetID, np)
}
*/
