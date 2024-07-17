package sg

import (
	"context"

	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Delete() {
	sg := GetSecurityGroup(false)

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	sr, err := nxc.DeleteSecurityGroup(context.TODO(), sg)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete security-group")
	}

	s.Stop()

	output.Show(sr)
}
