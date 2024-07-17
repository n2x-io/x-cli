package pod

import (
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) NewLabels(i interface{}, appLabel config.AppLabel) map[string]string {
	var ni *topology.NodeInstance

	switch appLabel {
	case config.AppLabelNode:
		ni = i.(*topology.NodeInstance)
	default:
		return nil
	}

	return config.NodeLabels(ni)
}
