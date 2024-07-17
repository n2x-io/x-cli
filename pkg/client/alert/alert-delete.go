package alert

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"n2x.dev/x-api-go/grpc/resources/events"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/utils/colors"
	"n2x.dev/x-lib/pkg/utils/msg"
)

func (api *API) Delete() {
	api.singleDelete()
	// api.multiDelete()
}

func (api *API) multiDelete() {
	al := alerts()

	var alertOptID string
	alertsOpts := make([]string, 0)
	alerts := make(map[string]*events.Alert)

	sep := colors.Black(strings.Repeat("-", 70))

	for _, a := range al {
		tm := colors.DarkYellow(output.Datetime(a.LastUpdated))
		nodeID := colors.DarkWhite(output.Fit(a.NodeID, 32))
		status := Output().AlertStatus(a.Status)
		component := colors.DarkWhite(output.Fit(strings.ToLower(a.Component), 52))
		group := fmt.Sprintf("%s %s", colors.Black("Group"), colors.DarkWhite(a.Group.String()))
		class := fmt.Sprintf("%s %s", colors.Black("Class"), colors.DarkWhite(a.Class.String()))
		l1 := fmt.Sprintf("%s %s %s", tm, nodeID, status)
		l2 := fmt.Sprintf("%s %s", colors.Black("Component"), component)
		l3 := fmt.Sprintf("%s %s", group, class)
		alertOptID = fmt.Sprintf("\n      %s\n      %s\n      %s\n%s", l1, l2, l3, sep)
		alertsOpts = append(alertsOpts, alertOptID)
		alerts[alertOptID] = a
	}

	sort.Strings(alertsOpts)

	alertsToDelete := input.GetMultiSelect("Delete Alerts:", "", alertsOpts, []string{}, nil)

	if len(alertsToDelete) == 0 {
		msg.Info("No alerts deleted.")
		return
	}

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetMonitoringAPIClient()
	defer grpcConn.Close()

	for _, alertOptID := range alertsToDelete {
		a := alerts[alertOptID]

		ar := &events.AlertReq{
			AccountID: a.AccountID,
			AlertID:   a.AlertID,
			TenantID:  a.TenantID,
			// NetID:     a.NetID,
			// SubnetID:  a.SubnetID,
			// NodeID: a.NodeID,
		}

		if _, err := nxc.DeleteAlert(context.TODO(), ar); err != nil {
			s.Stop()
			status.Error(err, "Unable to delete alert")
		}
	}

	s.Stop()

	msg.Info("Done.")
}

func (api *API) singleDelete() {
	a := getAlert()

	nxc, grpcConn := grpc.GetMonitoringAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	ar := &events.AlertReq{
		AccountID: a.AccountID,
		AlertID:   a.AlertID,
		TenantID:  a.TenantID,
		// NetID:     a.NetID,
		// SubnetID:  a.SubnetID,
		NodeID: a.NodeID,
	}

	sr, err := nxc.DeleteAlert(context.TODO(), ar)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete alert")
	}

	s.Stop()

	output.Show(sr)
}
