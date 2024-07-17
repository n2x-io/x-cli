package node

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/subnet"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/n2x"
)

func (api *API) Connect() {
	n := GetNodeByTenant(false, nil)

	s := subnet.GetSubnet(false)

	ss := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	unr := &topology.UpdateNodeNetworkingRequest{
		NodeReq: &topology.NodeReq{
			AccountID: n.AccountID,
			TenantID:  n.TenantID,
			NodeID:    n.NodeID,
		},
		NetID:    s.NetID,
		SubnetID: s.SubnetID,
	}

	sr, err := nxc.UpdateNodeNetworking(context.TODO(), unr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to update node network configuration")
	}

	ss.Stop()

	// output.Show(sr)
	Output().Show(sr)
}

func (api *API) Disconnect() {
	n := GetNodeByTenant(false, n2x.Bool(true))

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	unr := &topology.UpdateNodeNetworkingRequest{
		NodeReq: &topology.NodeReq{
			AccountID: n.AccountID,
			TenantID:  n.TenantID,
			NodeID:    n.NodeID,
		},
		NetID:    "",
		SubnetID: "",
	}

	sr, err := nxc.UpdateNodeNetworking(context.TODO(), unr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to update node network configuration")
	}

	s.Stop()

	// output.Show(sr)
	Output().Show(sr)
}
