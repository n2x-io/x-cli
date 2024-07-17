package node

import (
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
	"n2x.dev/x-lib/pkg/k8s/resources"
)

func (a *API) ConnectStatefulSet(ns, name string, ni *topology.NodeInstance) error {
	secret := resources.API(a.KubeConfig).Secret().New(ni, config.AppLabelNode)
	if _, err := resources.API(a.KubeConfig).Secret().Create(secret); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().Secret().Create()", errors.Trace())
	}

	c := resources.API(a.KubeConfig).Pod().NewContainer(ni, config.AppLabelNode)
	vols := resources.API(a.KubeConfig).Pod().NewVolumes(ni, config.AppLabelNode)
	labels := resources.API(a.KubeConfig).Pod().NewLabels(ni, config.AppLabelNode)
	if err := resources.API(a.KubeConfig).StatefulSet().AddContainer(ns, name, *c, vols, labels); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().StatefulSet().AddContainer()", errors.Trace())
	}

	return nil
}

func (a *API) DisconnectStatefulSet(ns, name string, ni *topology.NodeInstance) error {
	c := resources.API(a.KubeConfig).Pod().NewContainer(ni, config.AppLabelNode)
	vols := resources.API(a.KubeConfig).Pod().NewVolumes(ni, config.AppLabelNode)
	labels := resources.API(a.KubeConfig).Pod().NewLabels(ni, config.AppLabelNode)
	if err := resources.API(a.KubeConfig).StatefulSet().RemoveContainer(ns, name, *c, vols, labels); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().StatefulSet().RemoveContainer()", errors.Trace())
	}

	secretName := ni.Node.KubernetesAttrs.Name

	if err := resources.API(a.KubeConfig).Secret().Delete(ns, secretName); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().Secret().Delete()", errors.Trace())
	}

	return nil
}
