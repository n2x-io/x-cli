package pod

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) NewVolumes(i interface{}, appLabel config.AppLabel) []corev1.Volume {
	var ni *topology.NodeInstance

	switch appLabel {
	case config.AppLabelNode:
		ni = i.(*topology.NodeInstance)
	default:
		return nil
	}

	configVolumeName := fmt.Sprintf("%s-config", ni.Node.KubernetesAttrs.Name)

	return []corev1.Volume{
		{
			Name: "dev-net-tun",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/dev/net/tun",
				},
			},
		},
		{
			Name: configVolumeName,
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: ni.Node.KubernetesAttrs.Name,
				},
			},
		},
	}
}
