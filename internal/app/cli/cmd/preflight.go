package cmd

import (
	"fmt"

	"n2x.dev/x-cli/internal/app/cli/auth/login"
	"n2x.dev/x-cli/pkg/client"
	"n2x.dev/x-lib/pkg/utils/colors"
	"n2x.dev/x-lib/pkg/version"
)

func header() {
	fmt.Println(colors.Black(version.CLI_NAME + " " + version.GetVersion()))

	fmt.Println()

	fmt.Printf("%s is a CLI to control the %s.\n\n",
		version.CLI_NAME, version.PLATFORM_NAME)

	fmt.Printf("Find more information at %s\n\n", colors.White("https://n2x.io/docs"))

	// output.AppHeader(false)
}

func appHeader(str string) string {
	h1 := colors.Black(version.CLI_NAME + " " + version.GetVersion())
	// h2 := output.AppHeader(true)

	// return fmt.Sprintf("%s\n%s%s", h1, h2, str)
	return fmt.Sprintf("%s\n\n%s", h1, str)
}

func preflightNoLogin() {
	header()

	// if !isConfigured {
	// 	notConfigured()
	// 	os.Exit(0)
	// }

	// if _, err := auth.GetAccountID(); err != nil {
	// 	status.Error(fmt.Errorf("missing accountID"), "Unable to get account")
	// }
}

func preflight() {
	header()

	autoLogin()

	// if isConfigured {
	// 	// silent auto-login
	// 	autoLogin()
	// } else {
	// 	notConfigured()
	// 	// os.Exit(0)
	// }
}

func autoLogin() {
	if client.Auth().LoginRequired() {
		client.Auth().OTPSignin(login.NewRequestWithOTP(), true)
		// client.Auth().LoginWithToken(login.NewRequestWithToken(), true)
	}
}

/*
func notConfigured() {
	msg.Warn("Configuration not detected")

	fmt.Printf("%s\n", colors.DarkBlue("_"))
	cmd := colors.DarkWhite(fmt.Sprintf("%s setup", version.CLI_NAME))
	q := colors.DarkBlue("'")
	msg := fmt.Sprintf("%s %s%s%s", colors.Black("Please configure the client with"), q, cmd, q)
	fmt.Printf("%s %s\n\n", colors.Cyan("🢂"), msg)
}
*/
