package alert

import "n2x.dev/x-api-go/grpc/resources/events"

func (api *API) Show() *events.Alert {
	a := getAlert()

	Output().Show(a)

	return a
}
