package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"n2x.dev/x-lib/pkg/version"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate completion script",
	Long:  appHeader("Generate completion script."),
	// DisableFlagsInUseLine: true,
}

var completionBashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generate the autocompletion script for bash",
	Long: appHeader(fmt.Sprintf(`Generate the autocompletion script for bash.

To load completions:

  $ source <(%[1]s completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ %[1]s completion bash > /etc/bash_completion.d/%[1]s
  # macOS:
  $ %[1]s completion bash > /usr/local/etc/bash_completion.d/%[1]s
`, version.CLI_NAME)),
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Root().GenBashCompletion(os.Stdout)
	},
}

var completionZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generate the autocompletion script for zsh",
	Long: appHeader(fmt.Sprintf(`Generate the autocompletion script for zsh.

To load completions:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ %[1]s completion zsh > "${fpath[1]}/_%[1]s"

  # You will need to start a new shell for this setup to take effect.
`, version.CLI_NAME)),
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Root().GenZshCompletion(os.Stdout)
	},
}

var completionFishCmd = &cobra.Command{
	Use:   "fish",
	Short: "Generate the autocompletion script for fish",
	Long: appHeader(fmt.Sprintf(`Generate the autocompletion script for fish.

To load completions:

  $ %[1]s completion fish | source

  # To load completions for each session, execute once:
  $ %[1]s completion fish > ~/.config/fish/completions/%[1]s.fish
`, version.CLI_NAME)),
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Root().GenFishCompletion(os.Stdout, true)
	},
}

var completionPowerShellCmd = &cobra.Command{
	Use:   "powershell",
	Short: "Generate the autocompletion script for powershell",
	Long: appHeader(fmt.Sprintf(`Generate the autocompletion script for powershell.

To load completions:

  PS> %[1]s completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> %[1]s completion powershell > %[1]s.ps1
  # and source this file from your PowerShell profile.
`, version.CLI_NAME)),
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
	},
}

func init() {
	completionCmd.AddCommand(completionBashCmd)
	completionCmd.AddCommand(completionZshCmd)
	completionCmd.AddCommand(completionFishCmd)
	completionCmd.AddCommand(completionPowerShellCmd)
}
