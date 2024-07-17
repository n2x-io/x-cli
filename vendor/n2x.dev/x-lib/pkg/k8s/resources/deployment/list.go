package deployment

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) List(ns string) (*appsv1.DeploymentList, error) {
	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	ctx := context.TODO()

	listOpts := metav1.ListOptions{}

	dList, err := a.clientset.AppsV1().Deployments(ns).List(ctx, listOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.clientset.AppsV1().Deployments().List()", errors.Trace())
	}

	return dList, nil
}
