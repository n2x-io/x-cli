package clusterrolebinding

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) List() (*rbacv1.ClusterRoleBindingList, error) {
	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	ctx := context.TODO()

	listOpts := metav1.ListOptions{}

	crbList, err := a.clientset.RbacV1().ClusterRoleBindings().List(ctx, listOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.clientset.RbacV1().ClusterRoleBindings().List()", errors.Trace())
	}

	return crbList, nil
}
