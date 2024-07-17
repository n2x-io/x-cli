package tenant

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/resource"
	"n2x.dev/x-api-go/grpc/resources/tenant"
	"n2x.dev/x-cli/pkg/client/account"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-cli/pkg/vars"
	"n2x.dev/x-lib/pkg/utils/msg"
)

var tenantsMap map[string]*tenant.Tenant = nil
var selectedTenant *tenant.Tenant = nil

func GetTenant() *tenant.Tenant {
	if selectedTenant != nil {
		return selectedTenant
	}

	tl := Tenants()

	if len(tl) == 0 {
		msg.Info("No tenant found")
		os.Exit(1)
	}

	var tenantOptID string
	tenantsOpts := make([]string, 0)
	tenants := make(map[string]*tenant.Tenant)

	for _, t := range tl {
		tenantOptID = fmt.Sprintf("[%s] %s", t.Name, t.Description)
		tenantsOpts = append(tenantsOpts, tenantOptID)
		tenants[tenantOptID] = t
	}

	sort.Strings(tenantsOpts)

	tenantOptID = input.GetSelect("Tenant:", "", tenantsOpts, survey.Required)

	vars.TenantID = tenants[tenantOptID].TenantID
	selectedTenant = tenants[tenantOptID]

	return tenants[tenantOptID]
}

func Tenants() map[string]*tenant.Tenant {
	if tenantsMap != nil {
		return tenantsMap
	}

	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetTenantAPIClient()
	defer grpcConn.Close()

	lr := &tenant.ListTenantsRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	tenants := make(map[string]*tenant.Tenant)

	for {
		tl, err := nxc.ListTenants(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list tenants")
		}

		for _, t := range tl.Tenants {
			tenants[t.TenantID] = t
		}

		if len(tl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = tl.Meta.NextPageToken
		} else {
			break
		}
	}

	tenantsMap = tenants

	return tenants
}

func FetchTenant(tenantID string) *tenant.Tenant {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetTenantAPIClient()
	defer grpcConn.Close()

	tr := &tenant.TenantReq{
		AccountID: a.AccountID,
		TenantID:  tenantID,
	}

	t, err := nxc.GetTenant(context.TODO(), tr)
	if err != nil {
		status.Error(err, "Unable to get tenant")
	}

	return t
}

func validTenantName(val interface{}) error {
	if err := input.ValidID(val); err != nil {
		return err
	}

	tenantName := val.(string)

	if tenantsMap == nil {
		tenantsMap = Tenants()
	}

	for _, t := range tenantsMap {
		if t.Name == tenantName {
			return fmt.Errorf("tenant %s already exist", tenantName)
		}
	}

	return nil
}
