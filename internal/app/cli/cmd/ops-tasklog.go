package cmd

import (
	"github.com/spf13/cobra"
	"n2x.dev/x-cli/pkg/client"
	"n2x.dev/x-cli/pkg/vars"
)

// opsTaskLogsCmd represents the opsTaskLogs command
var opsTaskLogsCmd = &cobra.Command{
	Use:   "tasklog",
	Short: "Workflow tasklog administration",
	Long:  appHeader(`Workflow tasklog administration.`),
}

// opsTaskLogsListCmd represents the ops/operations list verb
var opsTaskLogsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasklogs",
	Long:  appHeader(`List all tasklogs.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.TaskLog().List()
	},
}

// opsTaskLogsShowCmd represents the ops/operations get verb
var opsTaskLogsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show log",
	Long:  appHeader(`Show log details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.TaskLog().Show()
	},
}

// opsTaskLogsDeleteCmd represents the ops/operations delete verb
var opsTaskLogsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete log",
	Long:  appHeader(`Remove log from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.TaskLog().Delete()
	},
}

func init() {
	opsCmd.AddCommand(opsTaskLogsCmd)
	opsTaskLogsCmd.AddCommand(opsTaskLogsListCmd)
	opsTaskLogsCmd.AddCommand(opsTaskLogsShowCmd)
	opsTaskLogsCmd.AddCommand(opsTaskLogsDeleteCmd)

	opsTaskLogsCmd.PersistentFlags().StringVarP(&vars.ProjectID, "project", "p", "", "project identifier")
	opsTaskLogsCmd.PersistentFlags().StringVarP(&vars.WorkflowID, "workflow", "w", "", "workflow identifier")
	opsTaskLogsCmd.PersistentFlags().StringVarP(&vars.TaskLogID, "tasklog", "t", "", "tasklog identifier")
}
