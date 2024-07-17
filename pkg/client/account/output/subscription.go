package output

import (
	"fmt"
	"time"

	"n2x.dev/x-api-go/grpc/resources/billing"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func showServiceSubscription(s *billing.Subscription, n int) {
	t := table.New()

	svcID := fmt.Sprintf("%s%s%s %s", colors.DarkMagenta("["), colors.Magenta(fmt.Sprintf("%02d", n+1)), colors.DarkMagenta("]"), colors.Black("Service ID"))
	t.AddRow(svcID, colors.DarkWhite(s.ServiceID))

	if s.Discount != nil {
		if len(s.Discount.PercentOff) > 0 {
			percentOff := fmt.Sprintf("%s%%", s.Discount.PercentOff)
			t.AddRow(colors.Black("     Promotional Discount"), colors.Green(percentOff))
		}
	}

	// var openInvoice bool
	var sstatus string
	switch s.Status {
	case billing.SubscriptionStatus_UNKNOWN:
		sstatus = output.StrInactive("unknown")
	case billing.SubscriptionStatus_NONE:
		sstatus = output.StrInactive("none")
	case billing.SubscriptionStatus_TRIALING:
		sstatus = output.StrEnabled("trial")
	case billing.SubscriptionStatus_ACTIVE:
		sstatus = output.StrEnabled("active")
	case billing.SubscriptionStatus_INCOMPLETE:
		// openInvoice = true
		sstatus = output.StrWarning("incomplete")
	case billing.SubscriptionStatus_INCOMPLETE_EXPIRED:
		// openInvoice = true
		sstatus = output.StrError("incomplete-expired")
	case billing.SubscriptionStatus_PAST_DUE:
		// openInvoice = true
		sstatus = output.StrWarning("past-due")
	case billing.SubscriptionStatus_CANCELED:
		// openInvoice = true
		sstatus = output.StrError("canceled")
	case billing.SubscriptionStatus_UNPAID:
		// openInvoice = true
		sstatus = output.StrError("unpaid")
	default:
		sstatus = output.StrInactive("n/a")
	}

	t.AddRow(colors.Black("     Status"), sstatus)

	tm := time.Unix(s.StartDate, 0).String()
	t.AddRow(colors.Black("     Start Date"), colors.DarkWhite(tm))
	if s.EndDate != 0 {
		tm = time.Unix(s.EndDate, 0).String()
		t.AddRow(colors.Black("     End Date"), colors.DarkWhite(tm))
	}

	// if s.TrialStartDate != 0 {
	// 	tm = time.Unix(s.TrialStartDate, 0).String()
	// 	t.AddRow(colors.Black("     Trial Period Start"), colors.DarkWhite(tm))
	// 	tm = time.Unix(s.TrialEndDate, 0).String()
	// 	t.AddRow(colors.Black("     Trial Period End"), colors.DarkWhite(tm))
	// }

	tm = time.Unix(s.CurrentPeriodStart, 0).String()
	t.AddRow(colors.Black("     Current Period Start"), colors.DarkWhite(tm))
	tm = time.Unix(s.CurrentPeriodEnd, 0).String()
	t.AddRow(colors.Black("     Current Period End"), colors.DarkWhite(tm))

	if s.CancelAtPeriodEnd {
		t.AddRow(colors.Black("     Cancel at Period End"), output.StrWarning("yes"))
		tm = time.Unix(s.CancelationDate, 0).String()
		t.AddRow(colors.Black("     Cancelation Date"), colors.DarkWhite(tm))
	}

	tm = time.Unix(s.CreationDate, 0).String()
	t.AddRow(colors.Black("     Creation Date"), colors.DarkWhite(tm))

	tm = time.Unix(s.LastModified, 0).String()
	t.AddRow(colors.Black("     Last Modified"), colors.DarkWhite(tm))

	t.Render()
	fmt.Println()

	// if openInvoice && len(s.LatestStripeHostedInvoiceURL) > 0 {
	// 	if input.GetConfirm("Detected unpaid invoice, open payment form now?", true) {
	// 		if err := open.Start(s.LatestStripeHostedInvoiceURL); err != nil {
	// 			status.Error(err, "Unable to open URL in your browser")
	// 		}

	// 		fmt.Printf("\n%s %s\n\n", colors.DarkWhite("🢂"), colors.Black("Opening URL in your browser..."))
	// 	} else {
	// 		fmt.Println()
	// 	}
	// }
}
