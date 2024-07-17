package user

import (
	"context"
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

func GetUser(edit bool) *iam.User {
	ul := users()

	if len(ul) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var userOptID string
	usersOpts := make([]string, 0)
	users := make(map[string]*iam.User)

	for _, u := range ul {
		userOptID = u.UserID
		usersOpts = append(usersOpts, userOptID)
		users[userOptID] = u
	}

	sort.Strings(usersOpts)

	if edit {
		usersOpts = append(usersOpts, input.NewResource)
	}

	userOptID = input.GetSelect("User:", "", usersOpts, survey.Required)

	if userOptID == input.NewResource {
		return nil
	}

	vars.UserID = users[userOptID].UserID

	return users[userOptID]
}

func users() map[string]*iam.User {
	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	lr := &iam.ListUsersRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	users := make(map[string]*iam.User)

	for {
		ul, err := nxc.ListUsers(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list users")
		}

		for _, u := range ul.Users {
			users[u.UserID] = u
		}

		if len(ul.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = ul.Meta.NextPageToken
		} else {
			break
		}
	}

	return users
}
