package tenant

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/tenant"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Delete() {
	t := GetTenant()

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTenantAPIClient()
	defer grpcConn.Close()

	tr := &tenant.TenantReq{
		AccountID: t.AccountID,
		TenantID:  t.TenantID,
	}

	sr, err := nxc.DeleteTenant(context.TODO(), tr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete tenant")
	}

	s.Stop()

	output.Show(sr)
}
