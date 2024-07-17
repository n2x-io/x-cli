package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"n2x.dev/x-lib/pkg/utils/colors"
)

type Table struct {
	table *tablewriter.Table
}

func New() *Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAutoMergeCells(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	// table.SetRowSeparator(colors.Black("─"))
	table.SetRowSeparator(colors.Black("-"))
	table.SetHeaderLine(false)
	table.SetRowLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	// t.table.SetColumnColor(
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
	// )

	return &Table{table: table}
}

func (t *Table) SetAutoWrapText(value bool) {
	t.table.SetAutoWrapText(value)
}

func (t *Table) SetReflowDuringAutoWrap(value bool) {
	t.table.SetReflowDuringAutoWrap(value)
}

func (t *Table) SetRowLine(char string) {
	t.table.SetRowSeparator(colors.Black(char))
	t.table.SetRowLine(true)
}

func (t *Table) UnsetRowLine() {
	t.table.SetRowLine(false)
}

func (t *Table) SetColWidth(width int) {
	t.table.SetColWidth(width)
}

func (t *Table) Header(header ...string) {
	t.table.SetAutoFormatHeaders(false)
	t.table.SetHeader(header)
	t.table.SetHeaderLine(true)

	// t.table.SetAutoWrapText(true)
	// t.table.SetReflowDuringAutoWrap(true)
	// t.table.SetRowLine(true)
}

func (t *Table) Footer(footer ...string) {
	t.table.SetFooter(footer)
}

func (t *Table) Caption(caption string) {
	t.table.SetCaption(true, caption)
}

func (t *Table) AddRow(row ...string) {
	t.table.Append(row)
}

func (t *Table) BulkData(data [][]string) {
	t.table.AppendBulk(data) // Add Bulk Data
}

func (t *Table) Render() {
	t.table.Render()
}
