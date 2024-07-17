package subnet

func (api *API) List() {
	Output().List(subnets())
}
