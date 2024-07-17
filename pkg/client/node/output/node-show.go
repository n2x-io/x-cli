package output

import (
	"fmt"

	// "github.com/c2h5oh/datasize"
	"github.com/gosuri/uitable"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/event"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) Show(n *topology.Node) {
	output.SectionHeader("Node Details")
	output.TitleT1("Node Information")

	agentInfo(n)

	if n.Agent.Routes != nil {
		output.SubTitleT2("Routing: Advertised Routes")

		for _, r := range n.Agent.Routes.Export {
			fmt.Printf(" ■ %s\n", colors.DarkGreen(r))
		}
		fmt.Println()

		output.SubTitleT2("Routing: Imported Routes")

		for _, r := range n.Agent.Routes.Import {
			fmt.Printf(" ■ %s\n", colors.DarkGreen(r))
		}
		fmt.Println()
	}

	if n.EventMetrics != nil {
		event.Output().ShowMetrics(n.EventMetrics)
	}

	// if n.Agent.Metrics.NetworkMonthlyHistory != nil {
	// 	output.SubTitleT2("Network Traffic")

	// 	t := uitable.New()
	// 	t.MaxColWidth = 16
	// 	t.Wrap = false
	// 	t.RightAlign(1)
	// 	t.RightAlign(2)
	// 	t.RightAlign(3)

	// 	// t.AddRow(output.TableHeader("Month"), output.TableHeader("Tx Total"), output.TableHeader("Rx Total"), output.TableHeader("Dropped Pkts"))
	// 	t.AddRow(colors.Black("Month"), colors.Black("Tx Total"), colors.Black("Rx Total"), colors.Black("Dropped Pkts"))
	// 	t.AddRow(colors.Black("-----"), colors.Black("--------"), colors.Black("--------"), colors.Black("------------"))
	// 	for month, nh := range n.Agent.Metrics.NetworkMonthlyHistory {
	// 		t.AddRow(
	// 			colors.DarkWhite(month),
	// 			datasize.ByteSize(nh.TxTotalBytes).HumanReadable(),
	// 			datasize.ByteSize(nh.RxTotalBytes).HumanReadable(),
	// 			nh.DroppedPkts,
	// 		)
	// 	}

	// 	fmt.Println(t)
	// 	fmt.Println()
	// }

	if len(n.Endpoints) > 0 {
		output.SubTitleT2("Node Endpoints")

		t := uitable.New()
		t.MaxColWidth = 36
		t.Wrap = false

		// t.AddRow(output.TableHeader("Endpoint ID / FQDN"), output.TableHeader("IPv4"), output.TableHeader("IPv6"))
		t.AddRow(colors.Black("Endpoint ID / FQDN"), colors.Black("IPv4"), colors.Black("IPv6"))
		t.AddRow(colors.Black("------------------"), colors.Black("----"), colors.Black("----"))
		for _, e := range n.Endpoints {
			fqdn := colors.DarkWhite(e.DNSName + ".n2x.local")
			t.AddRow(e.EndpointID)
			t.AddRow(fqdn, colors.DarkWhite(e.IPv4), e.IPv6)
		}

		fmt.Println(t)
		fmt.Println()
	}
}
