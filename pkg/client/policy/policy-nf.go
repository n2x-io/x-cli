package policy

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/subnet"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) NewNetworkFilter() {
	s := subnet.GetSubnet(false)

	nf := &topology.Filter{}

	setNetworkFilter(nf)

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	ss := output.Spinner()

	nnfr := &topology.NewNetworkFilterRequest{
		SubnetReq: &topology.SubnetReq{
			AccountID: s.AccountID,
			TenantID:  s.TenantID,
			NetID:     s.NetID,
			SubnetID:  s.SubnetID,
		},
		Index:       nf.Index,
		Description: nf.Description,
		SrcIPNet:    nf.SrcIPNet,
		DstIPNet:    nf.DstIPNet,
		Proto:       nf.Proto,
		DstPort:     nf.DstPort,
		Policy:      nf.Policy,
	}

	np, err := nxc.CreateNetworkFilter(context.TODO(), nnfr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to create network filter")
	}

	s.NetworkPolicy = np

	ss.Stop()

	// output.Show(np)
	Output().Show(s)
}

func (api *API) UpdateNetworkFilter() {
	s := subnet.GetSubnet(false)

	nf := getNetworkFilter(s.NetworkPolicy)

	setNetworkFilter(nf)

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	ss := output.Spinner()

	unfr := &topology.UpdateNetworkFilterRequest{
		SubnetReq: &topology.SubnetReq{
			AccountID: s.AccountID,
			TenantID:  s.TenantID,
			NetID:     s.NetID,
			SubnetID:  s.SubnetID,
		},
		NfID:        nf.NfID,
		Index:       nf.Index,
		Description: nf.Description,
		SrcIPNet:    nf.SrcIPNet,
		DstIPNet:    nf.DstIPNet,
		Proto:       nf.Proto,
		DstPort:     nf.DstPort,
		Policy:      nf.Policy,
	}

	np, err := nxc.UpdateNetworkFilter(context.TODO(), unfr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to update network filter")
	}

	s.NetworkPolicy = np

	ss.Stop()

	// output.Show(np)
	Output().Show(s)
}

func (api *API) DeleteNetworkFilter() {
	s := subnet.GetSubnet(false)

	nf := getNetworkFilter(s.NetworkPolicy)

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	ss := output.Spinner()

	dnfr := &topology.DeleteNetworkFilterRequest{
		SubnetReq: &topology.SubnetReq{
			AccountID: s.AccountID,
			TenantID:  s.TenantID,
			NetID:     s.NetID,
			SubnetID:  s.SubnetID,
		},
		NfID: nf.NfID,
	}

	np, err := nxc.DeleteNetworkFilter(context.TODO(), dnfr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to delete network filter")
	}

	s.NetworkPolicy = np

	ss.Stop()

	// output.Show(np)
	Output().Show(s)
}
