package policy

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/subnet"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) SetDefault() {
	s := subnet.GetSubnet(false)

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	usr := &topology.UpdateSubnetRequest{
		AccountID:     s.AccountID,
		TenantID:      s.TenantID,
		NetID:         s.NetID,
		SubnetID:      s.SubnetID,
		Description:   s.Description,
		DefaultPolicy: subnet.GetSecurityPolicy("Default Security Policy:"),
	}

	ss := output.Spinner()

	s, err := nxc.UpdateSubnet(context.TODO(), usr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to set network policy")
	}

	ss.Stop()

	// output.Show(s)
	Output().Show(s)
}
