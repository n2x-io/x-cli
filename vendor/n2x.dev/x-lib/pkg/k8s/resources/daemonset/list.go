package daemonset

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) List(ns string) (*appsv1.DaemonSetList, error) {
	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	ctx := context.TODO()

	listOpts := metav1.ListOptions{}

	dsList, err := a.clientset.AppsV1().DaemonSets(ns).List(ctx, listOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.clientset.AppsV1().DaemonSets().List()", errors.Trace())
	}

	return dsList, nil
}
