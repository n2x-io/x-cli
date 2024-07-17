package secret

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func (a *API) New(i interface{}, appLabel config.AppLabel) *corev1.Secret {
	var ni *topology.NodeInstance

	switch appLabel {
	case config.AppLabelNode:
		ni = i.(*topology.NodeInstance)
	default:
		return nil
	}

	ymlFile := fmt.Sprintf("%s.yml", appLabel.String())

	return &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ni.Node.KubernetesAttrs.Name,
			Namespace: ni.Node.KubernetesAttrs.Namespace,
			Labels:    config.NodeLabels(ni),
		},
		Type: corev1.SecretTypeOpaque,
		StringData: map[string]string{
			ymlFile: ni.Config.YAML,
		},
	}
}
