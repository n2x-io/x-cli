package output

import (
	"fmt"
	"os"
	"strings"
	"time"

	"n2x.dev/x-api-go/grpc/resources/ops"
	"n2x.dev/x-cli/pkg/auth"
	"n2x.dev/x-cli/pkg/client/event"
	"n2x.dev/x-cli/pkg/client/node"
	"n2x.dev/x-cli/pkg/client/tenant"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils"
	"n2x.dev/x-lib/pkg/utils/colors"
	"n2x.dev/x-lib/pkg/utils/msg"
)

func (api *API) Show(wf *ops.Workflow) {
	output.SectionHeader("Ops: Workflow Details")
	output.TitleT1("Workflow Information")

	t := table.New()

	// t.AddRow(colors.Black("Account ID"), colors.DarkWhite(wf.AccountID))
	t.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(wf.TenantID))
	t.AddRow(colors.Black("Project ID"), colors.DarkWhite(wf.ProjectID))
	t.AddRow(colors.Black("Workflow ID"), colors.DarkWhite(wf.WorkflowID))

	t.AddRow(colors.Black("Name"), colors.DarkWhite(wf.Name))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(wf.Description))

	// if wf.Owner != nil {
	// 	t.AddRow(colors.Black("Owner"), colors.DarkWhite(wf.Owner.UserID))
	// }

	if wf.Enabled {
		t.AddRow(colors.Black("Enabled"), output.StrEnabled("yes"))
	} else {
		t.AddRow(colors.Black("Enabled"), output.StrDisabled("no"))
	}

	if wf.Reviewed {
		t.AddRow(colors.Black("Reviewed"), output.StrEnabled("yes"))
	} else {
		t.AddRow(colors.Black("Reviewed"), output.StrDisabled("no"))
	}

	if wf.Approved {
		t.AddRow(colors.Black("Approved"), output.StrEnabled("yes"))
	} else {
		t.AddRow(colors.Black("Approved"), output.StrDisabled("no"))
	}

	t.Render()
	fmt.Println()

	if wf.Triggers != nil {
		output.SubTitleT2("Triggers")

		t = table.New()

		if wf.Triggers.Webhook {
			t.AddRow(colors.Black("Workflow Webhook"), output.StrEnabled("enabled"))
			t.Render()

			// /api/v1/accounts/{accountID}/webhooks/workflow/{wfToken}
			srv, err := auth.GetControllerEndpoint()
			if err == nil && len(srv) > 0 {
				host := strings.Split(srv, ":")[0]
				whURL := fmt.Sprintf("https://%s/api/v1/accounts/%s/webhooks/workflow/%s", host, wf.AccountID, wf.Token)
				fmt.Printf("\n%s %s\n", colors.Black("Workflow Webhook URL"), colors.DarkWhite(whURL))
			}
		} else {
			t.AddRow(colors.Black("Workflow Webhook"), output.StrDisabled("disabled"))
			t.Render()
		}

		fmt.Println()

		if wf.Triggers.Schedule != nil {
			t = table.New()

			if len(wf.Triggers.Schedule.Crontab) > 0 {
				if wf.Triggers.Schedule.Enabled {
					t.AddRow(colors.Black("Workflow Crontab"), output.StrEnabled("enabled"), colors.DarkWhite(wf.Triggers.Schedule.Crontab))
				} else {
					t.AddRow(colors.Black("Workflow Crontab"), output.StrDisabled("disabled"), colors.DarkWhite(wf.Triggers.Schedule.Crontab))
				}
			}

			if wf.Triggers.Schedule.DateTime != nil {
				timestamp, err := utils.GetDateTime(wf.Triggers.Schedule.DateTime)
				if err != nil {
					msg.Errorf("Unable to get workflow schedule: %v", err)
					os.Exit(1)
				}
				tm := time.Unix(timestamp, 0)

				if wf.Triggers.Schedule.Enabled {
					t.AddRow(colors.Black("Workflow Schedule"), output.StrEnabled("enabled"), colors.DarkWhite(tm.String()))
				} else {
					t.AddRow(colors.Black("Workflow Schedule"), output.StrDisabled("disabled"), colors.DarkWhite(tm.String()))
				}
			}

			t.Render()
			fmt.Println()
		}
	}

	if wf.Tasks != nil {
		output.SubTitleT2("Tasks")

		t = table.New()

		t.AddRow(colors.Black("Task Name"), colors.Black("Command"), colors.Black("Args"), colors.Black("UID"), colors.Black("GID"))
		t.AddRow(colors.Black("---------"), colors.Black("-------"), colors.Black("----"), colors.Black("---"), colors.Black("---"))

		for _, task := range wf.Tasks {
			if task.Command == nil {
				continue
			}
			taskName := colors.DarkWhite(task.Name)
			taskCmd := colors.DarkWhite(task.Command.Cmd)
			taskArgs := colors.DarkWhite(strings.Join(task.Command.Args, " "))
			taskUID := colors.DarkWhite(fmt.Sprintf("%d", task.Command.UID))
			taskGID := colors.DarkWhite(fmt.Sprintf("%d", task.Command.GID))
			t.AddRow(taskName, taskCmd, taskArgs, taskUID, taskGID)
		}

		t.Render()
		fmt.Println()
	}

	if len(wf.Targets) > 0 {
		output.SubTitleT2("Targets")

		s := output.Spinner()

		t = table.New()
		// t.Header(colors.Black("TENANT"), colors.Black("NETWORK"), colors.Black("SUBNET"), colors.Black("NODE"))
		// t.Header(colors.Black("Tenant"), colors.Black("Network"), colors.Black("Subnet"), colors.Black("Node"))

		// t.AddRow(colors.InvertedBlack("Tenant"), colors.InvertedBlack("Network"), colors.InvertedBlack("Subnet"), colors.InvertedBlack("Node"))
		// t.AddRow(colors.Black("Tenant"), colors.Black("Network"), colors.Black("Subnet"), colors.Black("Node"))
		// t.AddRow(colors.Black("------"), colors.Black("-------"), colors.Black("------"), colors.Black("----"))
		t.AddRow(colors.Black("Tenant"), colors.Black("Node"))
		t.AddRow(colors.Black("------"), colors.Black("----"))

		for _, nr := range wf.Targets {
			tenantName := tenant.FetchTenant(nr.TenantID).Name
			nodeName := node.FetchNode(nr).Cfg.NodeName
			t.AddRow(
				colors.DarkWhite(tenantName),
				// colors.DarkWhite(nr.NetID),
				// colors.DarkWhite(nr.SubnetID),
				// colors.DarkWhite(output.Fit(n.NodeID, 32)),
				colors.DarkWhite(output.Fit(nodeName, 32)),
				// colors.DarkWhite(" "),
			)
		}

		s.Stop()

		t.Render()
		fmt.Println()
	}

	if wf.Notify != nil {
		output.SubTitleT2("Notifications")

		t = table.New()
		// t.Header(colors.Black("RECIPIENT"), colors.Black("CHANNEL"))
		// t.Header(colors.Black("Recipient"), colors.Black("Channel"))

		// t.AddRow(colors.InvertedBlack("Recipient"), colors.InvertedBlack("Channel"))
		t.AddRow(colors.Black("Recipient"), colors.Black("Channel"))
		t.AddRow(colors.Black("---------"), colors.Black("-------"))

		if wf.Notify.Email != nil {
			for _, r := range wf.Notify.Email.Recipients {
				t.AddRow(colors.DarkWhite(r.Email), output.StrNormal("email"))
			}
		}

		if wf.Notify.Slack != nil {
			for _, r := range wf.Notify.Slack.Recipients {
				t.AddRow(colors.DarkWhite(r.Email), output.StrNormal("slack"))
			}
		}

		t.Render()
		fmt.Println()
	}

	if len(wf.Reviewers) > 0 {
		output.SubTitleT2("Reviewers")

		t = table.New()
		// t.Header(colors.Black("RESPONSIBLE"), colors.Black("VALIDATION"), colors.Black("TIMESTAMP"))
		// t.Header(colors.Black("Responsible"), colors.Black("Validation"), colors.Black("Timestamp"))

		// t.AddRow(colors.InvertedBlack("Responsible"), colors.InvertedBlack("Validation"), colors.InvertedBlack("Timestamp"))
		t.AddRow(colors.Black("Responsible"), colors.Black("Validation"), colors.Black("Timestamp"))
		t.AddRow(colors.Black("-----------"), colors.Black("----------"), colors.Black("---------"))

		for _, reviewer := range wf.Reviewers {
			for _, rby := range wf.ReviewedBy {
				if rby.Responsible.Email == reviewer {
					tm := time.Unix(rby.Timestamp, 0)
					if rby.Validated {
						t.AddRow(colors.DarkWhite(reviewer), output.StrEnabled("REVIEWED"), colors.DarkWhite(tm.String()))
					} else {
						t.AddRow(colors.DarkWhite(reviewer), output.StrDisabled("REJECTED"), colors.DarkWhite(tm.String()))
					}
				}
			}
		}

		t.Render()
		fmt.Println()
	}

	if len(wf.Approvers) > 0 {
		output.SubTitleT2("Approvers")

		t = table.New()
		// t.Header(colors.Black("RESPONSIBLE"), colors.Black("VALIDATION"), colors.Black("TIMESTAMP"))
		// t.Header(colors.Black("Responsible"), colors.Black("Validation"), colors.Black("Timestamp"))

		// t.AddRow(colors.InvertedBlack("Responsible"), colors.InvertedBlack("Validation"), colors.InvertedBlack("Timestamp"))
		t.AddRow(colors.Black("Responsible"), colors.Black("Validation"), colors.Black("Timestamp"))
		t.AddRow(colors.Black("-----------"), colors.Black("----------"), colors.Black("---------"))

		for _, approver := range wf.Approvers {
			for _, aby := range wf.ApprovedBy {
				if aby.Responsible.Email == approver {
					tm := time.Unix(aby.Timestamp, 0)
					if aby.Validated {
						t.AddRow(colors.DarkWhite(approver), output.StrEnabled("APPROVED"), colors.DarkWhite(tm.String()))
					} else {
						t.AddRow(colors.DarkWhite(approver), output.StrDisabled("REJECTED"), colors.DarkWhite(tm.String()))
					}
				}
			}
		}

		t.Render()
		fmt.Println()
	}

	if wf.EventMetrics != nil {
		event.Output().ShowMetrics(wf.EventMetrics)
	}
}
