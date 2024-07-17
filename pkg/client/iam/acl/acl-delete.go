package acl

import (
	"context"

	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Delete() {
	acl := GetACL(false)

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	sr, err := nxc.DeleteACL(context.TODO(), acl)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete ACL")
	}

	s.Stop()

	output.Show(sr)
}
