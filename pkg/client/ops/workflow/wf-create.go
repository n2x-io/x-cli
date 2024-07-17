package workflow

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/ops"
	"n2x.dev/x-cli/pkg/client/ops/project"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/utils"
)

func (api *API) Create(yamlFile string) {
	p := project.GetProject()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	wf := &ops.Workflow{}

	if err := utils.FileParser(yamlFile, wf); err != nil {
		s.Stop()
		status.Error(err, "Unable to parse YAML file")
	}

	wf.AccountID = p.AccountID
	wf.TenantID = p.TenantID
	wf.ProjectID = p.ProjectID

	wf, err := nxc.CreateWorkflow(context.TODO(), wf)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to create workflow")
	}

	s.Stop()

	Output().Show(wf)
}
