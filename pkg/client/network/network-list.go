package network

func (api *API) List() {
	//output.Show(nl)
	Output().List(networks())
}
