package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) List(subnets map[string]*topology.Subnet) {
	output.SectionHeader("Subnets")
	output.TitleT1("Subnet List")

	t := table.New()
	t.Header(colors.Black("SUBNET ID"), colors.Black("SUBNET CIDR"), colors.Black("DESCRIPTION"))

	for _, s := range subnets {
		if s.IPAM != nil {
			t.AddRow(s.SubnetID, colors.DarkWhite(s.IPAM.SubnetCIDR), output.Fit(s.Description, 32))
		}
	}

	t.Render()
	fmt.Println()
}
