package acl

func (api *API) Show() {
	Output().Show(GetACL(false))
}
