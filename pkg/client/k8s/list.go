package k8s

func (api *API) Services() {
	Output().List(api.getKubernetesServicesAnnotations())
}

func (api *API) Pods() {
	Output().List(api.getKubernetesPods())
}
