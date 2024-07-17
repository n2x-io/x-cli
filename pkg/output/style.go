package output

import (
	"fmt"
	"strings"
	"time"

	"n2x.dev/x-lib/pkg/utils/colors"
)

func SectionHeader(s string) {
	n := len(s)
	sep := strings.Repeat("─", 5)
	spc := strings.Repeat(" ", 80-n-10)
	fmt.Printf("%s%s %s %s\n", spc, colors.Black(sep), colors.White(s), colors.Black("≡"))
}

func Header(s string) {
	n := len(s)
	sep := strings.Repeat("-", n)
	fmt.Println(colors.White(s))
	fmt.Printf("%s\n", colors.Black(sep))

	// fmt.Printf("\n%s\n\n", colors.InvertedMagenta(s))
}

func TitleT1(s string) {
	n := len(s)
	// sep := strings.Repeat("=", n)
	sep := strings.Repeat("═", n)
	fmt.Printf("%s\n", colors.Black(sep))
	fmt.Println(colors.White(s))
	fmt.Printf("%s\n\n", colors.Black(sep))
}

func TitleT2(s string) {
	n := len(s)
	// sep := strings.Repeat("-", n)
	sep := strings.Repeat("─", n)
	fmt.Printf("%s\n", colors.Black(sep))
	fmt.Println(colors.White(s))
	fmt.Printf("%s\n\n", colors.Black(sep))
}

func SubTitleT1(s string) {
	n := len(s)
	// sep := strings.Repeat("=", n)
	sep := strings.Repeat("═", n)
	fmt.Println(colors.White(s))
	fmt.Printf("%s\n\n", colors.Black(sep))
}

func SubTitleT2(s string) {
	n := len(s)
	// sep := strings.Repeat("-", n)
	sep := strings.Repeat("─", n)
	fmt.Println(colors.White(s))
	fmt.Printf("%s\n\n", colors.Black(sep))
}

func BoxT1(s string) {
	n := len(s) + 2
	sep := strings.Repeat("─", n)
	fmt.Printf("┌%s┐\n", sep)
	fmt.Printf("│ %s │\n", colors.Black(s))
	fmt.Printf("└%s┘\n\n", sep)
}

func BoxT2(s string) {
	n := len(s) + 2
	sep := strings.Repeat("═", n)
	fmt.Printf("╔%s╗\n", sep)
	fmt.Printf("║ %s ║\n", colors.Black(s))
	fmt.Printf("╚%s╝\n\n", sep)
}

func TableHeader(s string) string {
	return colors.InvertedBlack(s)
}

func Choice(s string) {
	fmt.Printf("\n%s\n\n", colors.InvertedMagenta(s))
}

/*
func AccessingNetwork() {
	text := "Welcome to n2x network"
	prePad := strings.Repeat("─", 3)
	postPad := strings.Repeat("─", 70-len(text)-1)
	// pad := strings.Repeat("█", 60-len(text)-1)
	fmt.Printf("\n%s %s %s\n", colors.DarkCyan(prePad), colors.Black(text), colors.DarkCyan(postPad))
}
*/

/*
func WelcomeToTheNetwork() {
	text := "Welcome to n2x network"
	prePad := strings.Repeat("─", 3)
	postPad := strings.Repeat("─", 70-len(text)-1)
	// pad := strings.Repeat("█", 60-len(text)-1)
	fmt.Printf("\n%s %s %s\n", colors.DarkCyan(prePad), colors.Black(text), colors.DarkCyan(postPad))
}

func endOfTransmission() {
	text := "End-of-Transmission"
	prePad := strings.Repeat("─", 3)
	postPad := strings.Repeat("─", 70-len(text)-1)
	fmt.Printf("\n%s %s %s\n", colors.Black(prePad), colors.Black(text), colors.Black(postPad))
}
*/

func AuthenticationRequired() {
	fmt.Printf("%s\n\n", colors.DarkYellow("Authentication Required"))
}

func CmdLog(text string) {
	fmt.Printf("%s %s\n", colors.DarkGreen("🢂"), colors.Black(text))
}

func Authenticated() {
	fmt.Printf("%s\n\n", colors.InvertedDarkCyanAlt("Authenticated"))
}

func Logout() {
	fmt.Printf("%s\n\n", colors.InvertedDarkBlue("Logged out"))
}

func Datetime(tm int64) string {
	t := time.Unix(tm, 0)

	return datetime(t)
}

func DatetimeMilli(tm int64) string {
	t := time.UnixMilli(tm)

	return datetime(t)
}

func datetime(t time.Time) string {
	yr := t.Year()
	m := t.Month()
	d := t.Day()
	hr := t.Hour()
	min := t.Minute()
	sec := t.Second()

	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", yr, m, d, hr, min, sec)
}

func UserLocal(user string) string {
	return fmt.Sprintf("%s%s%s", colors.Green("<"), colors.DarkGreen(user), colors.Green(">"))
}

func UserRemote(user string) string {
	return fmt.Sprintf("%s%s%s", colors.Yellow("<"), colors.DarkYellow(user), colors.Yellow(">"))
	// return fmt.Sprintf("%s%s%s", colors.Red("<"), colors.DarkRed(user), colors.Red(">"))
}

func StrMRS() string {
	return colors.DarkBlue("[") + colors.Blue("mrs") + colors.DarkBlue("]")
}

func StrTier1() string {
	return colors.DarkBlue("[") + colors.Blue("tier-1") + colors.DarkBlue("]")
}

func StrOnline() string {
	// return colors.Green("online")
	return StrEnabled("online")
}

func StrOffline() string {
	// return colors.DarkRed("offline")
	return StrDisabled("offline")
}

func StrID(s string) string {
	return colors.DarkBlue("[") + colors.Cyan(s) + colors.DarkBlue("]")
}

func StrEnabled(s string) string {
	return colors.DarkGreen("[") + colors.Green(s) + colors.DarkGreen("]")
}

func StrDisabled(s string) string {
	// return colors.DarkRed("[") + colors.Red(s) + colors.DarkRed("]")
	return colors.Black("[") + colors.DarkRed(s) + colors.Black("]")
}

func StrOk(s string) string {
	return StrEnabled(s)
}

func StrError(s string) string {
	return StrDisabled(s)
}

func StrWarning(s string) string {
	return colors.DarkYellow("[") + colors.Yellow(s) + colors.DarkYellow("]")
}

func StrNormal(s string) string {
	return colors.DarkBlue("[") + colors.Blue(s) + colors.DarkBlue("]")
}

func StrFixedField(s string) string {
	return StrBlue(s)
}

func StrBlue(s string) string {
	return colors.DarkBlue("[") + colors.Blue(s) + colors.DarkBlue("]")
}

func StrMagenta(s string) string {
	return colors.DarkMagenta("[") + colors.Magenta(s) + colors.DarkMagenta("]")
}

func StrGreen(s string) string {
	return colors.Green("[") + colors.DarkGreen(s) + colors.Green("]")
}

func StrInactive(s string) string {
	return colors.Black("[") + s + colors.Black("]")
}

func Fit(s string, n int) string {
	max := n - 3

	length := len(s)

	if length > n {
		s = fmt.Sprintf("%s...", s[:max])
	} else {
		if (n - length) > 0 {
			spaces := strings.Repeat(" ", n-length)
			s = fmt.Sprintf("%s%s", s, spaces)
		}
	}

	return s
}

func AmountMoney(cents int64, currency string) string {
	var symbol string

	switch strings.ToLower(currency) {
	case "usd":
		symbol = "$"
	case "eur":
		symbol = "€"
	}

	var amount string
	if cents >= 0 {
		amount = colors.DarkGreen(fmt.Sprintf("%s%.2f", symbol, float64(cents)/100))
	} else {
		// amount = colors.DarkRed(fmt.Sprintf("%s%.2f", symbol, float64(cents)/100))
		amount = colors.DarkGreen(fmt.Sprintf("(%s%.2f)", symbol, float64(cents*-1)/100))
	}

	return fmt.Sprintf("%s", amount)
}

/*
func CustomerBalance(cents int64, currency string) string {
	var symbol string

	switch strings.ToLower(currency) {
	case "usd":
		symbol = "$"
	case "eur":
		symbol = "€"
	}

	var amount string
	if cents > 0 {
		amount = colors.DarkRed(fmt.Sprintf("%s%.2f", symbol, float64(cents)/100))
	} else {
		amount = colors.DarkGreen(fmt.Sprintf("%s%.2f", symbol, float64(cents)/100))
	}

	return fmt.Sprintf("%s", amount)
}
*/
