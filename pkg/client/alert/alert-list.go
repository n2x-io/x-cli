package alert

func (api *API) List() {
	Output().List(alerts())
}
