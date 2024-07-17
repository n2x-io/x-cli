package k8s

import (
	"n2x.dev/x-lib/pkg/k8s/objects"
	"n2x.dev/x-lib/pkg/k8s/resources"
)

type Interface interface {
	Resources() resources.Interface
	Objects() objects.Interface
}

type api struct {
	kubeConfig []byte
}

func API(kubeConfig []byte) Interface {
	return &api{kubeConfig: kubeConfig}
}

func (a *api) Resources() resources.Interface {
	return resources.API(a.kubeConfig)
}

func (a *api) Objects() objects.Interface {
	return &objects.API{KubeConfig: a.kubeConfig}
}
