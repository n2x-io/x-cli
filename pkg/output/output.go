package output

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/hokaccha/go-prettyjson"
	"n2x.dev/x-api-go/grpc/common/empty"
	"n2x.dev/x-api-go/grpc/common/status"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-lib/pkg/utils/colors"
	"n2x.dev/x-lib/pkg/utils/msg"
)

func Show(r interface{}) {
	switch sr := r.(type) {
	case *empty.Response:
		msg.Info("Done")

		return
	case *status.StatusResponse:
		switch sr.Code {
		case status.StatusCode_OK:
			msg.Info("Done")
		case status.StatusCode_CANCELED:
			msg.Warnf("Operation CANCELED: %s", sr.Message)
		case status.StatusCode_INTERRUPTED:
			msg.Warnf("Operation INTERRUPTED: %s", sr.Message)
		case status.StatusCode_INCOMPLETE:
			msg.Warnf("Operation INCOMPLETE: %s", sr.Message)
		case status.StatusCode_FAILED:
			msg.Errorf("Operation FAILED: %s", sr.Message)
		default:
			msg.Warn("Unknown code")
		}

		return
	}

	fmt.Println()

	s, err := prettyjson.Marshal(r)
	if err != nil {
		msg.Error("Invalid data")
		os.Exit(1)
	}
	fmt.Println(string(s))
}

/*
func ShowWaiting() {
	fmt.Printf("\nOk, proceeding, wait a minute..\n\n")
}
*/

func Spinner() *spinner.Spinner {
	// fmt.Println()

	s := spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	s.Prefix = "  "
	s.Suffix = colors.Black(" processing")
	// s.Color("blue", "bold")
	s.Start()

	// s.Stop() does not erase the line due to a bug in v1.16.0
	// workaround before after s.Stop():
	// fmt.Fprintf(s.Writer, "\r") // erases to end of line

	return s
}

func PercentBar(n int) string {
	var p int

	if n >= 0 && n <= 100 {
		p = n / 10
		return fmt.Sprintf("[%s%%] %s", colors.White(fmt.Sprintf("%d", n)), strings.Repeat("█", p)+strings.Repeat("░", 10-p))
	}

	return "n/a"
}

func ConfirmDeletion() {
	fmt.Println()

	if !input.GetConfirm("Confirm deletion?", false) {
		fmt.Println()
		os.Exit(0)
	}
}
