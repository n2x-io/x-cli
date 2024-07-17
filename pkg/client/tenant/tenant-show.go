package tenant

func (api *API) Show() {
	Output().Show(GetTenant())
}
