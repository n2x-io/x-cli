package output

import "n2x.dev/x-api-go/grpc/resources/ops"

type Interface interface {
	List(projects map[string]*ops.Project)
	Show(p *ops.Project)
}
type API struct{}
