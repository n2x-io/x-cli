package resource

import (
	"context"
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/grpc"
)

func (r *KubernetesResource) GetPodNodeInstance(s *topology.Subnet) (*topology.NodeInstance, error) {
	if len(r.Name) == 0 || len(r.Namespace) == 0 {
		return nil, fmt.Errorf("missing kubernetes resource metadata")
	}

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	if len(r.Labels.AccountID) > 0 && len(r.Labels.TenantID) > 0 && len(r.Labels.NodeGroupID) > 0 {
		// nodeGroup already exists
		ngr := &topology.NodeGroupReq{
			AccountID:   r.Labels.AccountID,
			TenantID:    r.Labels.TenantID,
			NodeGroupID: r.Labels.NodeGroupID,
		}

		return nxc.GetNodeGroupInstance(context.TODO(), ngr)
	}

	// nodeGroup does not exist

	if s == nil {
		return nil, fmt.Errorf("missing subnet")
	}

	nnr := &topology.NewNodeRequest{
		AccountID:   s.AccountID,
		TenantID:    s.TenantID,
		NetID:       s.NetID,
		SubnetID:    s.SubnetID,
		NodeName:    r.Name,
		Description: fmt.Sprintf("[k8s-pod] %s", r.Name),
		Type:        topology.NodeType_K8S_POD,
		ReplicaSet:  true,
		KubernetesAttrs: &topology.KubernetesAttrs{
			Namespace: r.Namespace,
			Name:      fmt.Sprintf("%s-n2x-node", r.Name),
			// PersistentVolume: false,
		},
	}

	return nxc.CreateKubernetesPod(context.TODO(), nnr)
}
