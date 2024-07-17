package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/ops"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) Show(p *ops.Project) {
	output.SectionHeader("Ops: Project Details")
	output.TitleT1("Project Information")

	t := table.New()

	// t.AddRow(colors.Black("Account ID"), colors.DarkWhite(p.AccountID))
	t.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(p.TenantID))
	t.AddRow(colors.Black("Project ID"), colors.DarkWhite(p.ProjectID))
	t.AddRow(colors.Black("Name"), colors.DarkWhite(p.Name))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(p.Description))

	t.Render()
	fmt.Println()

	output.SubTitleT2("ChatOps and Service Management Features")

	t = table.New()

	if p.ReviewRequired {
		t.AddRow(colors.Black("Review Required"), output.StrEnabled("yes"))
	} else {
		t.AddRow(colors.Black("Review Required"), output.StrDisabled("no"))
	}
	if p.ApprovalRequired {
		t.AddRow(colors.Black("Approval Required"), output.StrEnabled("yes"))
	} else {
		t.AddRow(colors.Black("Approval Required"), output.StrDisabled("no"))
	}

	t.Render()
	fmt.Println()
}
