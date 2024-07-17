package k8s

import (
	"context"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/subnet"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/k8s"
	"n2x.dev/x-lib/pkg/utils/msg"
)

func (api *API) DeleteGateway() {
	s := subnet.GetSubnet(false)
	if s == nil {
		msg.Alert("No subnet found.")
		msg.Alert("Please, configure at least one.")
		os.Exit(1)
	}

	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	ss := output.Spinner()

	resources, allIDs := api.getK8sResourceList(api.getKubernetesGateways(), true)

	ss.Stop()

	if len(allIDs) == 0 {
		msg.Info("No gateways connected found")
		os.Exit(1)
	}

	var selectedIDs []string

	selectMsg := "Delete n2x gateway"
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

	ss = output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	for _, rID := range selectedIDs {
		r, ok := resources[rID]
		if !ok {
			ss.Stop()
			msg.Error("Unable to parse response")
			os.Exit(1)
		}

		if !r.Connected {
			continue
		}

		ni, err := r.GetGatewayNodeInstance(nil, 0)
		if err != nil {
			ss.Stop()
			status.Error(err, "Unable to get kubernetes gateway node instance")
		}

		if err := k8s.API(api.kubeConfig).Objects().Node().DeleteGateway(r.Namespace, r.Name); err != nil {
			ss.Stop()
			status.Error(err, "Unable to delete kubernetes resources")
		}

		ngr := &topology.NodeGroupReq{
			AccountID:   ni.Node.AccountID,
			TenantID:    ni.Node.TenantID,
			NodeGroupID: ni.Node.NodeGroupID,
		}

		if _, err := nxc.DeleteNodeGroup(context.TODO(), ngr); err != nil {
			ss.Stop()
			status.Error(err, "Unable to delete kubernetes gateway nodeGroup")
		}
	}

	ss.Stop()

	fmt.Println()

	msg.Info("Kubernetes gateway deleted.")
}
