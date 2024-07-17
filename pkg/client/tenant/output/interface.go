package output

import "n2x.dev/x-api-go/grpc/resources/tenant"

type Interface interface {
	List(tenants map[string]*tenant.Tenant)
	Show(t *tenant.Tenant)
}
type API struct{}
