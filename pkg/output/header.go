package output

import (
	"fmt"

	"n2x.dev/x-lib/pkg/utils/colors"
)

func AppHeader(output bool) string {
	return logo(output)
}

func logo(output bool) string {
	website := fmt.Sprintf("%s %s", colors.Black("Main Website: "), colors.DarkWhite("https://n2x.io"))
	docsite := fmt.Sprintf("%s %s", colors.Black("Documentation:"), colors.DarkWhite("https://n2x.io/docs"))
	// discord := fmt.Sprintf("%s %s", colors.Black("Discord Server:"), colors.DarkWhite("https://n2x.io/discord"))

	if output {
		l1 := fmt.Sprintf("  ■   ▄  ▄▄▄▄ " + colors.Green("│") + "\n")
		l2 := fmt.Sprintf("■  ██    ▀  ▄ "+colors.Green("│")+" %s\n", website)
		l3 := fmt.Sprintf("  ▀   ■  ▀▀▀▀ "+colors.Green("│")+" %s\n", docsite)
		return fmt.Sprintf("%s%s%s\n", l1, l2, l3)
	}

	fmt.Printf("  ■   ▄  ▄▄▄▄ " + colors.Green("│") + "\n")
	fmt.Printf("■  ██    ▀  ▄ "+colors.Green("│")+" %s\n", website)
	fmt.Printf("  ▀   ■  ▀▀▀▀ "+colors.Green("│")+" %s\n", docsite)
	fmt.Println()

	return ""
}
