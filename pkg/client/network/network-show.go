package network

func (api *API) Show() {
	Output().Show(GetNetwork(false))
}
