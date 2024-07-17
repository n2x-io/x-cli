package namespace

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) Delete(name string) error {
	for {
		ns, err := a.Get(name)
		if err != nil {
			return errors.Wrapf(err, "[%v] function a.Get()", errors.Trace())
		}
		if ns == nil { // not found
			return nil
		}

		if a.clientset == nil {
			clientset, err := config.NewClient(a.KubeConfig)
			if err != nil {
				return errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
			}
			a.clientset = clientset
		}

		ctx := context.TODO()

		deletePolicy := metav1.DeletePropagationForeground
		deleteOpts := metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}

		if err := a.clientset.CoreV1().Namespaces().Delete(ctx, name, deleteOpts); err != nil {
			return errors.Wrapf(err, "[%v] function a.clientset.CoreV1().Namespaces().Delete()", errors.Trace())
		}

		time.Sleep(time.Second)
	}
}
