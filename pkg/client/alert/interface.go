package alert

import (
	"n2x.dev/x-api-go/grpc/resources/events"
	"n2x.dev/x-cli/pkg/client/alert/output"
)

type Interface interface {
	List()
	Show() *events.Alert
	Delete()
	NewNote(a *events.Alert)
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
