package deployment

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type Interface interface {
	List(ns string) (*appsv1.DeploymentList, error)
	AddContainer(ns, name string, c corev1.Container, vols []corev1.Volume, labels map[string]string) error
	RemoveContainer(ns, name string, c corev1.Container, vols []corev1.Volume, labels map[string]string) error
}

type API struct {
	KubeConfig []byte
	clientset  *kubernetes.Clientset
}
