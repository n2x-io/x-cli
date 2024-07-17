package role

func (api *API) List() {
	Output().List(Roles())
}
