package node

import (
	"n2x.dev/x-api-go/grpc/resources/topology"
)

type Interface interface {
	CreateGateway(ni *topology.NodeInstance) error
	DeleteGateway(ns, name string) error

	ConnectStatefulSet(ns, name string, ni *topology.NodeInstance) error
	DisconnectStatefulSet(ns, name string, ni *topology.NodeInstance) error
	ConnectDeployment(ns, name string, ni *topology.NodeInstance) error
	DisconnectDeployment(ns, name string, ni *topology.NodeInstance) error
	ConnectDaemonSet(ns, name string, ni *topology.NodeInstance) error
	DisconnectDaemonSet(ns, name string, ni *topology.NodeInstance) error
}

type API struct {
	KubeConfig []byte
}
