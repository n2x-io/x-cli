package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/iam"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) List(roles map[string]*iam.Role) {
	output.SectionHeader("IAM: Roles")
	output.TitleT1("Role List")

	t := table.New()
	t.Header(colors.Black("ACCOUNT ID"), colors.Black("ROLE"))

	for _, r := range roles {
		t.AddRow(r.AccountID, colors.DarkWhite(r.RoleID))
	}

	t.Render()
	fmt.Println()
}
