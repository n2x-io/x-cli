package tasklog

func (api *API) Show() {
	Output().Show(GetTaskLog())
}
