package clusterrolebinding

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) Update(ns, name string) error {
	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	ctx := context.TODO()

	updateOpts := metav1.UpdateOptions{}

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Service before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver

		crb, err := a.Get(name)
		if err != nil {
			return errors.Wrapf(err, "[%v] function a.Get()", errors.Trace())
		}
		if crb == nil { // not found
			return nil
		}

		if crb.Subjects == nil {
			crb.Subjects = make([]rbacv1.Subject, 0)
		}

		for _, subj := range crb.Subjects {
			if subj.Kind == "ServiceAccount" && subj.Name == name && subj.Namespace == ns {
				return nil
			}
		}

		crb.Subjects = append(crb.Subjects, rbacv1.Subject{
			Kind:      "ServiceAccount",
			Name:      name,
			Namespace: ns,
		})

		_, err = a.clientset.RbacV1().ClusterRoleBindings().Update(ctx, crb, updateOpts)
		return err
	})
	if retryErr != nil {
		return errors.Wrapf(retryErr, "[%v] function retry.RetryOnConflict()", errors.Trace())
	}

	return nil
}
