package subnet

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Delete() {
	s := GetSubnet(false)

	output.ConfirmDeletion()

	ss := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	sr := &topology.SubnetReq{
		AccountID: s.AccountID,
		TenantID:  s.TenantID,
		NetID:     s.NetID,
		SubnetID:  s.SubnetID,
	}

	r, err := nxc.DeleteSubnet(context.TODO(), sr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to delete subnet")
	}

	ss.Stop()

	output.Show(r)
}
