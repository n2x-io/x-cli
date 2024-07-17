package cmd

import (
	"github.com/spf13/cobra"
	"n2x.dev/x-cli/pkg/client"
	"n2x.dev/x-cli/pkg/vars"
)

var vsCmd = &cobra.Command{
	Use:   "vs",
	Short: "Virtual Servers (lb/ha services)",
	Long:  appHeader(`Virtual Servers (lb/ha services).`),
}

var vsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List virtual servers",
	Long:  appHeader(`List all virtual servers.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.VS().List()
	},
}

var vsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show virtual server",
	Long:  appHeader(`Show virtual server details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.VS().Show()
	},
}

var vsNewCmd = &cobra.Command{
	Use:   "create",
	Short: "Create virtual server",
	Long:  appHeader(`Create virtual server interactively.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.VS().New()
	},
}

var vsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update virtual server",
	Long:  appHeader(`Update virtual server interactively.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.VS().Update()
	},
}

var vsAddAppSvcCmd = &cobra.Command{
	Use:   "add-node-svc",
	Short: "Add node app service to VS high-availability group",
	Long:  appHeader(`Add node app service to VS high-availability group.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.VS().AddVSAppSvc()
	},
}

var vsDeleteAppSvcCmd = &cobra.Command{
	Use:   "remove-node-svc",
	Short: "Remove node app service from VS high-availability group",
	Long:  appHeader(`Remove node app service from VS high-availability group.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.VS().DeleteVSAppSvc()
	},
}

var vsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove virtual server",
	Long:  appHeader(`Remove virtual server.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.VS().Delete()
	},
}

func init() {
	vsCmd.AddCommand(vsListCmd)
	vsCmd.AddCommand(vsShowCmd)
	vsCmd.AddCommand(vsNewCmd)
	vsCmd.AddCommand(vsUpdateCmd)
	vsCmd.AddCommand(vsAddAppSvcCmd)
	vsCmd.AddCommand(vsDeleteAppSvcCmd)
	vsCmd.AddCommand(vsDeleteCmd)

	vsCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	vsCmd.PersistentFlags().StringVarP(&vars.VSID, "vs", "v", "", "vs identifier")
}
