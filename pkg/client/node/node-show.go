package node

func (api *API) Show() {
	Output().Show(GetNodeByTenant(false, nil))
}
