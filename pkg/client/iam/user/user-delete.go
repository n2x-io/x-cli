package user

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/iam"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Delete() {
	u := GetUser(false)

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	ur := &iam.UserReq{
		AccountID: u.AccountID,
		UserID:    u.UserID,
	}

	s := output.Spinner()

	sr, err := nxc.DeleteUser(context.TODO(), ur)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete user")
	}

	s.Stop()

	output.Show(sr)
}
