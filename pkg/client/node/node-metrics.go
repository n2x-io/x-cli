package node

func (api *API) Metrics() {
	Output().Metrics(GetNodeByTenant(false, nil))
}
