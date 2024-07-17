package output

import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/services/thirdParty"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func showIntegrations(i *thirdParty.Integrations) {
	if i == nil {
		return
	}

	if i.Github == nil &&
		i.Pagerduty == nil &&
		i.Slack == nil {
		return
	}

	output.SubTitleT2("Integrations")

	t := table.New()

	/*
		if i.Clickup != nil {
			if i.Clickup.Enabled {
				t.AddRow(colors.Black("ClickUp"), output.StrEnabled("enabled"))
				// t.AddRow(colors.Black("ClickUp API Key"), colors.DarkWhite(i.Clickup.ApiKey), colors.Black(" "))
				// t.AddRow(colors.Black("ClickUp Settings List URL"), colors.DarkWhite(i.Clickup.SettingsListURL), colors.Black(" "))
			} else {
				t.AddRow(colors.Black("ClickUp"), output.StrDisabled("not-configured"))
			}
			// t.AddRow()
		}
	*/
	if i.Github != nil {
		if i.Github.Enabled {
			t.AddRow(colors.Black("GitHub"), output.StrEnabled("enabled"))
			// t.AddRow(colors.Black("GitHub Access Token"), colors.DarkWhite(i.Github.AccessToken), colors.Black(" "))
			// t.AddRow(colors.Black("GitHub Webhooks Secret"), colors.DarkWhite(i.Github.WebhooksSecret), colors.Black(" "))
		} else {
			t.AddRow(colors.Black("GitHub"), output.StrDisabled("not-configured"))
		}
		// t.AddRow()
	}
	if i.Pagerduty != nil {
		if i.Pagerduty.Enabled {
			t.AddRow(colors.Black("PagerDuty"), output.StrEnabled("enabled"))
			// t.AddRow(colors.Black("PagerDuty Integration Key"), colors.DarkWhite(i.Pagerduty.IntegrationKey), colors.Black(" "))
		} else {
			t.AddRow(colors.Black("PagerDuty"), output.StrDisabled("not-configured"))
		}
		// t.AddRow()
	}
	if i.Slack != nil {
		if i.Slack.Enabled {
			t.AddRow(colors.Black("Slack"), output.StrEnabled("enabled"))
			if len(i.Slack.AlertsWebhook) > 0 {
				t.AddRow(colors.DarkWhite(" • Alert Notifications"), output.StrEnabled("enabled"))
				// t.AddRow(colors.Black("Slack Alert Notifications Webhook"), colors.DarkWhite(i.Slack.AlertsWebhook), colors.Black(" "))
			} else {
				t.AddRow(colors.DarkWhite(" • Alert Notifications"), output.StrDisabled("not-configured"))
			}
			if len(i.Slack.OpsWebhook) > 0 {
				t.AddRow(colors.DarkWhite(" • General Notifications & Reports"), output.StrEnabled("enabled"))
				// t.AddRow(colors.Black("Slack General Notifications Webhook"), colors.DarkWhite(i.Slack.OpsWebhook), colors.Black(" "))
			} else {
				t.AddRow(colors.DarkWhite(" • General Notifications & Reports"), output.StrDisabled("not-configured"))
			}
		} else {
			t.AddRow(colors.Black("Slack"), output.StrDisabled("not-configured"))
		}
		// t.AddRow()
	}
	/*
		if i.Crisp != nil {
			if i.Crisp.Enabled {
				t.AddRow(colors.Black("Crisp"), output.StrEnabled("enabled"))
				// t.AddRow(colors.Black("Crisp WebsiteID"), colors.DarkWhite(i.Crisp.WebsiteID), colors.Black(" "))
			} else {
				t.AddRow(colors.Black("Crisp"), output.StrDisabled("not-configured"))
			}
			// t.AddRow()
		}
	*/
	/*
		if i.Slack != nil {
			if i.Slack.Enabled {
				t.AddRow("Slack:", output.StrEnabled("enabled"))
				t.AddRow("Slack Bot User OAuth Access Token:", colors.Black(i.Slack.BotUserOAuthAccessToken), colors.Black(" "))
				t.AddRow("Slack Verification Token:", colors.Black(i.Slack.VerificationToken), colors.Black(" "))
				t.AddRow("Slack Signing Secret:", colors.Black(i.Slack.SigningSecret), colors.Black(" "))
			} else {
				t.AddRow("Slack:", output.StrDisabled("not-configured"))
			}
		}
	*/

	t.Render()
	fmt.Println()
}
