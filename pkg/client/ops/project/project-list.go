package project

func (api *API) List() {
	// output.Show(pl)
	Output().List(projects())
}
