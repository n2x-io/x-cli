package tenant

func (api *API) List() {
	Output().List(Tenants())
}
