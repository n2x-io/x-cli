package output

import "n2x.dev/x-api-go/grpc/resources/events"

type Interface interface {
	ShowMetrics(em *events.EventMetrics)
	FailureProbability(em *events.EventMetrics) string
}
type API struct{}
