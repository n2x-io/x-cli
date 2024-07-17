package cmd

import (
	"github.com/spf13/cobra"
	"n2x.dev/x-cli/pkg/client"
)

var alertsCmd = &cobra.Command{
	Use:   "alert",
	Short: "Alert system",
	Long:  appHeader(`Alert system.`),
}

var alertListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your alerts",
	Long:  appHeader(`List all your alerts.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Alert().List()
	},
}

var alertShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show alert",
	Long:  appHeader(`Show alert details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// a := client.Alert().Show()
		client.Alert().Show()

		// if input.GetConfirm("Add note to this alert?", false) {
		// 	client.Alert().NewNote(a)
		// } else {
		// 	fmt.Println()
		// }
	},
}

var alertDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete alert",
	Long:  appHeader(`Remove alert from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Alert().Delete()
	},
}

func init() {
	alertsCmd.AddCommand(alertListCmd)
	alertsCmd.AddCommand(alertShowCmd)
	alertsCmd.AddCommand(alertDeleteCmd)
}
