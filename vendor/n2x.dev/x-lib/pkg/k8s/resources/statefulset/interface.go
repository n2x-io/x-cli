package statefulset

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"n2x.dev/x-lib/pkg/k8s/config"
)

type Interface interface {
	List(ns string) (*appsv1.StatefulSetList, error)
	New(i interface{}, appLabel config.AppLabel) *appsv1.StatefulSet
	Get(ns, name string) (*appsv1.StatefulSet, error)
	Create(s *appsv1.StatefulSet) (*appsv1.StatefulSet, error)
	Delete(ns, name string) error
	AddContainer(ns, name string, c corev1.Container, vols []corev1.Volume, labels map[string]string) error
	RemoveContainer(ns, name string, c corev1.Container, vols []corev1.Volume, labels map[string]string) error
}

type API struct {
	KubeConfig []byte
	clientset  *kubernetes.Clientset
}
