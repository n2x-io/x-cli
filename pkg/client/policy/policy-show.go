package policy

import (
	"n2x.dev/x-cli/pkg/client/subnet"
)

func (api *API) Show() {
	s := subnet.GetSubnet(false)

	Output().Show(s)
}
