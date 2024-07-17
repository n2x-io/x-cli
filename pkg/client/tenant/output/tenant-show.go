package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/tenant"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) Show(t *tenant.Tenant) {
	output.SectionHeader("Tenant Details")
	output.TitleT1("Tenant Information")

	tbl := table.New()

	// tbl.AddRow(colors.Black("Account ID"), colors.DarkWhite(t.AccountID))
	tbl.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(t.TenantID))
	tbl.AddRow(colors.Black("Name"), colors.DarkWhite(t.Name))
	tbl.AddRow(colors.Black("Description"), colors.DarkWhite(t.Description))

	tbl.Render()
	fmt.Println()
}
