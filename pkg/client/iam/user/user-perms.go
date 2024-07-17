package user

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/iam"
	"n2x.dev/x-cli/pkg/client/iam/acl"
	"n2x.dev/x-cli/pkg/client/iam/role"
	"n2x.dev/x-cli/pkg/client/iam/sg"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) SetPermissions() {
	u := GetUser(false)

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	supr := &iam.SetUserPermissionsRequest{
		AccountID: u.AccountID,
		UserID:    u.UserID,
		RBAC:      setUserRBAC(u),
	}

	s := output.Spinner()

	u, err := nxc.SetUserPermissions(context.TODO(), supr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set user permissions")
	}

	s.Stop()

	Output().Show(u)
}

func setUserRBAC(u *iam.User) *iam.UserRBAC {
	rbac := u.RBAC

	rbac.SecurityGroups = input.GetMultiSelect("Security Groups:", "", getSecurityGroups(u.AccountID), rbac.SecurityGroups, nil)

	rbac.ACLs = input.GetMultiSelect("ACLs:", "", getACLs(u.AccountID), rbac.ACLs, nil)

	rbac.Roles = input.GetMultiSelect("Roles:", "", getRoles(u.AccountID), rbac.Roles, nil)

	return rbac
}

func getACLs(accountID string) []string {
	acll := acl.ACLs()

	acls := make([]string, 0)

	for _, acl := range acll {
		if acl.AccountID != accountID {
			continue
		}
		acls = append(acls, acl.ACLID)
	}

	return acls
}

func getRoles(accountID string) []string {
	rl := role.Roles()

	roles := make([]string, 0)

	for _, r := range rl {
		if r.AccountID != accountID {
			continue
		}
		roles = append(roles, r.RoleID)
	}

	return roles
}

func getSecurityGroups(accountID string) []string {
	sgl := sg.SecurityGroups()

	sgs := make([]string, 0)

	for _, sg := range sgl {
		if sg.AccountID != accountID {
			continue
		}
		sgs = append(sgs, sg.SecurityGroupID)
	}

	return sgs
}
