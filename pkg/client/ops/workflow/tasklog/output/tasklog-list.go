package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/ops"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) List(taskLogs []*ops.TaskLog) {
	output.SectionHeader("Ops: TaskLogs")
	output.TitleT1("TaskLog List")

	t := table.New()
	t.Header(colors.Black("TASK"), colors.Black("TIMESTAMP"), colors.Black("TARGET NODE"))

	for _, tl := range taskLogs {
		tm := tl.Status.Timestamp
		t.AddRow(colors.DarkWhite(tl.TaskName), output.DatetimeMilli(tm), tl.NodeName)
	}

	t.Render()
	fmt.Println()
}
