package secret

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"n2x.dev/x-lib/pkg/k8s/config"
)

type Interface interface {
	List(ns string) (*corev1.SecretList, error)
	New(i interface{}, appLabel config.AppLabel) *corev1.Secret
	Get(ns, name string) (*corev1.Secret, error)
	Create(s *corev1.Secret) (*corev1.Secret, error)
	Delete(ns, name string) error
}

type API struct {
	KubeConfig []byte
	clientset  *kubernetes.Clientset
}
