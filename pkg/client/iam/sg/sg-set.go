package sg

import (
	"context"
	"fmt"
	"sort"

	"n2x.dev/x-api-go/grpc/resources/iam"
	tenant_pb "n2x.dev/x-api-go/grpc/resources/tenant"
	"n2x.dev/x-cli/pkg/client/account"
	"n2x.dev/x-cli/pkg/client/tenant"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Set() {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	sg := GetSecurityGroup(true)
	if sg != nil { // editing existing resource
		output.Choice("Edit RBAC Security Group")
	} else { // <new> resource
		output.Choice("New RBAC Security Group")

		sg = &iam.SecurityGroup{
			AccountID: a.AccountID,
			Users:     make(map[string]bool),
		}

		sg.SecurityGroupID = input.GetInput("Security Group ID:", "", "", validSecurityGroupID)
	}

	tenantMap := tenant.Tenants()

	s := output.Spinner()

	// get all-tenant options
	tenantsOpts := make([]string, 0)
	tenants := make(map[string]*tenant_pb.Tenant)

	for _, t := range tenantMap {
		tenantOptID := fmt.Sprintf("[%s] %s", t.Name, t.Description)
		tenantsOpts = append(tenantsOpts, tenantOptID)
		tenants[tenantOptID] = t
	}

	sort.Strings(tenantsOpts)

	// get current security-group tenants
	currentTenantsOpts := make([]string, 0)
	for _, tenantID := range sg.TenantIDs {
		if t, ok := tenantMap[tenantID]; ok {
			tenantOptID := fmt.Sprintf("[%s] %s", t.Name, t.Description)
			currentTenantsOpts = append(currentTenantsOpts, tenantOptID)
		}
	}

	s.Stop()

	tenantsOpts = input.GetMultiSelect("Tenants:", "", tenantsOpts, currentTenantsOpts, nil)

	tenantIDs := make([]string, 0)

	for _, tenantOptID := range tenantsOpts {
		if t, ok := tenants[tenantOptID]; ok {
			tenantIDs = append(tenantIDs, t.TenantID)
		}
	}

	sg.TenantIDs = tenantIDs

	s = output.Spinner()

	sg, err := nxc.SetSecurityGroup(context.TODO(), sg)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set security-group")
	}

	s.Stop()

	Output().Show(sg, tenantMap)
}
