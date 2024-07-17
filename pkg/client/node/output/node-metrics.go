package output

import (
	"fmt"
	"time"

	// "github.com/c2h5oh/datasize"
	"github.com/gosuri/uitable"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) Metrics(n *topology.Node) {
	output.SectionHeader("Node Details")
	output.TitleT1("Node Metrics")

	agentInfo(n)

	output.SubTitleT2("Network Metrics")

	/*
		if n.Agent.Metrics.NetworkMetrics == nil {
			fmt.Printf("[%s]\n\n", colors.Black("not enough data yet"))
			return
		}

		// traffic history
		fmt.Println("Traffic History")
		fmt.Printf("%s\n\n", colors.Black("==============="))

		t := uitable.New()
		t.MaxColWidth = 16
		t.Wrap = false
		t.RightAlign(1)
		t.RightAlign(2)
		t.RightAlign(3)

		// t.AddRow(output.TableHeader("Month"), output.TableHeader("Tx Total"), output.TableHeader("Rx Total"), output.TableHeader("Dropped Pkts"))
		t.AddRow(colors.Black("Month"), colors.Black("Tx Total"), colors.Black("Rx Total"), colors.Black("Dropped Pkts"))
		t.AddRow(colors.Black("-----"), colors.Black("--------"), colors.Black("--------"), colors.Black("------------"))
		for month, nh := range n.Agent.Metrics.NetworkMonthlyHistory {
			t.AddRow(
				colors.DarkWhite(month),
				datasize.ByteSize(nh.TxTotalBytes).HumanReadable(),
				datasize.ByteSize(nh.RxTotalBytes).HumanReadable(),
				nh.DroppedPkts,
			)
		}
		fmt.Println(t)
		fmt.Println()

		// traffic history
		fmt.Println("Last Month")
		fmt.Printf("%s\n\n", colors.Black("=========="))

		t = uitable.New()
		t.MaxColWidth = 16
		t.Wrap = false
		t.RightAlign(1)
		t.RightAlign(2)
		t.RightAlign(3)

		// t.AddRow(fmt.Sprintf("%s     ", output.TableHeader("Day")), output.TableHeader("Tx Total"), output.TableHeader("Rx Total"), output.TableHeader("Dropped Pkts"))
		t.AddRow(fmt.Sprintf("%s     ", colors.Black("Day")), colors.Black("Tx Total"), colors.Black("Rx Total"), colors.Black("Dropped Pkts"))
		t.AddRow(colors.Black("---     "), colors.Black("--------"), colors.Black("--------"), colors.Black("------------"))
		today := time.Now().Day()
		for d := 1; d < 32; d++ {
			idx := resources.NetworkHistoryDailyIndex(d)
			if nh, ok := n.Agent.Metrics.NetworkDailyHistory[idx]; ok {
				switch {
				case today == d:
					t.AddRow(
						colors.Magenta(idx),
						colors.Magenta(datasize.ByteSize(nh.TxTotalBytes).HumanReadable()),
						colors.Magenta(datasize.ByteSize(nh.RxTotalBytes).HumanReadable()),
						colors.Magenta(fmt.Sprintf("%d", nh.DroppedPkts)),
					)
				case today < d:
					t.AddRow(
						colors.DarkMagenta(idx),
						colors.DarkMagenta(datasize.ByteSize(nh.TxTotalBytes).HumanReadable()),
						colors.DarkMagenta(datasize.ByteSize(nh.RxTotalBytes).HumanReadable()),
						colors.DarkMagenta(fmt.Sprintf("%d", nh.DroppedPkts)),
					)
				case today > d:
					t.AddRow(
						idx,
						datasize.ByteSize(nh.TxTotalBytes).HumanReadable(),
						datasize.ByteSize(nh.RxTotalBytes).HumanReadable(),
						nh.DroppedPkts,
					)
				}
			}
		}
		fmt.Println(t)
		fmt.Println()

		// traffic history
		fmt.Println("Last Day")
		fmt.Printf("%s\n\n", colors.Black("========"))

		t = uitable.New()
		t.MaxColWidth = 16
		t.Wrap = false
		t.RightAlign(1)
		t.RightAlign(2)
		t.RightAlign(3)

		// t.AddRow(output.TableHeader("Hour (UTC)"), output.TableHeader("Tx Total"), output.TableHeader("Rx Total"), output.TableHeader("Dropped Pkts"))
		t.AddRow(colors.Black("Hour (UTC)"), colors.Black("Tx Total"), colors.Black("Rx Total"), colors.Black("Dropped Pkts"))
		t.AddRow(colors.Black("----------"), colors.Black("--------"), colors.Black("--------"), colors.Black("------------"))
		hour := time.Now().UTC().Hour()
		for h := 0; h < 24; h++ {
			idx := resources.NetworkHistoryHourlyIndex(h)
			if nh, ok := n.Agent.Metrics.NetworkHourlyHistory[idx]; ok {
				switch {
				case hour == h:
					t.AddRow(
						colors.Magenta(idx),
						colors.Magenta(datasize.ByteSize(nh.TxTotalBytes).HumanReadable()),
						colors.Magenta(datasize.ByteSize(nh.RxTotalBytes).HumanReadable()),
						colors.Magenta(fmt.Sprintf("%d", nh.DroppedPkts)),
					)
				case hour < h:
					t.AddRow(
						colors.DarkMagenta(idx),
						colors.DarkMagenta(datasize.ByteSize(nh.TxTotalBytes).HumanReadable()),
						colors.DarkMagenta(datasize.ByteSize(nh.RxTotalBytes).HumanReadable()),
						colors.DarkMagenta(fmt.Sprintf("%d", nh.DroppedPkts)),
					)
				case hour > h:
					t.AddRow(
						idx,
						datasize.ByteSize(nh.TxTotalBytes).HumanReadable(),
						datasize.ByteSize(nh.RxTotalBytes).HumanReadable(),
						nh.DroppedPkts,
					)
				}
			}
		}
		fmt.Println(t)
		fmt.Println()

		// totals & monthly
		fmt.Println("Totals")
		fmt.Printf("%s\n\n", colors.Black("======"))

		t = uitable.New()
		t.MaxColWidth = 18
		t.Wrap = false

		nm := n.Agent.Metrics.NetworkMetrics
		//nm := n.Agent.Metrics.NetDevStats

		// t.AddRow(output.TableHeader("Tx Monthly"), "", output.TableHeader("Rx Monthly"), "")
		t.AddRow(colors.Black("Tx MONTHLY"), "", colors.Black("Rx MONTHLY"), "")
		t.AddRow(colors.Black("----------"), "", colors.Black("----------"), "")
		t.AddRow(colors.Black("Tx Monthly Bytes"), colors.DarkWhite(datasize.ByteSize(nm.TxMonthlyBytes).HumanReadable()),
			colors.Black("Rx Monthly Bytes"), colors.DarkWhite(datasize.ByteSize(nm.RxMonthlyBytes).HumanReadable()))
		t.AddRow(colors.Black("Tx Monthly Pkts"), colors.DarkWhite(fmt.Sprintf("%d", nm.TxMonthlyPkts)),
			colors.Black("Rx Monthly Pkts"), colors.DarkWhite(fmt.Sprintf("%d", nm.RxMonthlyPkts)))
		t.AddRow()
		t.AddRow(colors.Black("Dropped Pkts"), colors.DarkWhite(fmt.Sprintf("%d", nm.DroppedMonthlyPkts)))

		t.AddRow()

		// t.AddRow(output.TableHeader("Tx Totals"), "", output.TableHeader("Rx Totals"), "")
		t.AddRow(colors.Black("Tx TOTALS"), "", colors.Black("Rx TOTALS"), "")
		t.AddRow(colors.Black("---------"), "", colors.Black("---------"), "")
		t.AddRow(colors.Black("Tx Bps (5min avg)"), colors.DarkWhite(fmt.Sprintf("%.3f Kbps", nm.TxBps/1000)),
			colors.Black("Rx Bps (5min avg)"), colors.DarkWhite(fmt.Sprintf("%.3f Kbps", nm.RxBps/1000)))
		t.AddRow(colors.Black("Tx Total Bytes"), colors.DarkWhite(datasize.ByteSize(nm.TxTotalBytes).HumanReadable()),
			colors.Black("Rx Total Bytes"), colors.DarkWhite(datasize.ByteSize(nm.RxTotalBytes).HumanReadable()))
		t.AddRow(colors.Black("Tx Total Pkts"), colors.DarkWhite(fmt.Sprintf("%d", nm.TxTotalPkts)),
			colors.Black("Rx Total Pkts"), colors.DarkWhite(fmt.Sprintf("%d", nm.RxTotalPkts)))
		t.AddRow()
		t.AddRow(colors.Black("Dropped Pkts"), colors.DarkWhite(fmt.Sprintf("%d", nm.DroppedPkts)))

		fmt.Println(t)
		fmt.Println()

		// traffic matrix
		fmt.Println("Traffic Matrix")
		fmt.Printf("%s\n\n", colors.Black("=============="))

		t = uitable.New()
		t.MaxColWidth = 20
		t.Wrap = false
		t.RightAlign(1)
		t.RightAlign(2)
		t.RightAlign(3)
		t.RightAlign(4)
		t.RightAlign(5)

		// t.AddRow(output.TableHeader("Address"), output.TableHeader("Tx Bps"), output.TableHeader("Rx Bps"), output.TableHeader("Tx Total"), output.TableHeader("Rx Total"), output.TableHeader("Drops"))
		t.AddRow(colors.Black("Address"), colors.Black("Tx Bps"), colors.Black("Rx Bps"), colors.Black("Tx Total"), colors.Black("Rx Total"), colors.Black("Drops"))
		t.AddRow(colors.Black("-------"), colors.Black("------"), colors.Black("------"), colors.Black("--------"), colors.Black("--------"), colors.Black("-----"))
		for addr, nt := range n.Agent.Metrics.NetworkTraffic {
			t.AddRow(
				addr,
				nt.TxBps,
				nt.RxBps,
				datasize.ByteSize(nt.TxTotalBytes).HumanReadable(),
				datasize.ByteSize(nt.RxTotalBytes).HumanReadable(),
				nt.DroppedPkts,
			)
		}

		fmt.Println(t)
		fmt.Println()

	*/

	t := uitable.New()
	t.MaxColWidth = 60
	t.Wrap = false

	// var rmConns int32
	// for _, rm := range n.Agent.Metrics.RelayMetrics {
	// 	rmConns += rm.Connections
	// }
	// t.AddRow(colors.Black("Indirect Connections"), colors.DarkWhite(fmt.Sprintf("%d", rmConns)))
	// t.AddRow()

	if n.Agent.Metrics.LastUpdated > 0 {
		tm := time.UnixMilli(n.Agent.Metrics.LastUpdated)
		t.AddRow(colors.Black("Last Updated"), colors.DarkWhite(tm.String()))
	}

	fmt.Println(t)
	fmt.Println()
}
