package account

import (
	"context"
	"fmt"
	"os"

	"n2x.dev/x-api-go/grpc/resources/account"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/utils/msg"
)

func (api *API) Cancel() {
	a := FetchAccount()

	nxc, grpcConn := grpc.GetAccountAPIClient(true)
	defer grpcConn.Close()

	confirmCancelation()

	s := output.Spinner()

	ar := &account.AccountReq{
		AccountID: a.AccountID,
	}

	sr, err := nxc.CancelAccount(context.TODO(), ar)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to cancel account")
	}

	s.Stop()

	output.Show(sr)
}

func confirmCancelation() {
	msg.Alert("You are about to cancel your n2x account.")
	msg.Alert("All its remaining resources or configurations (rbac, users, etc) will be deleted.")
	msg.Alert("This action is irreversible, please, double check.")

	if !input.GetConfirm("Confirm account cancelation?", false) {
		fmt.Println()
		os.Exit(0)
	}

	if !input.GetConfirm("Last chance. Confirm account cancelation?", false) {
		fmt.Println()
		os.Exit(0)
	}
}
