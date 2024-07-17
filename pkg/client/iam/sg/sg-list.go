package sg

func (api *API) List() {
	Output().List(SecurityGroups())
}
