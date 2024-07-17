package objects

import (
	"n2x.dev/x-lib/pkg/k8s/objects/controller"
	"n2x.dev/x-lib/pkg/k8s/objects/node"
)

type Interface interface {
	Controller() controller.Interface
	Node() node.Interface
}

type API struct {
	KubeConfig []byte
}

func (a *API) Controller() controller.Interface {
	return &controller.API{KubeConfig: a.KubeConfig}
}

func (a *API) Node() node.Interface {
	return &node.API{KubeConfig: a.KubeConfig}
}
