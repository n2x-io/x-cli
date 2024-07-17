package output

import "n2x.dev/x-cli/pkg/client/k8s/resource"

type Interface interface {
	List(k8sResources map[string]*resource.KubernetesResource)
}
type API struct{}
