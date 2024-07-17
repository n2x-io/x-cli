package statefulset

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) RemoveContainer(ns, name string, c corev1.Container, vols []corev1.Volume, labels map[string]string) error {
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
		ss, err := a.clientset.AppsV1().StatefulSets(ns).Get(ctx, name, getOpts)
		if err != nil {
			return errors.Wrapf(err, "[%v] function a.clientset.AppsV1().StatefulSets().Get()", errors.Trace())
		}

		// remove n2x-node container
		containers := make([]corev1.Container, 0)
		for _, container := range ss.Spec.Template.Spec.Containers {
			if container.Name == c.Name {
				continue
			}
			containers = append(containers, container)
		}
		ss.Spec.Template.Spec.Containers = containers

		// remove n2x-node volumes
		volumes := make([]corev1.Volume, 0)
		for _, v := range ss.Spec.Template.Spec.Volumes {
			if v.Name == vols[0].Name || v.Name == vols[1].Name {
				continue
			}
			volumes = append(volumes, v)
		}
		ss.Spec.Template.Spec.Volumes = volumes

		// remove n2x-node labels
		if ss.ObjectMeta.Labels != nil {
			for k := range labels {
				delete(ss.ObjectMeta.Labels, k)
			}
		}

		_, err = a.clientset.AppsV1().StatefulSets(ns).Update(ctx, ss, updateOpts)
		return err
	})
	if retryErr != nil {
		return errors.Wrapf(retryErr, "[%v] function retry.RetryOnConflict()", errors.Trace())
	}

	return nil
}
