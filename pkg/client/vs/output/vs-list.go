package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) List(vss map[string]*topology.VS) {
	output.SectionHeader("Virtual Servers")
	output.TitleT1("VS List")

	t := table.New()
	t.Header(colors.Black("VS ID"), colors.Black("NAME"), colors.Black("DESCRIPTION"))

	for _, vs := range vss {
		t.AddRow(vs.VSID, colors.DarkWhite(vs.Name), output.Fit(vs.Description, 32))
	}

	t.Render()
	fmt.Println()
}
