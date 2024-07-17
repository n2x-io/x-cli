package role

func (api *API) Show() {
	Output().Show(GetRole(false))
}
