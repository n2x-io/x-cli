package vs

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Delete() {
	vs := getVS(false)

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	vsr := &topology.VSReq{
		AccountID: vs.AccountID,
		VSID:      vs.VSID,
	}

	sr, err := nxc.DeleteVS(context.TODO(), vsr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete virtual server")
	}

	s.Stop()

	output.Show(sr)
}
