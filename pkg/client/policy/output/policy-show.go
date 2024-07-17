package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) Show(s *topology.Subnet) {
	output.SectionHeader("Security Policy Details")
	output.TitleT1("Security Policy")

	t := table.New()

	t.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(s.TenantID))
	t.AddRow(colors.Black("Network ID"), colors.DarkWhite(s.NetID))
	t.AddRow(colors.Black("Subnet ID"), colors.DarkWhite(s.SubnetID))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(s.Description))
	t.AddRow(colors.Black("Subnet CIDR"), colors.DarkWhite(s.IPAM.SubnetCIDR))

	t.Render()

	ShowNetworkPolicy(s.NetworkPolicy)
}

func ShowNetworkPolicy(np *topology.Policy) {
	t := table.New()

	switch np.DefaultPolicy {
	case topology.SecurityPolicy_ACCEPT:
		t.AddRow(colors.Black("Default Policy"), output.StrEnabled(topology.SecurityPolicy_ACCEPT.String()))
	case topology.SecurityPolicy_DROP:
		t.AddRow(colors.Black("Default Policy"), output.StrDisabled(topology.SecurityPolicy_DROP.String()))
	}

	t.Render()
	fmt.Println()

	t = table.New()
	// t.Header(colors.Black("Index"), colors.Black("Source"), colors.Black("Destination"), colors.Black("Port/Proto"), colors.Black("Policy"))

	// t.AddRow(output.TableHeader("Index"), output.TableHeader("Source"), output.TableHeader("Destination"), output.TableHeader("Port/Proto"), output.TableHeader("Policy"))

	t.AddRow(colors.Black("Index"), colors.Black("Source"), colors.Black("Destination"), colors.Black("Port/Proto"), colors.Black("Policy"))
	t.AddRow(colors.Black("-----"), colors.Black("------"), colors.Black("-----------"), colors.Black("----------"), colors.Black("------"))

	for _, nf := range np.NetworkFilters {
		var pol string
		switch nf.Policy {
		case topology.SecurityPolicy_ACCEPT:
			pol = output.StrEnabled(topology.SecurityPolicy_ACCEPT.String())
		case topology.SecurityPolicy_DROP:
			pol = output.StrDisabled(topology.SecurityPolicy_DROP.String())
		}

		var portProto string
		if nf.DstPort == 0 {
			portProto = nf.Proto.String()
		} else {
			portProto = fmt.Sprintf("%d/%s", nf.DstPort, nf.Proto.String())
		}

		t.AddRow(
			fmt.Sprintf("%d", nf.Index),
			nf.SrcIPNet,
			nf.DstIPNet,
			portProto,
			pol,
		)
	}

	t.Render()
	fmt.Println()
}
