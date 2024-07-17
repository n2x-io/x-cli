package thirdParty

import "n2x.dev/x-api-go/grpc/resources/services/thirdParty"

type Interface interface {
	// ClickUp(c *thirdParty.ClickUp) *thirdParty.ClickUp
	GitHub(gh *thirdParty.GitHub) *thirdParty.GitHub
	PagerDuty(pd *thirdParty.PagerDuty) *thirdParty.PagerDuty
	Slack(slck *thirdParty.Slack) *thirdParty.Slack
}
type API struct{}

func Setup() Interface {
	return &API{}
}
