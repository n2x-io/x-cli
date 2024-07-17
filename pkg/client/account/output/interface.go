package output

import (
	"n2x.dev/x-api-go/grpc/resources/account"
	"n2x.dev/x-api-go/grpc/resources/billing"
	user_output "n2x.dev/x-cli/pkg/client/iam/user/output"
	"n2x.dev/x-cli/pkg/output"
)

type Interface interface {
	show(a *account.Account)

	// Show(a *account.Account)

	Basic(a *account.Account)
	Admin(a *account.Account)
	Service(a *account.Account)
	Subscription(s *billing.Subscription, n int)
	Configuration(a *account.Account)
	Stats(a *account.Account)
}
type API struct{}

// func (api *API) Show(a *account.Account) {
// 	api.show(a)
// 	showServiceSubscription(a)
// 	showIntegrations(a.Integrations)
// 	// showTraffic(a.Traffic)
// 	showUsage(a.Usage)
// }

func (api *API) Basic(a *account.Account) {
	api.show(a)
}

func (api *API) Admin(a *account.Account) {
	// api.show(a)
	output.SectionHeader("Account Details")
	output.TitleT1("Account Admin Information")

	// output.SubTitleT1("Account Admin")
	user_output.ShowUser(a.Owner.Admin)
}

func (api *API) Service(a *account.Account) {
	api.show(a)
	showService(a)
}

func (api *API) Subscription(s *billing.Subscription, n int) {
	showServiceSubscription(s, n)
}

func (api *API) Configuration(a *account.Account) {
	api.show(a)
	showIntegrations(a.Integrations)
}

func (api *API) Stats(a *account.Account) {
	api.show(a)
	// showTraffic(a.Traffic)
	showUsage(a.Usage)
}
