package tenant

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/tenant"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Update() {
	t := GetTenant()

	utr := &tenant.UpdateTenantRequest{
		AccountID:   t.AccountID,
		TenantID:    t.TenantID,
		Name:        input.GetInput("Tenant Name:", "", t.Name, survey.Required),
		Description: input.GetInput("Description:", "", t.Description, survey.Required),
	}

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTenantAPIClient()
	defer grpcConn.Close()

	t, err := nxc.UpdateTenant(context.TODO(), utr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set tenant")
	}

	s.Stop()

	// output.Show(t)
	Output().Show(t)
}
