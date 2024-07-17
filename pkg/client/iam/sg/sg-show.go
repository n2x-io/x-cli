package sg

import "n2x.dev/x-cli/pkg/client/tenant"

func (api *API) Show() {
	tenantMap := tenant.Tenants()

	Output().Show(GetSecurityGroup(false), tenantMap)
}
