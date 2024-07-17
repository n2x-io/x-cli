package node

func (api *API) ListByTenant() {
	// output.Show(nl)
	Output().List(nodesByTenant())
}

func (api *API) ListBySubnet() {
	// output.Show(nl)
	Output().List(nodesBySubnet())
}
