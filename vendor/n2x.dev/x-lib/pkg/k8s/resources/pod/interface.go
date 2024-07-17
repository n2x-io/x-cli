package pod

import (
	corev1 "k8s.io/api/core/v1"
	"n2x.dev/x-lib/pkg/k8s/config"
)

type Interface interface {
	NewLabels(i interface{}, appLabel config.AppLabel) map[string]string
	NewContainer(i interface{}, appLabel config.AppLabel) *corev1.Container
	NewVolumes(i interface{}, appLabel config.AppLabel) []corev1.Volume
}

type API struct{}
