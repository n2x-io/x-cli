package service

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) RemoveAnnotations(ns string, svcName string, annotations map[string]string) error {
	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	ctx := context.TODO()

	getOpts := metav1.GetOptions{}
	updateOpts := metav1.UpdateOptions{}

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Service before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		s, err := a.clientset.CoreV1().Services(ns).Get(ctx, svcName, getOpts)
		if err != nil {
			return errors.Wrapf(err, "[%v] function a.clientset.CoreV1().Services().Get()", errors.Trace())
		}

		if s.ObjectMeta.Annotations == nil {
			return nil
		}

		for k := range annotations {
			delete(s.ObjectMeta.Annotations, k)
		}

		_, err = a.clientset.CoreV1().Services(ns).Update(ctx, s, updateOpts)
		return err
	})
	if retryErr != nil {
		return errors.Wrapf(retryErr, "[%v] function retry.RetryOnConflict()", errors.Trace())
	}

	return nil
}
