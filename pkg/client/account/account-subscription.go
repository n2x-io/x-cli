package account

import (
	"context"
	"fmt"
	"strings"

	"github.com/skratchdot/open-golang/open"
	"n2x.dev/x-api-go/grpc/resources/account"
	"n2x.dev/x-api-go/grpc/resources/billing"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/utils/colors"
	"n2x.dev/x-lib/pkg/utils/msg"
)

const serviceBusiness string = "business"

func (api *API) BillingPortal(a *account.Account) {
	if a == nil {
		a = FetchAccount()
	}

	s := output.Spinner()

	nxc, grpcConn := grpc.GetBillingAPIClient(true)
	defer grpcConn.Close()

	customerPortalRequest := &billing.CustomerPortalRequest{
		AccountID:  a.AccountID,
		CustomerID: a.Owner.Customer.StripeCustomerID,
	}

	r, err := nxc.GetCustomerPortal(context.TODO(), customerPortalRequest)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get customer portal URL")
	}

	if err := open.Start(r.URL); err != nil {
		s.Stop()
		status.Error(err, "Unable to open URL in your browser")
	}

	s.Stop()

	fmt.Printf("\n%s %s\n\n", colors.DarkWhite("🢂"), colors.Black("Opening Billing Portal in your browser..."))
}

func (api *API) Subscription(a *account.Account, interactive bool) {
	if a == nil {
		a = fetchAccountStats()
	}

	Output().Service(a)

	if a.Service == nil {
		return
	}

	requiresPaymentMethod := false

	if a.Service.Subscriptions == nil {
		msg.Warn("Your free account does not have a payment method configured")
		if !input.GetConfirm("[Limited-Time Offer] Add a payment method NOW and get $200 in FREE credits!", true) {
			fmt.Println()
			return
		}

		api.BillingPortal(a)

		return
	}

	for _, s := range a.Service.Subscriptions {
		if strings.ToLower(s.ServiceID) != serviceBusiness {
			continue
		}

		Output().Subscription(s, 0)

		if !interactive {
			return
		}

		switch s.LatestStripeInvoicePaymentIntentStatus {
		case "requires_payment_method":
			requiresPaymentMethod = true
		case "requires_action":
			if len(s.LatestStripeHostedInvoiceURL) > 0 {
				if input.GetConfirm(`Your bank/card issuer is requesting additional authentication
to authorize an ongoing payment.

Open the payment form now?`, true) {
					if err := open.Start(s.LatestStripeHostedInvoiceURL); err != nil {
						status.Error(err, "Unable to open URL in your browser")
					}

					fmt.Printf("\n%s %s\n\n", colors.DarkWhite("->"), colors.Black("Opening URL in your browser..."))
				} else {
					fmt.Println()
				}
			}
		}

		if !interactive {
			return
		}

		if requiresPaymentMethod {
			if input.GetConfirm("A failed payment attempt requires your attention, open the Billing Portal now?", true) {
				api.BillingPortal(a)
			} else {
				fmt.Println()
			}
		}
	}
}

/*
func (api *API) Subscription(a *account.Account, interactive bool) {
	if a == nil {
		a = fetchAccountStats()
	}

	Output().Service(a)

	if a.Service == nil {
		return
	}

	requiresPaymentMethod := false

	s := a.Service.Subscription

	Output().Subscription(s, 0)

	if !interactive {
		return
	}

	switch s.LatestStripeInvoicePaymentIntentStatus {
	case "requires_payment_method":
		requiresPaymentMethod = true
	case "requires_action":
		if len(s.LatestStripeHostedInvoiceURL) > 0 {
			if input.GetConfirm(`Your bank/card issuer is requesting additional authentication
  to authorize an ongoing payment.

  Open the payment form now?`, true) {
				if err := open.Start(s.LatestStripeHostedInvoiceURL); err != nil {
					status.Error(err, "Unable to open URL in your browser")
				}

				fmt.Printf("\n%s %s\n\n", colors.DarkWhite("🢂"), colors.Black("Opening URL in your browser..."))
			} else {
				fmt.Println()
			}
		}
	}

	if !interactive {
		return
	}

	if requiresPaymentMethod {
		if input.GetConfirm("A failed payment attempt requires your attention, open the Billing Portal now?", true) {
			api.BillingPortal(a)
		} else {
			fmt.Println()
		}
	} else {
		opt := !checkLimit(a)
		if !input.GetConfirm("Upgrade subscription now?", opt) {
			fmt.Println()
			return
		}

		api.BillingPortal(a)
	}
}
*/

/*
func (api *API) ApplyPromotion() {
	a := FetchAccount()

	api.Subscription(a, false)

	apr := &billing.ApplyPromotionRequest{
		AccountID:            a.AccountID,
		StripeSubscriptionID: a.Service.Subscription.StripeSubscriptionID,
		PromotionCode:        input.GetInput("Promotion Code:", "", "", survey.Required),
	}

	s := output.Spinner()

	nxc, grpcConn := grpc.GetBillingAPIClient(true)
	defer grpcConn.Close()

	sr, err := nxc.ApplyPromotion(context.TODO(), apr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to apply promotion code")
	}

	s.Stop()

	output.Show(sr)
}
*/

/*
func checkLimit(a *account.Account) bool {
	if a.Usage == nil {
		return true
	}

	if a.Usage.Limit == nil {
		return true
	}

	if !a.Usage.Limit.OverLimit {
		return true
	}

	msg.Warn("Your Free account is over limits, upgrade recommended.")

	return false
}
*/
