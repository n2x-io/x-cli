package alert

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/events"
	"n2x.dev/x-api-go/grpc/resources/resource"
	tenant_pb "n2x.dev/x-api-go/grpc/resources/tenant"
	"n2x.dev/x-cli/pkg/client/tenant"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-cli/pkg/vars"
	"n2x.dev/x-lib/pkg/utils/msg"
)

func getAlert() *events.Alert {
	al := alerts()

	if len(al) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var alertOptID string
	alertsOpts := make([]string, 0)
	alerts := make(map[string]*events.Alert)

	for _, a := range al {
		tm := output.DatetimeMilli(a.LastUpdated)
		nodeName := output.Fit(a.NodeName, 32)
		// component := output.Fit(a.Component, 32)
		component := strings.ToLower(a.Component)
		l1 := fmt.Sprintf("%s %s [%s]", tm, nodeName, a.Status.String())
		// l2 := fmt.Sprintf("Component: %s", component)
		l2 := fmt.Sprintf("Component: %s | Group: %s | Class: %s", component, a.Group, a.Class)
		// l3 := fmt.Sprintf("Group: %s | Class: %s", a.Group, a.Class)
		// alertOptID = fmt.Sprintf("%s\n  %s\n  %s\n", l1, l2, l3)
		alertOptID = fmt.Sprintf("%s\n  %s\n", l1, l2)
		alertsOpts = append(alertsOpts, alertOptID)
		alerts[alertOptID] = a
	}

	sort.Strings(alertsOpts)

	alertOptID = input.GetSelect("Alert:", "", alertsOpts, survey.Required)

	vars.AlertID = alerts[alertOptID].AlertID

	return alerts[alertOptID]
}

func alerts() map[string]*events.Alert {
	t := tenant.GetTenant()

	s := output.Spinner()
	// defer s.Stop()

	nxc, grpcConn := grpc.GetMonitoringAPIClient()
	defer grpcConn.Close()

	lr := &events.ListAlertsRequest{
		Meta: &resource.ListRequest{},
		Tenant: &tenant_pb.TenantReq{
			AccountID: t.AccountID,
			TenantID:  t.TenantID,
		},
	}

	alerts := make(map[string]*events.Alert)

	for {
		al, err := nxc.ListAlerts(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list alerts")
		}

		for _, a := range al.Alerts {
			alerts[a.AlertID] = a
		}

		if len(al.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = al.Meta.NextPageToken
		} else {
			break
		}
	}

	s.Stop()

	return alerts
}
