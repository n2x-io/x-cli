package workflow

func (api *API) Show() {
	Output().Show(GetWorkflow())
}
