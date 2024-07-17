package workflow

func (api *API) List() {
	Output().List(workflows())
}
