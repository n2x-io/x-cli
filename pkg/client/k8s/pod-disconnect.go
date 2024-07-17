package k8s

import (
	"context"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/k8s/resource"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/k8s"
	"n2x.dev/x-lib/pkg/utils/msg"
)

func (api *API) DisconnectPod() {
	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	s := output.Spinner()

	resources, allIDs := api.getK8sResourceList(api.getKubernetesPods(), true)

	s.Stop()

	if len(allIDs) == 0 {
		msg.Info("No pods connected found")
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

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	for _, rID := range selectedIDs {
		r, ok := resources[rID]
		if !ok {
			s.Stop()
			msg.Error("Unable to parse response")
			os.Exit(1)
		}

		if !r.Connected {
			continue
		}

		ni, err := r.GetPodNodeInstance(nil)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to get node instance")
		}

		switch r.KubernetesResourceType {
		case resource.KubernetesResourceTypeStatefulSet:
			if err := k8s.API(api.kubeConfig).Objects().Node().DisconnectStatefulSet(r.Namespace, r.Name, ni); err != nil {
				s.Stop()
				status.Error(err, "Unable to disconnect kubernetes statefulSet")
			}
		case resource.KubernetesResourceTypeDeployment:
			if err := k8s.API(api.kubeConfig).Objects().Node().DisconnectDeployment(r.Namespace, r.Name, ni); err != nil {
				s.Stop()
				status.Error(err, "Unable to disconnect kubernetes deployment")
			}
		case resource.KubernetesResourceTypeDaemonSet:
			if err := k8s.API(api.kubeConfig).Objects().Node().DisconnectDaemonSet(r.Namespace, r.Name, ni); err != nil {
				s.Stop()
				status.Error(err, "Unable to disconnect kubernetes daemonSet")
			}
		}

		ngr := &topology.NodeGroupReq{
			AccountID:   ni.Node.AccountID,
			TenantID:    ni.Node.TenantID,
			NodeGroupID: ni.Node.NodeGroupID,
		}

		if _, err := nxc.DeleteNodeGroup(context.TODO(), ngr); err != nil {
			s.Stop()
			status.Error(err, "Unable to delete kubernetes node group")
		}
	}

	s.Stop()

	fmt.Println()

	api.Pods()
}
