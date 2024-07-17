package pod

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-lib/pkg/k8s/config"
	"n2x.dev/x-lib/pkg/n2x"
)

func (a *API) NewContainer(i interface{}, appLabel config.AppLabel) *corev1.Container {
	var ni *topology.NodeInstance

	switch appLabel {
	case config.AppLabelNode:
		ni = i.(*topology.NodeInstance)
	default:
		return nil
	}

	configVolumeName := fmt.Sprintf("%s-config", ni.Node.KubernetesAttrs.Name)

	return &corev1.Container{
		Name:            ni.Node.KubernetesAttrs.Name,
		Image:           ni.Node.KubernetesAttrs.Image,
		ImagePullPolicy: corev1.PullAlways,
		SecurityContext: &corev1.SecurityContext{
			Privileged: n2x.Bool(true), // only needed by sysctl to enable ipv6
			Capabilities: &corev1.Capabilities{
				Add: []corev1.Capability{
					"net_admin",
				},
			},
			RunAsUser:    n2x.Int64(0),
			RunAsGroup:   n2x.Int64(0),
			RunAsNonRoot: n2x.Bool(false),
		},
		Args: []string{
			"start",
		},
		Ports: []corev1.ContainerPort{
			{
				Name:          "n2x-p2p",
				Protocol:      corev1.ProtocolTCP,
				ContainerPort: ni.Node.Agent.Port,
			},
		},
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      "dev-net-tun",
				MountPath: "/dev/net/tun",
				ReadOnly:  true,
			},
			{
				Name:      configVolumeName,
				MountPath: "/etc/n2x",
				ReadOnly:  true,
			},
		},
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				// cpu: 30m
				corev1.ResourceCPU: *resource.NewMilliQuantity(30, resource.DecimalSI),
				// memory: 200Mi
				corev1.ResourceMemory: *resource.NewQuantity(200*1024*1024, resource.BinarySI),
			},
			Requests: corev1.ResourceList{
				// cpu: 15m
				corev1.ResourceCPU: *resource.NewMilliQuantity(15, resource.DecimalSI),
				// memory: 50Mi
				corev1.ResourceMemory: *resource.NewQuantity(50*1024*1024, resource.BinarySI),
			},
		},
	}
}
