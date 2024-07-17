package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/iam"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) List(users map[string]*iam.User) {
	output.SectionHeader("IAM: Users")
	output.TitleT1("User List")

	t := table.New()
	t.Header(colors.Black("ACCOUNT ID"), colors.Black("USER EMAIL"))

	for _, u := range users {
		t.AddRow(u.AccountID, colors.DarkWhite(u.Email))
	}

	t.Render()
	fmt.Println()
}
