package cmd

import (
	"github.com/spf13/cobra"
	"n2x.dev/x-cli/pkg/client"
	"n2x.dev/x-cli/pkg/vars"
)

// tenantCmd represents the tenant command
var tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "Tenants administration",
	Long:  appHeader(`Tenants administration.`),
}

// tenantListCmd represents the admin/tenants list verb
var tenantListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tenants",
	Long:  appHeader(`List all tenants.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Tenant().List()
	},
}

// tenantShowCmd represents the admin/tenants get verb
var tenantShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show tenant",
	Long:  appHeader(`Show tenant details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Tenant().Show()
	},
}

var tenantNewCmd = &cobra.Command{
	Use:   "create",
	Short: "Create tenant",
	Long:  appHeader(`Create tenant interactively.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Tenant().New()
	},
}

var tenantUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update tenant details",
	Long:  appHeader(`Update tenant details interactively.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Tenant().Update()
	},
}

// tenantDeleteCmd represents the admin/tenants delete verb
var tenantDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove tenant",
	Long:  appHeader(`Remove tenant from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Tenant().Delete()
	},
}

func init() {
	tenantCmd.AddCommand(tenantListCmd)
	tenantCmd.AddCommand(tenantShowCmd)
	tenantCmd.AddCommand(tenantNewCmd)
	tenantCmd.AddCommand(tenantUpdateCmd)
	tenantCmd.AddCommand(tenantDeleteCmd)

	tenantCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
}
