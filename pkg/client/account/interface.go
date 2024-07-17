package account

import (
	"n2x.dev/x-api-go/grpc/resources/account"
	"n2x.dev/x-cli/pkg/client/account/output"
)

type Interface interface {
	Show()
	// Settings()
	Cancel()
	Subscription(a *account.Account, interactive bool)
	// ApplyPromotion()
	BillingPortal(a *account.Account)
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
