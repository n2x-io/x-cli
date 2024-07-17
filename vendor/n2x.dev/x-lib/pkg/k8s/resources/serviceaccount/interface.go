package serviceaccount

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type Interface interface {
	List(ns string) (*corev1.ServiceAccountList, error)
	New(ns, name string) *corev1.ServiceAccount
	Get(ns, name string) (*corev1.ServiceAccount, error)
	Create(ns, name string) (*corev1.ServiceAccount, error)
	Delete(ns, name string) error
}

type API struct {
	KubeConfig []byte
	clientset  *kubernetes.Clientset
}
