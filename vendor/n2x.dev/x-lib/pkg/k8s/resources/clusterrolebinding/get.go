package clusterrolebinding

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
	"n2x.dev/x-lib/pkg/xlog"
)

func (a *API) Get(name string) (*rbacv1.ClusterRoleBinding, error) {
	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	ctx := context.TODO()

	getOpts := metav1.GetOptions{}

	crb, err := a.clientset.RbacV1().ClusterRoleBindings().Get(ctx, name, getOpts)
	if k8sErrors.IsNotFound(err) {
		xlog.Debugf("ClusterRoleBinding %s not found", name)
		return nil, nil
	} else if statusError, isStatus := err.(*k8sErrors.StatusError); isStatus {
		xlog.Errorf("Unable to get ClusterRoleBinding %v", statusError.ErrStatus.Message)
		return nil, errors.Wrapf(err, "[%v] function a.clientset.RbacV1().ClusterRoleBindings().Get()", errors.Trace())
	} else if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.clientset.RbacV1().ClusterRoleBindings().Get()", errors.Trace())
	}

	return crb, nil
}
