package k8s

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/k8s"
	"n2x.dev/x-lib/pkg/utils/msg"
)

func (api *API) DisconnectService() {
	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	s := output.Spinner()

	resources, allIDs := api.getK8sResourceList(api.getKubernetesServicesAnnotations(), true)

	s.Stop()

	if len(allIDs) == 0 {
		msg.Info("No services connected found")
		os.Exit(1)
	}

	var selectedIDs []string

	selectMsg := "Disconnect from n2x"
	if err := survey.AskOne(
		&survey.MultiSelect{
			Message:  selectMsg,
			Options:  allIDs,
			PageSize: 10,
		},
		&selectedIDs,
		survey.WithIcons(input.SurveySetIcons),
	); err != nil {
		status.Error(err, "Unable to get response")
	}

	s = output.Spinner()

	for _, rID := range selectedIDs {
		r, ok := resources[rID]
		if !ok {
			msg.Error("Unable to parse response")
			os.Exit(1)
		}

		annt := removedAnnotations()
		if err := k8s.API(api.kubeConfig).Resources().Service().RemoveAnnotations(r.Namespace, r.Name, annt); err != nil {
			status.Error(err, "Unable to remove kubernetes service annotations")
		}
	}

	s.Stop()

	fmt.Println()

	api.Services()
}

func removedAnnotations() map[string]string {
	return map[string]string{
		"n2x.io/account": "",
		"n2x.io/tenant":  "",
		"n2x.io/network": "",
		"n2x.io/subnet":  "",
		"n2x.io/dnsName": "",
		"n2x.io/ipv4":    "",
	}
}
