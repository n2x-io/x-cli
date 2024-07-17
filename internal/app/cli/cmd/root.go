package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"n2x.dev/x-cli/pkg/config"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/update"
	"n2x.dev/x-lib/pkg/utils"
	"n2x.dev/x-lib/pkg/utils/colors"
	"n2x.dev/x-lib/pkg/utils/msg"
	"n2x.dev/x-lib/pkg/version"
)

var cfgFile string
var newVersionAvailable, isConfigured bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   version.CLI_NAME,
	Short: version.CLI_NAME + " is a CLI to control the " + version.PLATFORM_NAME,
	Long: version.CLI_NAME + ` is a CLI to control the ` +
		version.PLATFORM_NAME + `.

Find support and more information:

  Project Website:     ` + version.N2X_URL + `
  Documentation:       ` + version.N2X_DOC_URL + `
  Join us on Discord:  ` + version.DISCORD_URL,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func Init() {
	if err := consoleInit(); err != nil {
		msg.Error(err)
		os.Exit(1)
	}

	initConfig()

	rootCmd.Long = appHeader(rootCmd.Long)

	cobra.EnableCommandSorting = false

	cobra.OnInitialize()

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", defaultConfigFileHelp())
	// rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "output format (table, json)")
	// rootCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	execute()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// flag.StringVar(&vars.AccountID, "a", "", "account identifier")
	flag.StringVar(&cfgFile, "config", "", defaultConfigFileHelp())
	flag.Parse()

	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	cfgFile = getConfigFile(cfgFile)

	viper.SetConfigFile(cfgFile)

	viper.SetEnvPrefix("mm") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv() // read in environment variables that match

	if !utils.FileExists(cfgFile) {
		config.SetDefaults()
	}

	if utils.FileExists(cfgFile) {
		if err := viper.ReadInConfig(); err != nil {
			status.Error(err, "Unable to read config file")
		}
		isConfigured = true
		// msg.Debugf("Using configuration file: %v", viper.ConfigFileUsed())
	}

	config.Init()
}

func execute() {
	go checkVersionUpdate()

	if err := rootCmd.Execute(); err != nil {
		msg.Error(errors.Cause(err))
		os.Exit(1)
	}

	if newVersionAvailable {
		fmt.Printf("%s\n", colors.DarkBlue("_"))
		cmd := colors.DarkWhite(fmt.Sprintf("%s version update", version.CLI_NAME))
		q := colors.DarkBlue("'")
		msg := fmt.Sprintf("%s %s%s%s", colors.Black("New version available, please update with"), q, cmd, q)
		fmt.Printf("%s %s\n\n", colors.Cyan("🢂"), msg)
	}
}

func checkVersionUpdate() {
	newVersionAvailable, _ = update.IsBinaryOutdated(version.CLI_NAME)
}
