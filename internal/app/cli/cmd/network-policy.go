package cmd

import (
	"github.com/spf13/cobra"
	"n2x.dev/x-cli/pkg/client"
	"n2x.dev/x-cli/pkg/vars"
)

var networkPolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Security policies administration",
	Long:  appHeader(`Security policies administration.`),
}

var networkPolicyShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show security policy",
	Long:  appHeader(`Show security policy details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().Show()
	},
}

/*
var networkPolicyImportCmd = &cobra.Command{
	Use:   "import -f <yamlFile>",
	Short: "Import security policy from YAML file",
	Long:  appHeader(`Import security policy from YAML file.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().Import(vars.YAMLFile)
	},
}

var networkPolicyExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export security policy to standard output",
	Long:  appHeader(`Export security policy  to standard output.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().Export()
	},
}
*/

var networkPolicySetDefaultCmd = &cobra.Command{
	Use:   "set-default",
	Short: "Set default policy",
	Long:  appHeader(`Set default policy.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().SetDefault()
	},
}

/*
var networkPolicySetRuleCmd = &cobra.Command{
	Use:   "set-rule",
	Short: "Set security policy rule",
	Long:  appHeader(`Set security policy rule.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().SetRule()
	},
}

var networkPolicyUnsetRuleCmd = &cobra.Command{
	Use:   "unset-rule",
	Short: "Unset security policy rule",
	Long:  appHeader(`Unset security policy rule.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().UnsetRule()
	},
}
*/

var networkPolicyNewNetworkFilterCmd = &cobra.Command{
	Use:   "add-rule",
	Short: "Add security policy rule",
	Long:  appHeader(`Add security policy rule.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().NewNetworkFilter()
	},
}

var networkPolicyUpdateNetworkFilterCmd = &cobra.Command{
	Use:   "edit-rule",
	Short: "Edit security policy rule",
	Long:  appHeader(`Edit security policy rule.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().UpdateNetworkFilter()
	},
}

var networkPolicyDeleteNetworkFilterCmd = &cobra.Command{
	Use:   "delete-rule",
	Short: "Delete security policy rule",
	Long:  appHeader(`Delete security policy rule.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().DeleteNetworkFilter()
	},
}

var networkPolicyDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete security policy (all rules)",
	Long:  appHeader(`Remove security policy (all rules) from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().Delete()
	},
}

func init() {
	networkPolicyCmd.AddCommand(networkPolicyShowCmd)
	// networkPolicyCmd.AddCommand(networkPolicyImportCmd)
	// networkPolicyCmd.AddCommand(networkPolicyExportCmd)
	networkPolicyCmd.AddCommand(networkPolicySetDefaultCmd)
	// networkPolicyCmd.AddCommand(networkPolicySetRuleCmd)
	// networkPolicyCmd.AddCommand(networkPolicyUnsetRuleCmd)
	networkPolicyCmd.AddCommand(networkPolicyNewNetworkFilterCmd)
	networkPolicyCmd.AddCommand(networkPolicyUpdateNetworkFilterCmd)
	networkPolicyCmd.AddCommand(networkPolicyDeleteNetworkFilterCmd)
	networkPolicyCmd.AddCommand(networkPolicyDeleteCmd)

	networkPolicyCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	networkPolicyCmd.PersistentFlags().StringVarP(&vars.NetID, "network", "n", "", "network identifier")
	networkPolicyCmd.PersistentFlags().StringVarP(&vars.SubnetID, "subnet", "s", "", "subnet identifier")

	// networkPolicyImportCmd.Flags().StringVarP(&vars.YAMLFile, "yamlFile", "f", "", "yaml workflow file")
	// networkPolicyImportCmd.MarkFlagRequired("yamlFile")
}
