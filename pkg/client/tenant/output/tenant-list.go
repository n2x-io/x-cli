package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/tenant"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) List(tenants map[string]*tenant.Tenant) {
	output.SectionHeader("Tenants")
	output.TitleT1("Tenant List")

	t := table.New()
	t.Header(colors.Black("TENANT NAME"), colors.Black("DESCRIPTION"))

	for _, tenant := range tenants {
		t.AddRow(colors.DarkWhite(tenant.Name), output.Fit(tenant.Description, 48))
	}

	t.Render()
	fmt.Println()
}
