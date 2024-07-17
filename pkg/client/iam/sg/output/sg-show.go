package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/iam"
	"n2x.dev/x-api-go/grpc/resources/tenant"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) Show(sg *iam.SecurityGroup, tenantMap map[string]*tenant.Tenant) {
	output.SectionHeader("IAM: Security Group Details")
	output.TitleT1("Security Group Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(sg.AccountID))
	t.AddRow(colors.Black("Security Group"), colors.DarkWhite(sg.SecurityGroupID))

	t.Render()
	fmt.Println()

	if len(sg.TenantIDs) > 0 {
		output.SubTitleT2("Tenants")

		for _, tenantID := range sg.TenantIDs {
			if tenantID == ".*" {
				fmt.Printf(" ■ %s\n", colors.DarkGreen(tenantID))
				continue
			}

			if t, ok := tenantMap[tenantID]; ok {
				tenantStr := fmt.Sprintf("%s: %s", t.Name, t.Description)
				fmt.Printf(" ■ %s\n", colors.DarkGreen(tenantStr))
			}
		}
		fmt.Println()
	}

	if len(sg.Users) > 0 {
		output.SubTitleT2("Users")

		for u := range sg.Users {
			fmt.Printf(" ■ %s\n", colors.DarkGreen(u))
		}
		fmt.Println()
	}
}
