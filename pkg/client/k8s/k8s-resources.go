package k8s

import (
	"fmt"
	"os"
	"sort"

	"n2x.dev/x-cli/pkg/client/k8s/resource"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/k8s"
	"n2x.dev/x-lib/pkg/utils/msg"
)

func (api *API) getKubernetesPods() map[string]*resource.KubernetesResource {
	rl := make(map[string]*resource.KubernetesResource)

	for k, v := range api.getKubernetesStatefulSets() {
		rl[k] = v
	}

	for k, v := range api.getKubernetesDeployments() {
		rl[k] = v
	}

	for k, v := range api.getKubernetesDaemonSets() {
		rl[k] = v
	}

	return rl
}

func (api *API) getKubernetesGateways() map[string]*resource.KubernetesResource {
	rl := make(map[string]*resource.KubernetesResource)

	for k, v := range api.getKubernetesServices() {
		rl[k] = v
	}

	return rl
}

func (api *API) getKubernetesServices() map[string]*resource.KubernetesResource {
	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	svcl, err := k8s.API(api.kubeConfig).Resources().Service().List("")
	if err != nil {
		status.Error(err, "Unable to get kubernetes services")
	}

	rl := make(map[string]*resource.KubernetesResource)

	for _, s := range svcl.Items {
		r := &resource.KubernetesResource{
			KubernetesResourceType: resource.KubernetesResourceTypeService,
			Namespace:              s.ObjectMeta.Namespace,
			Name:                   s.ObjectMeta.Name,
			Connected:              false,
		}
		r.ParseLabelsGateway(s.ObjectMeta.Labels)

		rID := fmt.Sprintf("%s:%s", s.ObjectMeta.Namespace, s.ObjectMeta.Name)

		rl[rID] = r
	}

	return rl
}

func (api *API) getKubernetesServicesAnnotations() map[string]*resource.KubernetesResource {
	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	svcl, err := k8s.API(api.kubeConfig).Resources().Service().List("")
	if err != nil {
		status.Error(err, "Unable to get kubernetes services")
	}

	rl := make(map[string]*resource.KubernetesResource)

	for _, s := range svcl.Items {
		r := &resource.KubernetesResource{
			KubernetesResourceType: resource.KubernetesResourceTypeService,
			Namespace:              s.ObjectMeta.Namespace,
			Name:                   s.ObjectMeta.Name,
			Connected:              false,
			ServiceAnnotations:     make(map[string]string),
		}
		r.ParseServiceAnnotations(s.ObjectMeta.Annotations)

		rID := fmt.Sprintf("%s:%s", s.ObjectMeta.Namespace, s.ObjectMeta.Name)

		rl[rID] = r
	}

	return rl
}

func (api *API) getKubernetesStatefulSets() map[string]*resource.KubernetesResource {
	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	stfsl, err := k8s.API(api.kubeConfig).Resources().StatefulSet().List("")
	if err != nil {
		status.Error(err, "Unable to get kubernetes statefulSets")
	}

	rl := make(map[string]*resource.KubernetesResource)

	for _, ss := range stfsl.Items {
		r := &resource.KubernetesResource{
			KubernetesResourceType: resource.KubernetesResourceTypeStatefulSet,
			Namespace:              ss.ObjectMeta.Namespace,
			Name:                   ss.ObjectMeta.Name,
			Connected:              false,
		}
		r.ParseLabelsPod(ss.ObjectMeta.Labels)

		rID := fmt.Sprintf("%s:%s", ss.ObjectMeta.Namespace, ss.ObjectMeta.Name)

		rl[rID] = r
	}

	return rl
}

func (api *API) getKubernetesDeployments() map[string]*resource.KubernetesResource {
	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	dl, err := k8s.API(api.kubeConfig).Resources().Deployment().List("")
	if err != nil {
		status.Error(err, "Unable to get kubernetes deployments")
	}

	rl := make(map[string]*resource.KubernetesResource)

	for _, d := range dl.Items {
		r := &resource.KubernetesResource{
			KubernetesResourceType: resource.KubernetesResourceTypeDeployment,
			Namespace:              d.ObjectMeta.Namespace,
			Name:                   d.ObjectMeta.Name,
			Connected:              false,
		}
		r.ParseLabelsPod(d.ObjectMeta.Labels)

		rID := fmt.Sprintf("%s:%s", d.ObjectMeta.Namespace, d.ObjectMeta.Name)

		rl[rID] = r
	}

	return rl
}

func (api *API) getKubernetesDaemonSets() map[string]*resource.KubernetesResource {
	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	dsl, err := k8s.API(api.kubeConfig).Resources().DaemonSet().List("")
	if err != nil {
		status.Error(err, "Unable to get kubernetes daemonSets")
	}

	rl := make(map[string]*resource.KubernetesResource)

	for _, ds := range dsl.Items {
		r := &resource.KubernetesResource{
			KubernetesResourceType: resource.KubernetesResourceTypeDaemonSet,
			Namespace:              ds.ObjectMeta.Namespace,
			Name:                   ds.ObjectMeta.Name,
			Connected:              false,
		}
		r.ParseLabelsPod(ds.ObjectMeta.Labels)

		rID := fmt.Sprintf("%s:%s", ds.ObjectMeta.Namespace, ds.ObjectMeta.Name)

		rl[rID] = r
	}

	return rl
}

func (api *API) getK8sResourceList(k8sResources map[string]*resource.KubernetesResource, connected bool) (map[string]*resource.KubernetesResource, []string) {
	if len(k8sResources) == 0 {
		msg.Info("No suitable resources found in kubernetes cluster")
		os.Exit(0)
	}

	var allIDs []string

	resources := make(map[string]*resource.KubernetesResource)

	for _, r := range k8sResources {
		// var status, subnet string
		var subnet string
		if r.Connected {
			if !connected {
				continue
			}
			// status = output.StrEnabled("■")
			// status = colors.Green("■")
			// status = "🖧"
			// status = "🟩"
			// status = "🟢"
			// subnet = fmt.Sprintf(" %s", output.StrNormal(fmt.Sprintf("%s:%s", r.NetID, r.VRFID)))
			// subnet = fmt.Sprintf("%s ", output.StrEnabled(fmt.Sprintf("%s:%s", r.NetID, r.VRFID)))
			subnet = fmt.Sprintf("[%s:%s] ", r.NetStatus.NetID, r.NetStatus.SubnetID)
		} else {
			if connected {
				continue
			}
			// status = output.StrDisabled("■")
			// status = colors.DarkRed("■")
			// status = "🟥"
			// status = "🔴"
			// status = "🟠"
		}

		// rID := fmt.Sprintf("%s %s: %s%s", status, r.Namespace, colors.DarkWhite(r.Name), subnet)
		// rID := fmt.Sprintf("%s %s: %s%s", status, r.Namespace, r.Name, subnet)
		// rID := fmt.Sprintf("%s: %s%s", r.Namespace, r.Name, subnet)
		rID := fmt.Sprintf("%s%s: %s", subnet, r.Namespace, r.Name)
		resources[rID] = r
		allIDs = append(allIDs, rID)
	}

	sort.Strings(allIDs)

	return resources, allIDs
}
