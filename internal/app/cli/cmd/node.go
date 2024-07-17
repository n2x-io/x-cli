package cmd

import (
	"github.com/spf13/cobra"
	"n2x.dev/x-cli/pkg/client"
	"n2x.dev/x-cli/pkg/vars"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Network node operations",
	Long:  appHeader(`Network node operations for network administrators.`),
}

var nodeAddNodeCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new node",
	Long:  appHeader(`Register a new node.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().AddNode()
	},
}

/*
var nodeGetInstallationWebhookCmd = &cobra.Command{
	Use:   "get-magic-link",
	Short: "Generate magic link to setup a new node (linux only)",
	Long:  appHeader(`Generate magic link to setup a new node (linux only).`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().GetInstallationWebhook()
	},
}
*/

var nodeListByTenant, nodeListBySubnet bool

var nodeListCmd = &cobra.Command{
	Use:   "list --by-tenant | --by-subnet",
	Short: "List nodes",
	Long:  appHeader(`List nodes.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if nodeListBySubnet {
			client.Node().ListBySubnet()
		} else if nodeListByTenant {
			client.Node().ListByTenant()
		}
	},
}

var nodeShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show node",
	Long:  appHeader(`Show node details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().Show()
	},
}

var nodeConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect (move) node to a subnet",
	Long:  appHeader(`Connect (move) node to a subnet.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().Connect()
	},
}

var nodeDisconnectCmd = &cobra.Command{
	Use:   "disable-networking",
	Short: "Disable n2x networking and disconnect node from a subnet",
	Long:  appHeader(`Disable n2x networking and disconnect node from a subnet.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().Disconnect()
	},
}

var nodeDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove node from database",
	Long:  appHeader(`Remove node from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().Delete()
	},
}

var nodeMetricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Show detailed metrics",
	Long:  appHeader(`Show detailed metrics.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().Metrics()
	},
}

// var nodeResetTrafficMatrixCmd = &cobra.Command{
// 	Use:   "reset-traffic-matrix",
// 	Short: "Reset traffic matrix metrics",
// 	Long:  appHeader(`Reset traffic matrix metrics.`),
// 	Args:  cobra.NoArgs,
//  PreRun: func(cmd *cobra.Command, args []string) {
// 	    preflight()
//  },
// 	Run: func(cmd *cobra.Command, args []string) {
// 		client.Node().ResetNetworkTraffic()
// 	},
// }

var nodeShowEndpointCmd = &cobra.Command{
	Use:   "show-endpoint",
	Short: "Show network endpoint details",
	Long:  appHeader(`Show network endpoint details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().ShowEndpoint()
	},
}

var nodeDeleteEndpointCmd = &cobra.Command{
	Use:   "delete-endpoint",
	Short: "Delete network endpoint",
	Long:  appHeader(`Remove network endpoint from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().DeleteEndpoint()
	},
}

func init() {
	nodeCmd.AddCommand(nodeAddNodeCmd)
	// nodeCmd.AddCommand(nodeGetInstallationWebhookCmd)
	nodeCmd.AddCommand(nodeListCmd)
	nodeCmd.AddCommand(nodeShowCmd)
	nodeCmd.AddCommand(nodeConnectCmd)
	nodeCmd.AddCommand(nodeDisconnectCmd)
	nodeCmd.AddCommand(nodeDeleteCmd)
	nodeCmd.AddCommand(nodeMetricsCmd)
	// nodeCmd.AddCommand(nodeResetTrafficMatrixCmd)
	nodeCmd.AddCommand(nodeShowEndpointCmd)
	nodeCmd.AddCommand(nodeDeleteEndpointCmd)

	nodeCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	nodeCmd.PersistentFlags().StringVarP(&vars.NetID, "network", "n", "", "network identifier")
	nodeCmd.PersistentFlags().StringVarP(&vars.SubnetID, "subnet", "s", "", "subnet identifier")
	nodeCmd.PersistentFlags().StringVarP(&vars.NodeID, "node", "x", "", "node identifier")

	nodeListCmd.Flags().BoolVar(&nodeListByTenant, "by-tenant", false, "list nodes by tenant")
	nodeListCmd.Flags().BoolVar(&nodeListBySubnet, "by-subnet", false, "list nodes by subnet")
	nodeListCmd.MarkFlagsOneRequired("by-tenant", "by-subnet")
}
