package acl

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/iam"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/tenant"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) Set() {
	t := tenant.GetTenant()

	nxc1, grpcConn1 := grpc.GetTopologyAPIClient()
	defer grpcConn1.Close()

	acl := GetACL(true)
	if acl != nil { // editing existing resource
		output.Choice("Edit RBAC ACL")
	} else { // <new> resource
		output.Choice("New RBAC ACL")

		acl = &iam.ACL{
			AccountID: t.AccountID,
			ACLID:     input.GetInput("ACL ID:", "", "", validACLID),
			Users:     make(map[string]bool, 0),
			Tags:      make([]string, 0),
		}
	}

	s := output.Spinner()

	req := &topology.TopologyRequest{
		AccountID: t.AccountID,
		TenantID:  t.TenantID,
	}

	ttags, err := nxc1.GetTopologyTags(context.TODO(), req)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get topology tags")
	}

	s.Stop()

	acl.Tags = input.GetMultiSelect("Tags:", "", ttags.Tags, acl.Tags, nil)

	s = output.Spinner()

	nxc2, grpcConn2 := grpc.GetIAMAPIClient()
	defer grpcConn2.Close()

	acl, err = nxc2.SetACL(context.TODO(), acl)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set ACL")
	}

	s.Stop()

	Output().Show(acl)
}
