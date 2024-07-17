package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/tenant"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/ipnet"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) Show(vs *topology.VS) {
	output.SectionHeader("Virtual Server Details")
	output.TitleT1("Virtual Server Information")

	t := table.New()

	t.AddRow(colors.Black("VS ID"), colors.DarkWhite(vs.VSID))
	t.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(vs.TenantID))
	t.AddRow(colors.Black("Name"), colors.DarkWhite(vs.Name))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(vs.Description))
	t.AddRow(colors.Black("Connectivity Zone"), colors.DarkWhite(vs.LocationID))
	if len(vs.Cname) > 0 {
		t.AddRow(colors.Black("Custom DNS CNAME"), colors.DarkWhite(vs.Cname))
	}
	if vs.ReqAuth {
		t.AddRow(colors.Black("Authentication"), output.StrEnabled("required"))
	} else {
		t.AddRow(colors.Black("Authentication"), output.StrDisabled("no"))
	}
	// t.AddRow(colors.Black("Protocol"), colors.DarkWhite(getProto(as.Proto)))
	// t.AddRow(colors.Black("Port"), colors.DarkWhite(fmt.Sprintf("%d", vs.VSPort)))

	t.Render()
	fmt.Println()

	if len(vs.AppSvcs) > 0 {
		output.SubTitleT2("Node App Services")

		s := output.Spinner()

		t = table.New()

		t.AddRow(colors.Black("Tenant"), colors.Black("Subnet"), colors.Black("Node"), colors.Black("Proto"), colors.Black("Port"))
		t.AddRow(colors.Black("------"), colors.Black("------"), colors.Black("----"), colors.Black("-----"), colors.Black("----"))

		for _, as := range vs.AppSvcs {
			tenantName := tenant.FetchTenant(vs.TenantID).Name
			t.AddRow(
				colors.DarkWhite(tenantName),
				// colors.DarkWhite(as.NetID),
				colors.DarkWhite(as.SubnetID),
				colors.DarkWhite(output.Fit(as.NodeName, 32)),
				colors.DarkWhite(getProto(as.Proto)),
				colors.DarkWhite(fmt.Sprintf("%d", as.RSPort)),
				// colors.DarkWhite(" "),
			)
		}

		s.Stop()

		t.Render()
		fmt.Println()
	}

	if len(vs.Tags) > 0 {
		output.SubTitleT2("Tags")

		for _, tag := range vs.Tags {
			fmt.Printf(" ■ %s\n", colors.DarkGreen(tag))
		}
		fmt.Println()
	}

	fmt.Println()

	fmt.Printf("%s %s: %s\n\n",
		colors.DarkCyan("->"),
		colors.DarkWhite("URL"),
		colors.DarkBlue(getURL(vs)),
	)
}

func getURL(vs *topology.VS) string {
	if len(vs.Cname) > 0 {
		return fmt.Sprintf("https://%s", vs.Cname)
	}

	return fmt.Sprintf("https://%s.%s.%s", vs.VSID, vs.LocationID, ipnet.IAPDomain())
}

func getProto(proto topology.VSProto) string {
	switch proto {
	case topology.VSProto_PROTO_TCP_GENERIC:
		return "TCP"
	case topology.VSProto_PROTO_TCP_HTTP:
		return "TCP"
	case topology.VSProto_PROTO_TCP_HTTPS:
		return "TCP"
	case topology.VSProto_PROTO_TCP_SSH:
		return "TCP"
	case topology.VSProto_PROTO_UDP_GENERIC:
		return "UDP"
	}

	return "unknown"
}
