package output

import "n2x.dev/x-api-go/grpc/resources/events"

type Interface interface {
	List(alerts map[string]*events.Alert)
	Show(a *events.Alert)

	AlertStatus(s events.Status) string
}
type API struct{}
