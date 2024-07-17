package service

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) List(ns string) (*corev1.ServiceList, error) {
	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	listOpts := metav1.ListOptions{
		// FieldSelector: fmt.Sprintf("spec.loadBalancerIP=%s", externalIPv4),
	}

	svcList, err := a.clientset.CoreV1().Services(ns).List(context.TODO(), listOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.clientset.CoreV1().Services().List()", errors.Trace())
	}

	return svcList, nil
}
