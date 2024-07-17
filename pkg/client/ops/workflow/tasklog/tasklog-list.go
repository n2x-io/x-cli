package tasklog

func (api *API) List() {
	Output().List(taskLogs())
}
