package sg

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/iam"
	"n2x.dev/x-api-go/grpc/resources/resource"
	"n2x.dev/x-cli/pkg/client/account"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-cli/pkg/vars"
	"n2x.dev/x-lib/pkg/utils/msg"
)

var securityGroupsMap map[string]*iam.SecurityGroup = nil

func GetSecurityGroup(edit bool) *iam.SecurityGroup {
	sgl := SecurityGroups()

	if len(sgl) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var sgOptID string
	sgsOpts := make([]string, 0)
	sgs := make(map[string]*iam.SecurityGroup)

	for _, sg := range sgl {
		sgOptID = sg.SecurityGroupID
		sgsOpts = append(sgsOpts, sgOptID)
		sgs[sgOptID] = sg
	}

	sort.Strings(sgsOpts)

	if edit {
		sgsOpts = append(sgsOpts, input.NewResource)
	}

	sgOptID = input.GetSelect("Security Group:", "", sgsOpts, survey.Required)

	if sgOptID == input.NewResource {
		return nil
	}

	vars.SecurityGroupID = sgs[sgOptID].SecurityGroupID

	return sgs[sgOptID]
}

func SecurityGroups() map[string]*iam.SecurityGroup {
	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	lr := &iam.ListSecurityGroupsRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	securityGroups := make(map[string]*iam.SecurityGroup)

	for {
		sgl, err := nxc.ListSecurityGroups(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list security-groups")
		}

		for _, sg := range sgl.SecurityGroups {
			securityGroups[sg.SecurityGroupID] = sg
		}

		if len(sgl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = sgl.Meta.NextPageToken
		} else {
			break
		}
	}

	return securityGroups
}

func validSecurityGroupID(val interface{}) error {
	if err := input.ValidID(val); err != nil {
		return err
	}

	sgID := val.(string)

	if securityGroupsMap == nil {
		securityGroupsMap = SecurityGroups()
	}

	if _, ok := securityGroupsMap[sgID]; ok {
		return fmt.Errorf("security group %s already exist", sgID)
	}

	return nil
}
