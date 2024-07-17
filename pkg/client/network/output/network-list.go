package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) List(networks map[string]*topology.Network) {
	output.SectionHeader("Networks")
	output.TitleT1("Network List")

	t := table.New()
	t.Header(colors.Black("NETWORK ID"), colors.Black("NETWORK CIDR"), colors.Black("DESCRIPTION"))

	for _, n := range networks {
		t.AddRow(n.NetID, colors.DarkWhite(n.NetworkCIDR), output.Fit(n.Description, 32))
	}

	t.Render()
	fmt.Println()
}
