package cmd

import (
	"github.com/spf13/cobra"
	"n2x.dev/x-cli/pkg/client"
	"n2x.dev/x-cli/pkg/vars"
)

var opsProjectsCmd = &cobra.Command{
	Use:   "project",
	Short: "Workflow projects administration",
	Long:  appHeader(`Workflow projects administration.`),
}

var opsProjectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Long:  appHeader(`List all projects.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Project().List()
	},
}

var opsProjectsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show project",
	Long:  appHeader(`Show project details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Project().Show()
	},
}

var opsProjectsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project",
	Long:  appHeader(`Create a new project interactively.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Project().Create()
	},
}

var opsProjectsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a project",
	Long:  appHeader(`Update a project interactively.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Project().Update()
	},
}

var opsProjectsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove project",
	Long:  appHeader(`Remove project from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Project().Delete()
	},
}

func init() {
	opsCmd.AddCommand(opsProjectsCmd)
	opsProjectsCmd.AddCommand(opsProjectsListCmd)
	opsProjectsCmd.AddCommand(opsProjectsShowCmd)
	opsProjectsCmd.AddCommand(opsProjectsCreateCmd)
	opsProjectsCmd.AddCommand(opsProjectsUpdateCmd)
	opsProjectsCmd.AddCommand(opsProjectsDeleteCmd)

	opsProjectsCmd.PersistentFlags().StringVarP(&vars.ProjectID, "project", "p", "", "project identifier")
}
