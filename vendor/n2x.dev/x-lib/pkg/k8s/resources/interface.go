package resources

import (
	"n2x.dev/x-lib/pkg/k8s/resources/clusterrolebinding"
	"n2x.dev/x-lib/pkg/k8s/resources/daemonset"
	"n2x.dev/x-lib/pkg/k8s/resources/deployment"
	"n2x.dev/x-lib/pkg/k8s/resources/namespace"
	"n2x.dev/x-lib/pkg/k8s/resources/pod"
	"n2x.dev/x-lib/pkg/k8s/resources/secret"
	"n2x.dev/x-lib/pkg/k8s/resources/service"
	"n2x.dev/x-lib/pkg/k8s/resources/serviceaccount"
	"n2x.dev/x-lib/pkg/k8s/resources/statefulset"
)

type Interface interface {
	ServiceAccount() serviceaccount.Interface
	ClusterRoleBinding() clusterrolebinding.Interface
	Namespace() namespace.Interface
	Secret() secret.Interface
	StatefulSet() statefulset.Interface
	Deployment() deployment.Interface
	DaemonSet() daemonset.Interface
	Service() service.Interface
	Pod() pod.Interface
}

type api struct {
	kubeConfig []byte
}

func API(kubeConfig []byte) Interface {
	return &api{kubeConfig: kubeConfig}
}

func (a *api) ServiceAccount() serviceaccount.Interface {
	return &serviceaccount.API{KubeConfig: a.kubeConfig}
}

func (a *api) ClusterRoleBinding() clusterrolebinding.Interface {
	return &clusterrolebinding.API{KubeConfig: a.kubeConfig}
}

func (a *api) Namespace() namespace.Interface {
	return &namespace.API{KubeConfig: a.kubeConfig}
}

func (a *api) Secret() secret.Interface {
	return &secret.API{KubeConfig: a.kubeConfig}
}

func (a *api) StatefulSet() statefulset.Interface {
	return &statefulset.API{KubeConfig: a.kubeConfig}
}

func (a *api) Deployment() deployment.Interface {
	return &deployment.API{KubeConfig: a.kubeConfig}
}

func (a *api) DaemonSet() daemonset.Interface {
	return &daemonset.API{KubeConfig: a.kubeConfig}
}

func (a *api) Service() service.Interface {
	return &service.API{KubeConfig: a.kubeConfig}
}

func (a *api) Pod() pod.Interface {
	return &pod.API{}
}
