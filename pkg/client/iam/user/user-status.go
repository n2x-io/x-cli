package user

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/iam"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Enable() {
	setStatus(true)
}

func (api *API) Disable() {
	setStatus(false)
}

func setStatus(enabled bool) {
	u := GetUser(false)

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	ur := &iam.UserReq{
		AccountID: u.AccountID,
		UserID:    u.UserID,
	}

	s := output.Spinner()

	var err error
	if enabled {
		u, err = nxc.EnableUser(context.TODO(), ur)
	} else {
		u, err = nxc.DisableUser(context.TODO(), ur)
	}
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set user")
	}

	s.Stop()

	Output().Show(u)
}
