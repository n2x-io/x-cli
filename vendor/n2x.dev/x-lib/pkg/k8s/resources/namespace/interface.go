package namespace

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type Interface interface {
	List() (*corev1.NamespaceList, error)
	New(name string) *corev1.Namespace
	Get(name string) (*corev1.Namespace, error)
	Create(name string) (*corev1.Namespace, error)
	Delete(name string) error
}

type API struct {
	KubeConfig []byte
	clientset  *kubernetes.Clientset
}
