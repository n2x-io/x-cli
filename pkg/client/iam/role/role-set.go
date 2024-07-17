package role

import (
	"context"

	"n2x.dev/x-api-go/grpc/common/empty"
	"n2x.dev/x-api-go/grpc/resources/iam"
	"n2x.dev/x-cli/pkg/client/account"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Set() {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	role := GetRole(true)
	if role != nil { // editing existing resource
		output.Choice("Edit RBAC Role")
	} else { // <new> resource
		output.Choice("New RBAC Role")

		role = &iam.Role{
			AccountID: a.AccountID,
			Users:     make(map[string]bool),
		}

		role.RoleID = input.GetInput("Role ID:", "", "", validRoleID)
	}

	perms, err := nxc.ListIAMPermissions(context.TODO(), &empty.Request{})
	if err != nil {
		status.Error(err, "Unable to get IAM permissions")
	}

	var permList []string
	for _, p := range perms.Permissions {
		permList = append(permList, p)
	}

	role.Permissions = input.GetMultiSelect("Permissions:", "", permList, role.Permissions, nil)

	s := output.Spinner()

	role, err = nxc.SetRole(context.TODO(), role)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set role")
	}

	s.Stop()

	Output().Show(role)
}
