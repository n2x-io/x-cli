package output

import (
	"fmt"
	"sort"
	"time"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) List(nodes map[string]*topology.Node) {
	output.SectionHeader("Nodes")
	output.TitleT1("Node List")

	t := table.New()
	t.Header(colors.Black("NODE NAME / FQDN"), colors.Black("IPv4"), colors.Black("FLAGS / ENDPOINT"))

	t.SetRowLine("-")

	nodesSort := make([]string, 0)
	for _, n := range nodes {
		nodesSort = append(nodesSort, n.Cfg.NodeName)
	}
	sort.Strings(nodesSort)

	for _, nodeName := range nodesSort {
		n := nodes[nodeName]

		var c1, c2, c3 string
		var nodeName string
		var status string

		if n.Cfg != nil {
			nodeName = colors.Black(output.Fit(n.Cfg.NodeName, 32))

			// hostTitle := fmt.Sprintf(" %s %s:", colors.DarkWhite("└─"), colors.Black("Node Instance"))
			p := fmt.Sprintf("prio-%d", n.Cfg.Priority)
			prio := fmt.Sprintf("%s%s%s", colors.Yellow("["), colors.DarkYellow(p), colors.Yellow("]"))

			var nodeType string
			if n.Agent.CanRelay && !n.Cfg.DisableRelay {
				nodeType = output.StrTier1()
			}

			c3 = fmt.Sprintf("%s %s", prio, nodeType)
		}

		tm := time.UnixMilli(n.LastSeen)
		if time.Since(tm) > nodeTimeout*time.Second {
			status = colors.DarkRed("█")
		} else {
			status = colors.Green("█")
		}

		c1 = fmt.Sprintf("%s %s", status, nodeName)

		if n.Endpoints != nil {
			for _, e := range n.Endpoints {
				fqdn := colors.DarkWhite(output.Fit(e.DNSName+".n2x.local", 32))
				endpoint := output.Fit(e.EndpointID, 20)
				ipv4 := colors.DarkWhite(e.IPv4)

				c1 = fmt.Sprintf("%s\n%s %s", c1, status, fqdn)
				c2 = fmt.Sprintf("%s\n%s", c2, ipv4)
				c3 = fmt.Sprintf("%s\n%s", c3, endpoint)
			}
		}
		t.AddRow(c1, c2, c3)
	}

	t.Render()
	fmt.Println()
}
