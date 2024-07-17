package controller

import (
	"n2x.dev/x-api-go/grpc/resources/topology"
)

type Interface interface {
	Create(ni *topology.NodeInstance) error
}

type API struct {
	KubeConfig []byte
}

func (a *API) Create(ni *topology.NodeInstance) error {
	// if _, err := resources.API(a.KubeConfig).Namespace().Create(ni.K8SOpts.Ns); err != nil {
	// 	return errors.Wrapf(err, "[%v] function resources.API().Namespace().Create()", errors.Trace())
	// }

	// svc := resources.API(a.KubeConfig).Service().New(ni, config.AppLabelNode)
	// if _, err := resources.API(a.KubeConfig).Service().Create(svc); err != nil {
	// 	return errors.Wrapf(err, "[%v] function resources.API().Service().Create()", errors.Trace())
	// }

	// secret := resources.API(a.KubeConfig).Secret().New(ni, config.AppLabelNode)
	// if _, err := resources.API(a.KubeConfig).Secret().Create(secret); err != nil {
	// 	return errors.Wrapf(err, "[%v] function resources.API().Secret().Create()", errors.Trace())
	// }

	// stfs := resources.API(a.KubeConfig).Statefulset().New(ni, config.AppLabelNode)
	// if _, err := resources.API(a.KubeConfig).Statefulset().Create(stfs); err != nil {
	// 	return errors.Wrapf(err, "[%v] function resources.API().Statefulset().Create()", errors.Trace())
	// }

	return nil
}
