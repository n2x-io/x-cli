package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) ShowEndpoint(e *topology.Endpoint) {
	output.SectionHeader("Endpoint Details")
	output.TitleT1("Network Endpoint")

	t := table.New()

	t.AddRow(colors.Black("Endpoint"), colors.DarkWhite(e.EndpointID))
	t.AddRow(colors.Black("FQDN"), colors.DarkWhite(e.DNSName+".n2x.local"))
	t.AddRow(colors.Black("IPv4"), colors.DarkWhite(e.IPv4))
	t.AddRow(colors.Black("IPv6"), colors.DarkWhite(e.IPv6))

	t.Render()
	fmt.Println()
}
