package namespace

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) Create(name string) (*corev1.Namespace, error) {
	ns, err := a.Get(name)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.Get()", errors.Trace())
	}
	if ns != nil { // already found
		return ns, nil
	}

	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	ctx := context.TODO()

	ns = a.New(name)

	createOpts := metav1.CreateOptions{}

	ns, err = a.clientset.CoreV1().Namespaces().Create(ctx, ns, createOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.clientset.CoreV1().Namespaces().Create()", errors.Trace())
	}

	return ns, nil
}
