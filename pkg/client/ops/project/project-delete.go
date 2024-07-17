package project

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/ops"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Delete() {
	p := GetProject()

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	pr := &ops.ProjectReq{
		AccountID: p.AccountID,
		TenantID:  p.TenantID,
		ProjectID: p.ProjectID,
	}

	sr, err := nxc.DeleteProject(context.TODO(), pr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete project")
	}

	s.Stop()

	output.Show(sr)
}
