package subnet

func (api *API) Show() {
	Output().Show(GetSubnet(false))
}
