package output

import (
	"fmt"
	"time"

	"n2x.dev/x-api-go/grpc/resources/ops"
	"n2x.dev/x-cli/pkg/client/tenant"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) Show(tl *ops.TaskLog) {
	output.SectionHeader("Ops: TaskLog Details")
	output.TitleT1("TaskLog Information")

	t := table.New()

	// t.AddRow(colors.Black("Account ID"), colors.DarkWhite(tl.AccountID))
	t.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(tl.TenantID))
	t.AddRow(colors.Black("Project ID"), colors.DarkWhite(tl.ProjectID))
	t.AddRow(colors.Black("Workflow ID"), colors.DarkWhite(tl.WorkflowID))
	t.AddRow(colors.Black("TaskLog ID"), colors.DarkWhite(tl.TaskLogID))
	t.AddRow(colors.Black("Task Name"), colors.DarkWhite(tl.TaskName))
	t.AddRow(colors.Black("Task Description"), colors.DarkWhite(tl.TaskDescription))
	// t.AddRow("Target Node:", colors.DarkWhite(tl.NodeID))
	// t.AddRow()
	// t.AddRow(colors.DarkWhite("Target"))
	tenantName := tenant.FetchTenant(tl.TenantID).Name
	t.AddRow(colors.Black("Tenant"), colors.DarkWhite(tenantName))
	// t.AddRow(colors.Black("Network"), colors.DarkWhite(tl.NetID))
	// t.AddRow(colors.Black("Subnet"), colors.DarkWhite(tl.SubnetID))
	// t.AddRow(colors.Black("Node ID"), colors.DarkWhite(tl.NodeID))
	t.AddRow(colors.Black("Node"), colors.DarkWhite(tl.NodeName))

	if tl.Result != nil {
		switch tl.Result.Status {
		case ops.CommandResultStatus_EXECUTED:
			t.AddRow(colors.Black("Result"), output.StrOk("EXECUTED"))
		case ops.CommandResultStatus_FAILED:
			t.AddRow(colors.Black("Result"), output.StrError("FAILED"))
		}
	}
	if tl.Status != nil {
		tm := time.UnixMilli(tl.Status.Timestamp)
		t.AddRow(colors.Black("Timestamp"), colors.DarkWhite(tm.String()))
	}

	t.Render()
	fmt.Println()

	if len(tl.StdoutStderr) > 0 {
		output.SubTitleT2("Activity Log")

		if len(tl.StdoutStderr) > 0 {
			fmt.Println(colors.Black("-----BEGIN OUTPUT-----"))
			fmt.Printf("%s", tl.StdoutStderr)
			fmt.Println(colors.Black("-----END OUTPUT-----"))
			fmt.Println()
		}

		// if len(tl.Stdout) > 0 {
		// 	fmt.Println(colors.Black("-----BEGIN STDOUT-----"))
		// 	fmt.Printf("%s", tl.Stdout)
		// 	fmt.Println(colors.Black("-----END STDOUT-----"))
		// 	fmt.Println()
		// }

		// if len(tl.Stderr) > 0 {
		// 	fmt.Println(colors.Black("-----BEGIN STDERR-----"))
		// 	fmt.Printf("%s", tl.Stderr)
		// 	fmt.Println(colors.Black("-----END STDERR-----"))
		// 	fmt.Println()
		// }
	}
}
