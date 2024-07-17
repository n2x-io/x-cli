package policy

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/subnet"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Delete() {
	s := subnet.GetSubnet(false)

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	ss := output.Spinner()

	sr := &topology.SubnetReq{
		AccountID: s.AccountID,
		TenantID:  s.TenantID,
		NetID:     s.NetID,
		SubnetID:  s.SubnetID,
	}

	np, err := nxc.DeleteNetworkPolicy(context.TODO(), sr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to delete network policy")
	}

	s.NetworkPolicy = np

	ss.Stop()

	// output.Show(s)
	Output().Show(s)
}
