package workflow

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/ops"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Delete() {
	wf := GetWorkflow()

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	wfr := &ops.WorkflowReq{
		AccountID:  wf.AccountID,
		TenantID:   wf.TenantID,
		ProjectID:  wf.ProjectID,
		WorkflowID: wf.WorkflowID,
	}

	sr, err := nxc.DeleteWorkflow(context.TODO(), wfr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete workflow")
	}

	s.Stop()

	output.Show(sr)
}
