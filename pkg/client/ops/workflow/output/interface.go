package output

import "n2x.dev/x-api-go/grpc/resources/ops"

type Interface interface {
	List(workflows map[string]*ops.Workflow)
	Show(wf *ops.Workflow)
}
type API struct{}
