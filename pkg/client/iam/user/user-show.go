package user

func (api *API) Show() {
	Output().Show(GetUser(false))
}
