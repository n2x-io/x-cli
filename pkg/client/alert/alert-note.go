package alert

import (
	"context"
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/events"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) NewNote(a *events.Alert) {
	if a == nil {
		a = getAlert()
	}

	nxc, grpcConn := grpc.GetMonitoringAPIClient()
	defer grpcConn.Close()

	acr := &events.AlertNewCommentRequest{
		AccountID: a.AccountID,
		AlertID:   a.AlertID,
		Comment: &events.AlertComment{
			Timestamp: time.Now().UnixMilli(),
			// UserID: userID,
			// UserEmail: userEmail,
			Text: input.GetMultiline("Note:", "", "", survey.Required),
		},
	}

	s := output.Spinner()

	_, err := nxc.NewAlertComment(context.TODO(), acr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to add alert note")
	}

	s.Stop()

	fmt.Printf(`

Your input has been sent!

Thanks for using n2x!

`)
}
