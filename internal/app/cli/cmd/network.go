package cmd

import (
	"github.com/spf13/cobra"
	"n2x.dev/x-cli/pkg/client"
	"n2x.dev/x-cli/pkg/vars"
)

// networkCmd represents the network command
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Networks administration",
	Long:  appHeader(`Networks administration.`),
}

// networkListCmd represents the admin/networks list verb
var networkListCmd = &cobra.Command{
	Use:   "list",
	Short: "List networks",
	Long:  appHeader(`List all networks.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Network().List()
	},
}

// networkShowCmd represents the admin/networks get verb
var networkShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show network",
	Long:  appHeader(`Show network details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Network().Show()
	},
}

// networkNewCmd represents the network create verb
var networkNewCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a network",
	Long:  appHeader(`Create a network interactively.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Network().New()
	},
}

// networkUpdateCmd represents the network update verb
var networkUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a network",
	Long:  appHeader(`Update a network interactively.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Network().Update()
	},
}

// networkDeleteCmd represents the admin/tenants delete verb
var networkDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove network",
	Long:  appHeader(`Remove network from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Network().Delete()
	},
}

func init() {
	networkCmd.AddCommand(networkListCmd)
	networkCmd.AddCommand(networkShowCmd)
	networkCmd.AddCommand(networkNewCmd)
	networkCmd.AddCommand(networkUpdateCmd)
	networkCmd.AddCommand(networkDeleteCmd)

	networkCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	networkCmd.PersistentFlags().StringVarP(&vars.NetID, "network", "n", "", "network identifier")
}
