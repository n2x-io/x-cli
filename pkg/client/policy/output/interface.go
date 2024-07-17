package output

import "n2x.dev/x-api-go/grpc/resources/topology"

type Interface interface {
	Show(s *topology.Subnet)
}
type API struct{}
