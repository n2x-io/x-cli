package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/iam"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) List(securityGroups map[string]*iam.SecurityGroup) {
	output.SectionHeader("IAM: Security Groups")
	output.TitleT1("Security Group List")

	t := table.New()
	t.Header(colors.Black("ACCOUNT ID"), colors.Black("SECURITY GROUP"))

	for _, sg := range securityGroups {
		t.AddRow(sg.AccountID, colors.DarkWhite(sg.SecurityGroupID))
	}

	t.Render()
	fmt.Println()
}
