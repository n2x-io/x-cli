package cmd

import (
	"github.com/spf13/cobra"
	"n2x.dev/x-cli/pkg/client"
	"n2x.dev/x-cli/pkg/vars"
)

var opsWorkflowsCmd = &cobra.Command{
	Use:   "workflow",
	Short: "Automation workflows administration",
	Long:  appHeader(`Automation workflows administration.`),
}

var opsWorkflowsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List workflows",
	Long:  appHeader(`List all workflows.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Workflow().List()
	},
}

var opsWorkflowsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show workflow",
	Long:  appHeader(`Show workflow details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Workflow().Show()
	},
}

var opsWorkflowsCreateCmd = &cobra.Command{
	Use:   "create -f <yamlFile>",
	Short: "Create workflow from YAML file",
	Long:  appHeader(`Create workflow from YAML file.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Workflow().Create(vars.YAMLFile)
	},
}

var opsWorkflowsUpdateCmd = &cobra.Command{
	Use:   "update -f <yamlFile>",
	Short: "Update workflow from YAML file",
	Long:  appHeader(`Update workflow from YAML file.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Workflow().Update(vars.YAMLFile)
	},
}

var opsWorkflowsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove workflow",
	Long:  appHeader(`Remove workflow from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Workflow().Delete()
	},
}

var opsWorkflowsEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable workflow",
	Long:  appHeader(`Enable workflow.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Workflow().Enable()
	},
}

var opsWorkflowsDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable workflow",
	Long:  appHeader(`Disable workflow.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Workflow().Disable()
	},
}

func init() {
	opsCmd.AddCommand(opsWorkflowsCmd)
	opsWorkflowsCmd.AddCommand(opsWorkflowsListCmd)
	opsWorkflowsCmd.AddCommand(opsWorkflowsShowCmd)
	opsWorkflowsCmd.AddCommand(opsWorkflowsCreateCmd)
	opsWorkflowsCmd.AddCommand(opsWorkflowsUpdateCmd)
	opsWorkflowsCmd.AddCommand(opsWorkflowsDeleteCmd)
	opsWorkflowsCmd.AddCommand(opsWorkflowsEnableCmd)
	opsWorkflowsCmd.AddCommand(opsWorkflowsDisableCmd)

	opsWorkflowsCmd.PersistentFlags().StringVarP(&vars.ProjectID, "project", "p", "", "project identifier")
	opsWorkflowsCmd.PersistentFlags().StringVarP(&vars.WorkflowID, "workflow", "w", "", "workflow identifier")

	opsWorkflowsCreateCmd.Flags().StringVarP(&vars.YAMLFile, "yamlFile", "f", "", "yaml workflow file")
	opsWorkflowsCreateCmd.MarkFlagRequired("yamlFile")
	opsWorkflowsUpdateCmd.Flags().StringVarP(&vars.YAMLFile, "yamlFile", "f", "", "yaml workflow file")
	opsWorkflowsUpdateCmd.MarkFlagRequired("yamlFile")
}
