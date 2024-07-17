package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/iam"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) List(acls map[string]*iam.ACL) {
	output.SectionHeader("IAM: ACLs")
	output.TitleT1("ACL List")

	t := table.New()
	t.Header(colors.Black("ACCOUNT ID"), colors.Black("ACL"))

	for _, a := range acls {
		t.AddRow(a.AccountID, colors.DarkWhite(a.ACLID))
	}

	t.Render()
	fmt.Println()
}
