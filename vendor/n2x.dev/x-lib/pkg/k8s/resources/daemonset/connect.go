package daemonset

import (
	"context"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) AddContainer(ns, name string, c corev1.Container, vols []corev1.Volume, labels map[string]string) error {
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
		ds, err := a.clientset.AppsV1().DaemonSets(ns).Get(ctx, name, getOpts)
		if err != nil {
			return errors.Wrapf(err, "[%v] function a.clientset.AppsV1().DaemonSets().Get()", errors.Trace())
		}

		// add n2x-node container
		cUpdated := false
		for _, container := range ds.Spec.Template.Spec.Containers {
			if strings.Contains(container.Image, "n2x-node") {
				cUpdated = true
				break
			}
		}
		if !cUpdated {
			ds.Spec.Template.Spec.Containers = append(ds.Spec.Template.Spec.Containers, c)
		}

		// add n2x-node volumes
		vol0Updated := false
		vol1Updated := false
		for _, volume := range ds.Spec.Template.Spec.Volumes {
			if volume.Name == vols[0].Name {
				vol0Updated = true
			}
			if volume.Name == vols[1].Name {
				vol1Updated = true
			}
		}
		if !vol0Updated {
			ds.Spec.Template.Spec.Volumes = append(ds.Spec.Template.Spec.Volumes, vols[0])
		}
		if !vol1Updated {
			ds.Spec.Template.Spec.Volumes = append(ds.Spec.Template.Spec.Volumes, vols[1])
		}

		// add n2x-node labels
		if ds.ObjectMeta.Labels == nil {
			ds.ObjectMeta.Labels = make(map[string]string)
		}
		for k, v := range labels {
			ds.ObjectMeta.Labels[k] = v
		}

		_, err = a.clientset.AppsV1().DaemonSets(ns).Update(ctx, ds, updateOpts)
		return err
	})
	if retryErr != nil {
		return errors.Wrapf(retryErr, "[%v] function retry.RetryOnConflict()", errors.Trace())
	}

	return nil
}
