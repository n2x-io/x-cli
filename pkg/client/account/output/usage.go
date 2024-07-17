package output

import (
	"fmt"
	"time"

	"n2x.dev/x-api-go/grpc/resources/account"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func showUsage(usage *account.Usage) {
	if usage == nil {
		return
	}

	output.SubTitleT2("Account Stats")

	t := table.New()

	// t.AddRow(output.TableHeader("Network"), "", output.TableHeader("IAM"), "", output.TableHeader("Automation"), "")
	t.AddRow(colors.Black("NETWORK"), "", colors.Black("IAM"), "", colors.Black("AUTOMATION"), "")
	t.AddRow(colors.Black("-------"), "", colors.Black("---"), "", colors.Black("----------"), "")

	t.AddRow(colors.Black("Tenants"), colors.DarkWhite(fmt.Sprintf("%v", usage.Tenants)),
		colors.Black("Users"), colors.DarkWhite(fmt.Sprintf("%v", usage.Users)),
		colors.Black("Projects"), colors.DarkWhite(fmt.Sprintf("%v", usage.Projects)))
	t.AddRow(colors.Black("Networks"), colors.DarkWhite(fmt.Sprintf("%v", usage.Networks)),
		colors.Black("Security Groups"), colors.DarkWhite(fmt.Sprintf("%v", usage.SecurityGroups)),
		colors.Black("Workflows"), colors.DarkWhite(fmt.Sprintf("%v", usage.Workflows)))
	t.AddRow(colors.Black("Subnets"), colors.DarkWhite(fmt.Sprintf("%v", usage.Subnets)),
		colors.Black("Roles"), colors.DarkWhite(fmt.Sprintf("%v", usage.Roles)))
	t.AddRow(colors.Black("Nodes"), colors.DarkWhite(fmt.Sprintf("%v", usage.Nodes)),
		colors.Black("ACLs"), colors.DarkWhite(fmt.Sprintf("%v", usage.ACLs)))
	t.AddRow(colors.Black("Endpoints"), colors.DarkWhite(fmt.Sprintf("%v", usage.Endpoints)))
	t.AddRow(colors.Black("Tier-1 Nodes"), colors.DarkWhite(fmt.Sprintf("%v", usage.Relays)))

	t.Render()
	fmt.Println()

	if usage.LastUpdated > 0 {
		t := table.New()

		tm := time.Unix(usage.LastUpdated, 0)
		t.AddRow(colors.Black("Last Updated"), colors.DarkWhite(tm.String()))

		t.Render()
		fmt.Println()
	}
}
