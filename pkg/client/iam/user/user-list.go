package user

func (api *API) List() {
	Output().List(users())
}
