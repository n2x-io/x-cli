package tenant

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/tenant"
	"n2x.dev/x-cli/pkg/client/account"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) New() {
	a := account.GetAccount()

	ntr := &tenant.NewTenantRequest{
		AccountID:   a.AccountID,
		Name:        input.GetInput("Tenant Name:", "", "", validTenantName),
		Description: input.GetInput("Description:", "", "", survey.Required),
	}

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTenantAPIClient()
	defer grpcConn.Close()

	t, err := nxc.CreateTenant(context.TODO(), ntr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set tenant")
	}

	s.Stop()

	// output.Show(t)
	Output().Show(t)
}
