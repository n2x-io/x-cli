package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/iam"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) Show(r *iam.Role) {
	output.SectionHeader("IAM: Role Details")
	output.TitleT1("Role Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(r.AccountID))
	t.AddRow(colors.Black("Role"), colors.DarkWhite(r.RoleID))

	t.Render()
	fmt.Println()

	if len(r.Permissions) > 0 {
		output.SubTitleT2("Permissions")

		for _, perm := range r.Permissions {
			fmt.Printf(" ■ %s\n", colors.DarkGreen(perm))
		}
		fmt.Println()
	}

	if len(r.Users) > 0 {
		output.SubTitleT2("Users")

		for u := range r.Users {
			fmt.Printf(" ■ %s\n", colors.DarkGreen(u))
		}
		fmt.Println()
	}
}
