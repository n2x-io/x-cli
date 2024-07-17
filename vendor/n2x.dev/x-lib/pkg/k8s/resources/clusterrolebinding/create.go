package clusterrolebinding

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) Create(ns, name string) (*rbacv1.ClusterRoleBinding, error) {
	if err := a.Update(ns, name); err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.Update()", errors.Trace())
	}

	crb, err := a.Get(name)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.Get()", errors.Trace())
	}
	if crb != nil { // already found
		return crb, nil
	}

	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	ctx := context.TODO()

	crb = a.New(ns, name)

	createOpts := metav1.CreateOptions{}

	crb, err = a.clientset.RbacV1().ClusterRoleBindings().Create(ctx, crb, createOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.clientset.RbacV1().ClusterRoleBindings().Create()", errors.Trace())
	}

	return crb, nil
}
