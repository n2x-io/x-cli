package output

import (
	"fmt"
	"time"

	"n2x.dev/x-api-go/grpc/resources/events"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) Show(a *events.Alert) {
	// fmt.Println() // extra line needed due to the long select option above
	output.SectionHeader("Alerts")
	output.TitleT1("Alert Information")

	t := table.New()
	t.BulkData([][]string{
		// {colors.Black("Account ID"), colors.DarkWhite(a.AccountID)},
		{colors.Black("Tenant ID"), colors.DarkWhite(a.TenantID)},
		{colors.Black("Alert ID"), colors.DarkWhite(output.Fit(a.AlertID, 62))},
		{colors.Black("Timestamp"), colors.DarkWhite(time.UnixMilli(a.Timestamp).String())},
		// {colors.Black("Network ID"), colors.DarkWhite(a.NetID)},
		// {colors.Black("Subnet ID"), colors.DarkWhite(a.SubnetID)},
		{colors.Black("Node ID"), colors.DarkWhite(a.NodeID)},
		{colors.Black("Node Name"), colors.DarkWhite(a.NodeName)},
		{colors.Black("Class"), colors.DarkWhite(a.Class.String())},
		{colors.Black("Group"), colors.DarkWhite(a.Group.String())},
		{colors.Black("Component"), colors.DarkWhite(output.Fit(a.Component, 62))},
		{colors.Black("Severity"), alertSeverity(a.Severity)},
		{colors.Black("Status"), api.AlertStatus(a.Status)},
		{colors.Black("Last Updated"), colors.DarkWhite(time.UnixMilli(a.LastUpdated).String())},
		// {colors.Black("Summary"), colors.DarkWhite(a.Summary)},
	})

	t.AddRow(colors.Black("Summary"), colors.DarkWhite(output.Fit(a.Summary, 62)))

	// t.SetAutoWrapText(true)
	// t.SetReflowDuringAutoWrap(true)

	var alertDetails string
	n := 0
	for k, v := range a.CustomDetails {
		line := fmt.Sprintf("%s: %s", k, v)
		if n == len(a.CustomDetails) {
			alertDetails = fmt.Sprintf("%s%s", alertDetails, line)
		} else {
			alertDetails = fmt.Sprintf("%s%s\n", alertDetails, line)
		}
		n++
	}
	t.AddRow(colors.Black("Details"), colors.DarkWhite(alertDetails))

	t.Render()

	// if a.Comments == nil {
	// 	return
	// }

	// output.SubTitleT2("Notes")

	// alertComments(a)
}

/*
func alertComments(a *events.Alert) {
	if a.Comments == nil {
		return
	}

	var tmSort []int64
	commentMap := make(map[int64]*events.AlertComment, 0)

	for _, c := range a.Comments {
		tmSort = append(tmSort, c.Timestamp)
		commentMap[c.Timestamp] = c
	}
	sort.Slice(tmSort, func(i, j int) bool { return tmSort[i] < tmSort[j] })

	t := table.New()
	// t.SetRowLine("─")

	sep := strings.Repeat("─", 50)
	n := len(a.Comments)
	var m int

	for _, tm := range tmSort {
		c := commentMap[tm]
		m++
		userNickname := strings.Split(c.UserEmail, "@")[0]
		author := output.UserLocal(fmt.Sprintf("%s @%s", userNickname, a.AccountID))
		timestamp := time.Unix(tm, 0).String()
		hdr := fmt.Sprintf("%s %s", colors.Black(timestamp), author)
		t.AddRow(hdr)
		t.AddRow(fmt.Sprintf("%s", c.Text))
		if m < n {
			t.AddRow(colors.Black(sep))
		}
	}

	t.Render()
	fmt.Println()
}
*/

func (api *API) AlertStatus(s events.Status) string {
	switch s {
	case events.Status_TRIGGERED:
		return output.StrError(s.String())
	case events.Status_ACKNOWLEDGED:
		return output.StrWarning(s.String())
	case events.Status_RESOLVED:
		return output.StrOk(s.String())
	}

	return output.StrInactive(s.String())
}

func alertSeverity(s events.Severity) string {
	switch s {
	case events.Severity_INFO:
		return output.StrBlue(s.String())
	case events.Severity_WARNING:
		return output.StrWarning(s.String())
	case events.Severity_ERROR:
		return output.StrError(s.String())
	case events.Severity_CRITICAL:
		return output.StrMagenta(s.String())
	}

	return output.StrInactive(s.String())
}
