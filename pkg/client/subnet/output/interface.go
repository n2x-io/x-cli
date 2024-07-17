package output

import "n2x.dev/x-api-go/grpc/resources/topology"

type Interface interface {
	List(vrfs map[string]*topology.Subnet)
	Show(v *topology.Subnet)
}
type API struct{}
