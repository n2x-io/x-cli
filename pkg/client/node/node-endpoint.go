package node

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/n2x"
)

func (api *API) ShowEndpoint() {
	n := GetNodeByTenant(false, n2x.Bool(true))
	eID := GetEndpoint(n).EndpointID

	for endpointID, e := range n.Endpoints {
		if endpointID == eID {
			// output.Show(e)
			Output().ShowEndpoint(e)
		}
	}
}

func (api *API) DeleteEndpoint() {
	n := GetNodeByTenant(false, n2x.Bool(true))
	eID := GetEndpoint(n).EndpointID

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	er := &topology.EndpointRequest{
		NodeReq: &topology.NodeReq{
			AccountID: n.AccountID,
			TenantID:  n.TenantID,
			NodeID:    n.NodeID,
		},
		EndpointID: eID,
	}

	sr, err := nxc.DeleteNetworkEndpoint(context.TODO(), er)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete node network endpoint")
	}

	s.Stop()

	output.Show(sr)
}
