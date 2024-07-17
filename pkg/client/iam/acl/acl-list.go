package acl

func (api *API) List() {
	Output().List(ACLs())
}
