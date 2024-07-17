package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) Show(n *topology.Network) {
	output.SectionHeader("Network Details")
	output.TitleT1("Network Information")

	t := table.New()

	// t.AddRow(colors.Black("Account ID"), colors.DarkWhite(n.AccountID))
	t.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(n.TenantID))
	t.AddRow(colors.Black("Network ID"), colors.DarkWhite(n.NetID))
	t.AddRow(colors.Black("Network CIDR"), colors.DarkWhite(n.NetworkCIDR))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(n.Description))

	if n.RoutedSubnets {
		t.AddRow(colors.Black("Routed Subnets"), output.StrEnabled("yes"))
	} else {
		t.AddRow(colors.Black("Routed Subnets"), output.StrDisabled("no"))
	}
	t.AddRow(colors.Black("Connectivity Zone"), colors.DarkWhite(n.LocationID))

	t.Render()
	fmt.Println()
}
