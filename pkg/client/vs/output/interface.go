package output

import "n2x.dev/x-api-go/grpc/resources/topology"

type Interface interface {
	List(vss map[string]*topology.VS)
	Show(vs *topology.VS)
}
type API struct{}
