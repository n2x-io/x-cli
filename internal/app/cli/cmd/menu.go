package cmd

import (
	"github.com/spf13/cobra"
)

/*
var iamCmd = &cobra.Command{
	Use:   "iam",
	Short: "Users administration and access management (RBAC)",
	Long:  appHeader(`Users administration and access management (RBAC).`),
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User settings",
	Long:  appHeader(`User settings.`),
}
*/

var opsCmd = &cobra.Command{
	Use:   "ops",
	Short: "Automation and GitOps: projects / workflows / logs",
	Long:  appHeader(`Automation and workflows management commands.`),
}

// var auditCmd = &cobra.Command{
// 	Use:   "audit",
// 	Short: "Audit commands",
// 	Long:  appHeader(`Platform audit commands.`),
// }

// var statusCmd = &cobra.Command{
// 	Use:   "status",
// 	Short: "Platform status report",
// 	Long:  appHeader(`Platform status report.`),
// }

// var runtimeCmd = &cobra.Command{
// 	Use:   "runtime",
// 	Short: "Runtime configuration administration",
// 	Long:  appHeader(`Runtime configuration administration.`),
// }

func init() {
	rootCmd.AddCommand(accountCmd)
	// rootCmd.AddCommand(iamCmd)
	rootCmd.AddCommand(authCmd)
	// rootCmd.AddCommand(userCmd)
	rootCmd.AddCommand(alertsCmd)
	rootCmd.AddCommand(tenantCmd)
	rootCmd.AddCommand(networkCmd)
	rootCmd.AddCommand(subnetCmd)
	rootCmd.AddCommand(nodeCmd)
	rootCmd.AddCommand(vsCmd)
	rootCmd.AddCommand(networkPolicyCmd)
	rootCmd.AddCommand(k8sCmd)
	// rootCmd.AddCommand(serviceCmd)
	rootCmd.AddCommand(opsCmd)
	// rootCmd.AddCommand(execCmd)
	// rootCmd.AddCommand(transferCmd)
	// rootCmd.AddCommand(portFwdCmd)
	// rootCmd.AddCommand(auditCmd)
	// rootCmd.AddCommand(supportCmd)
	// rootCmd.AddCommand(statusCmd)
	// rootCmd.AddCommand(runtimeCmd)
	rootCmd.AddCommand(versionCmd)
	// rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(completionCmd)
}
