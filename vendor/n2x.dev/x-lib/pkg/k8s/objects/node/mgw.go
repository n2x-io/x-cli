package node

import (
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
	"n2x.dev/x-lib/pkg/k8s/resources"
)

func (a *API) CreateGateway(ni *topology.NodeInstance) error {
	ns := ni.Node.KubernetesAttrs.Namespace

	if _, err := resources.API(a.KubeConfig).Namespace().Create(ns); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().Namespace().Create()", errors.Trace())
	}

	if _, err := resources.API(a.KubeConfig).ServiceAccount().Create(ns, config.ServiceAccountView); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().ServiceAccount().Create()", errors.Trace())
	}

	if _, err := resources.API(a.KubeConfig).ClusterRoleBinding().Create(ns, config.ClusterRoleBindingView); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().ClusterRoleBinding().Create()", errors.Trace())
	}

	svc := resources.API(a.KubeConfig).Service().New(ni, config.AppLabelNode)
	if _, err := resources.API(a.KubeConfig).Service().Create(svc); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().Service().Create()", errors.Trace())
	}

	secret := resources.API(a.KubeConfig).Secret().New(ni, config.AppLabelNode)
	if _, err := resources.API(a.KubeConfig).Secret().Create(secret); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().Secret().Create()", errors.Trace())
	}

	stfs := resources.API(a.KubeConfig).StatefulSet().New(ni, config.AppLabelNode)
	if _, err := resources.API(a.KubeConfig).StatefulSet().Create(stfs); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().StatefulSet().Create()", errors.Trace())
	}

	return nil
}

func (a *API) DeleteGateway(ns, name string) error {
	if err := resources.API(a.KubeConfig).StatefulSet().Delete(ns, name); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().StatefulSet().Delete()", errors.Trace())
	}

	if err := resources.API(a.KubeConfig).Secret().Delete(ns, name); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().Secret().Delete()", errors.Trace())
	}

	if err := resources.API(a.KubeConfig).Service().Delete(ns, name); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().Service().Delete()", errors.Trace())
	}

	return nil
}
