package clusterrolebinding

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/client-go/kubernetes"
)

type Interface interface {
	List() (*rbacv1.ClusterRoleBindingList, error)
	New(ns, name string) *rbacv1.ClusterRoleBinding
	Get(name string) (*rbacv1.ClusterRoleBinding, error)
	Create(ns, name string) (*rbacv1.ClusterRoleBinding, error)
	Update(ns, name string) error
	Delete(name string) error
}

type API struct {
	KubeConfig []byte
	clientset  *kubernetes.Clientset
}
