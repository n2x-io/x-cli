package output

import (
	"fmt"
	"sort"
	"strings"

	"n2x.dev/x-api-go/grpc/resources/events"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
	"n2x.dev/x-lib/pkg/utils/msg"
)

func (api *API) List(alerts map[string]*events.Alert) {
	output.SectionHeader("Alerts")
	output.TitleT1("Alert List")

	if len(alerts) == 0 {
		msg.Info("No records found")
		return
	}

	t := table.New()
	t.Header(colors.Black("LAST UPDATED / NODE NAME"), colors.Black("STATUS / COMPONENT"))

	t.SetRowLine("-")

	tmSort := make([]string, 0)
	tmSortUniq := make(map[string]struct{})
	for _, a := range alerts {
		tm := output.DatetimeMilli(a.LastUpdated)
		if _, ok := tmSortUniq[tm]; !ok {
			tmSortUniq[tm] = struct{}{}
			tmSort = append(tmSort, tm)
		}
	}
	sort.Strings(tmSort)

	for _, tm := range tmSort {
		for _, a := range alerts {
			if tm == output.DatetimeMilli(a.LastUpdated) {
				nodeName := colors.DarkWhite(output.Fit(a.NodeName, 36))
				component := colors.DarkWhite(output.Fit(strings.ToLower(a.Component), 32))

				var status string
				switch a.Status {
				case events.Status_TRIGGERED:
					status = colors.DarkRed("█")
				case events.Status_ACKNOWLEDGED:
					status = colors.DarkYellow("█")
				case events.Status_RESOLVED:
					status = colors.Green("█")
				}

				alertStatus := api.AlertStatus(a.Status)

				c1 := fmt.Sprintf("%s %s\n%s %s", status, colors.Black(tm), status, nodeName)
				c2 := fmt.Sprintf("%s\n%s", alertStatus, component)

				t.AddRow(c1, c2)
			}
		}
	}

	t.Render()
	fmt.Println()
}
