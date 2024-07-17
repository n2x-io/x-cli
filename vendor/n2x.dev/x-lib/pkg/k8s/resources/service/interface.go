package service

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"n2x.dev/x-lib/pkg/k8s/config"
)

type Interface interface {
	List(ns string) (*corev1.ServiceList, error)
	New(i interface{}, appLabel config.AppLabel) *corev1.Service
	Get(ns, name string) (*corev1.Service, error)
	Create(s *corev1.Service) (*corev1.Service, error)
	Delete(ns, name string) error

	InsertAnnotations(ns string, svcName string, annotations map[string]string) error
	RemoveAnnotations(ns string, svcName string, annotations map[string]string) error

	GetNodePort() (int, error)
}

type API struct {
	KubeConfig []byte
	clientset  *kubernetes.Clientset
}
