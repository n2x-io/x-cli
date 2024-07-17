package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/iam"
	"n2x.dev/x-cli/pkg/client/event"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) Show(u *iam.User) {
	output.SectionHeader("IAM: User Details")
	output.TitleT1("User Information")

	ShowUser(u)
}

func ShowUser(u *iam.User) {
	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(u.AccountID))
	t.AddRow(colors.Black("User ID"), colors.DarkWhite(u.UserID))
	t.AddRow(colors.Black("Email"), colors.DarkWhite(u.Email))

	if u.Type == iam.UserType_USER_TYPE_ADM {
		t.AddRow(colors.Black("User Type"), output.StrBlue("adm"))
		// t.AddRow(colors.Black("Admin User"), output.StrEnabled("enabled"))
		// t.AddRow(colors.Black("IAP User"), output.StrDisabled("disabled"))
	}
	if u.Type == iam.UserType_USER_TYPE_IAP {
		t.AddRow(colors.Black("User Type"), output.StrBlue("proxy"))
		// t.AddRow(colors.Black("Admin User"), output.StrDisabled("disabled"))
		// t.AddRow(colors.Black("IAP User"), output.StrEnabled("enabled"))
	}

	if u.Status.Enabled {
		t.AddRow(colors.Black("User Status"), output.StrEnabled("enabled"))
	} else {
		t.AddRow(colors.Black("User Status"), output.StrDisabled("disabled"))
	}

	t.Render()
	fmt.Println()

	if u.RBAC != nil {
		if len(u.RBAC.SecurityGroups) > 0 {
			output.SubTitleT2("Security Groups")

			for _, sg := range u.RBAC.SecurityGroups {
				fmt.Printf(" ■ %s\n", colors.DarkGreen(sg))
			}
			fmt.Println()
		}

		if len(u.RBAC.ACLs) > 0 {
			output.SubTitleT2("ACLs")

			for _, acl := range u.RBAC.ACLs {
				fmt.Printf(" ■ %s\n", colors.DarkGreen(acl))
			}
			fmt.Println()
		}

		if len(u.RBAC.Roles) > 0 {
			output.SubTitleT2("Roles")

			for _, role := range u.RBAC.Roles {
				fmt.Printf(" ■ %s\n", colors.DarkGreen(role))
			}
			fmt.Println()
		}
	}

	if u.EventMetrics != nil {
		event.Output().ShowMetrics(u.EventMetrics)
	}
}
