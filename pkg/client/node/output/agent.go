package output

import (
	"fmt"
	"os"
	"strings"
	"time"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
	"n2x.dev/x-lib/pkg/utils/msg"
)

const nodeTimeout = 960

func agentInfo(n *topology.Node) {
	ncfg := n.Cfg
	if ncfg == nil {
		msg.Alert("No data found, node is not configured yet")
		os.Exit(0)
	}

	t := table.New()

	// t.AddRow(colors.Black("Account ID"), colors.DarkWhite(n.AccountID))
	t.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(n.TenantID))
	if len(n.Cfg.NetID) > 0 {
		t.AddRow(colors.Black("Network ID"), colors.DarkWhite(n.Cfg.NetID))
	}
	if len(n.Cfg.SubnetID) > 0 {
		t.AddRow(colors.Black("Subnet ID"), colors.DarkWhite(n.Cfg.SubnetID))
	}
	t.AddRow(colors.Black("Node ID"), colors.DarkWhite(n.NodeID))
	// t.AddRow()
	t.AddRow(colors.Black("Node Name"), colors.DarkWhite(ncfg.NodeName))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(ncfg.Description))

	var status string
	tm := time.UnixMilli(n.LastSeen)
	if time.Since(tm) > nodeTimeout*time.Second {
		status = output.StrOffline()
	} else {
		status = output.StrOnline()
	}
	t.AddRow(colors.Black("Status"), status)

	t.Render()
	fmt.Println()

	fmt.Print(colors.Black("-----NODE AUTHORIZATION TOKEN-----\n"))
	fmt.Printf("%s\n", colors.DarkWhite(n.NodeToken))
	fmt.Print(colors.Black("-----NODE AUTHORIZATION TOKEN-----\n"))
	fmt.Println()

	na := n.Agent
	if na == nil {
		// msg.Alert("No data found, agent is not registered yet")
		// os.Exit(0)
		return
	}

	if na.Metrics == nil {
		// fmt.Printf("[%s]\n\n", colors.Black("not enough data yet"))
		// os.Exit(0)
		return
	}

	if na.Metrics.HostMetrics == nil {
		// fmt.Printf("[%s]\n\n", colors.Black("not enough data yet"))
		// os.Exit(0)
		return
	}

	t = table.New()

	t.AddRow(colors.Black("OS"), colors.DarkWhite(strings.Title(na.Metrics.HostMetrics.OS)))

	if len(na.Metrics.HostMetrics.Uptime) > 0 {
		t.AddRow(colors.Black("Uptime"), colors.DarkWhite(na.Metrics.HostMetrics.Uptime))
	}

	var autoUpdate, maintSched string
	if ncfg.Maintenance != nil {
		if ncfg.Maintenance.AutoUpdate {
			autoUpdate = output.StrEnabled("auto-update")
		} else {
			autoUpdate = output.StrDisabled("auto-update")
		}
		if ncfg.Maintenance.Schedule != nil {
			maintHour := fmt.Sprintf("%02d", ncfg.Maintenance.Schedule.Hour)
			maintMin := fmt.Sprintf("%02d", ncfg.Maintenance.Schedule.Minute)
			maintSched = colors.DarkGreen(fmt.Sprintf("%s:%s", maintHour, maintMin))
			maintSched = fmt.Sprintf("%s [%s]", colors.DarkWhite("Scheduled"), maintSched)
		} else {
			maintSched = colors.DarkWhite("Not scheduled")
		}
	}
	maint := fmt.Sprintf("%s %s", autoUpdate, maintSched)
	t.AddRow(colors.Black("Maintenance"), maint)

	var disableExec, disableTransfer, disablePortFwd, disableOps string
	if ncfg.Management != nil {
		if ncfg.Management.DisableExec {
			disableExec = output.StrDisabled("exec")
		} else {
			disableExec = output.StrEnabled("exec")
		}
		if ncfg.Management.DisableTransfer {
			disableTransfer = output.StrDisabled("transfer")
		} else {
			disableTransfer = output.StrEnabled("transfer")
		}
		if ncfg.Management.DisablePortForwarding {
			disablePortFwd = output.StrDisabled("portForward")
		} else {
			disablePortFwd = output.StrEnabled("portForward")
		}
		if ncfg.Management.DisableOps {
			disableOps = output.StrDisabled("workflows")
		} else {
			disableOps = output.StrEnabled("workflows")
		}
	}
	opts := fmt.Sprintf("%s %s %s %s", disableExec, disableTransfer, disablePortFwd, disableOps)
	t.AddRow(colors.Black("Options"), opts)

	// t.AddRow()

	var extIP string
	if len(na.ExternalIPv4) > 0 {
		extIP = na.ExternalIPv4
	} else {
		extIP = "n/a"
	}
	t.AddRow(colors.Black("External IP"), colors.DarkWhite(extIP))
	t.AddRow(colors.Black("Port"), colors.DarkWhite(fmt.Sprintf("%d", na.Port)))

	var k8sGw string
	if n.Type == topology.NodeType_K8S_GATEWAY {
		k8sGw = output.StrEnabled("k8sGw")
	}

	dnsPort := fmt.Sprintf("udp/%d", na.DNSPort)

	var relay string
	if na.CanRelay && !ncfg.DisableRelay {
		relay = output.StrTier1()
	}
	prio := fmt.Sprintf("[%s-%s]", colors.DarkWhite("priority"), colors.Yellow(fmt.Sprintf("%d", ncfg.Priority)))
	routing := fmt.Sprintf("%s %s %s", prio, relay, k8sGw)
	t.AddRow(colors.Black("DNS Port"), colors.DarkWhite(dnsPort))
	t.AddRow(colors.Black("Routing"), routing)

	// t.AddRow()

	hostResources := fmt.Sprintf("Load %s | RAM %s | Disk %s",
		colors.DarkWhite(fmt.Sprintf("%f", na.Metrics.HostMetrics.LoadAvg)),
		output.PercentBar(int(na.Metrics.HostMetrics.MemoryUsage)),
		output.PercentBar(int(na.Metrics.HostMetrics.DiskUsage)))
	t.AddRow(colors.Black("Resources"), hostResources)

	// t.AddRow(colors.Black("Memory"), output.PercentBar(int(na.Metrics.HostMetrics.MemoryUsage)))
	// t.AddRow(colors.Black("Load Average"), colors.DarkWhite(fmt.Sprintf("%f", na.Metrics.HostMetrics.LoadAvg)))
	// t.AddRow(colors.Black("Disk"), output.PercentBar(int(na.Metrics.HostMetrics.DiskUsage)))

	t.Render()
	fmt.Println()
}
