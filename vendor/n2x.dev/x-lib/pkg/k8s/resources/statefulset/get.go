package statefulset

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
	"n2x.dev/x-lib/pkg/xlog"
)

func (a *API) Get(ns, name string) (*appsv1.StatefulSet, error) {
	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	ctx := context.TODO()

	getOpts := metav1.GetOptions{}

	s, err := a.clientset.AppsV1().StatefulSets(ns).Get(ctx, name, getOpts)
	if k8sErrors.IsNotFound(err) {
		xlog.Debugf("Secret %s/%s not found", ns, name)
		return nil, nil
	} else if statusError, isStatus := err.(*k8sErrors.StatusError); isStatus {
		xlog.Errorf("Unable to get secret %v", statusError.ErrStatus.Message)
		return nil, errors.Wrapf(err, "[%v] function a.clientset.AppsV1().StatefulSets().Get()", errors.Trace())
	} else if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.clientset.AppsV1().StatefulSets().Get()", errors.Trace())
	}

	return s, nil
}
