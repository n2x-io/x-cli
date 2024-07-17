package role

import (
	"context"

	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Delete() {
	role := GetRole(false)

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	sr, err := nxc.DeleteRole(context.TODO(), role)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete role")
	}

	s.Stop()

	output.Show(sr)
}
