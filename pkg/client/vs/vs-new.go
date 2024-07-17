package vs

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/network"
	"n2x.dev/x-cli/pkg/client/node"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/ipnet"
	"n2x.dev/x-lib/pkg/n2x"
)

func (api *API) New() {
	nw := network.GetNetwork(false)

	nvsr := &topology.NewVSRequest{
		AccountID:   nw.AccountID,
		TenantID:    nw.TenantID,
		NetID:       nw.NetID,
		Name:        input.GetInput("VS Name:", "", "", input.ValidName),
		Description: input.GetInput("Description:", "", "", nil),
		LocationID:  nw.LocationID,
		Cname:       input.GetInput("Custom DNS CNAME:", "Fully Qualified Domain Name", "", input.ValidFQDN),
		ReqAuth:     input.GetConfirm("Authentication:", true),
		Proto:       topology.VSProto_PROTO_TCP_HTTPS,
		VSPort:      443,
		NodeAppSvcs: make([]*topology.NodeAppSvcReq, 0),
		Tags:        input.GetMultiInput("Tags:", "Tags separated by comma", nil, input.ValidTags),
	}

	if len(nvsr.Cname) > 0 {
		if err := ipnet.VSCNAMEIsValid(nvsr.Cname, nvsr.LocationID); err != nil {
			status.Error(err, "Unable to create virtual server")
		}
	}

	protocols := []string{"HTTP", "HTTPS"}
	var defaultPort string
	var rsProto topology.VSProto

	nodeAppSvcMap := make(map[string]*topology.NodeAppSvcReq, 0)

	nodeCount := 0
	msg1 := "Add node app svc to virtual server?"
	msgN := "Add another node app svc to virtual server (high-availability/lb group)?"
	for {
		msg := msg1
		if nodeCount > 0 {
			msg = msgN
		}
		ok := input.GetConfirm(msg, true)
		if !ok {
			break
		}

		n := node.GetNodeByTenant(false, n2x.Bool(true))

		if nw.TenantID != n.TenantID {
			status.Error(fmt.Errorf("vs/node tenantID mismatch"), "Unable to create virtual server")
		}

		if nw.NetID != n.Cfg.NetID {
			status.Error(fmt.Errorf("vs/node netID mismatch"), "Unable to create virtual server")
		}

		proto := input.GetSelect("Protocol:", "", protocols, survey.Required)

		switch proto {
		case "HTTP":
			rsProto = topology.VSProto_PROTO_TCP_HTTP
			defaultPort = "8080"
		case "HTTPS":
			rsProto = topology.VSProto_PROTO_TCP_HTTPS
			defaultPort = "443"
		}

		nodeAppSvcMap[n.NodeID] = &topology.NodeAppSvcReq{
			TenantID: n.TenantID,
			NetID:    n.Cfg.NetID,
			SubnetID: n.Cfg.SubnetID,
			NodeID:   n.NodeID,
			NodeName: n.Cfg.NodeName,
			// AppSvcName:        input.GetInput("App Svc Name:", "", "", input.ValidName),
			// AppSvcDescription: input.GetInput("Description:", "", "", nil),
			Proto:  rsProto,
			RSPort: getAppSvcPort(defaultPort),
		}

		nodeCount++
	}

	for _, nodeAppSvc := range nodeAppSvcMap {
		nvsr.NodeAppSvcs = append(nvsr.NodeAppSvcs, nodeAppSvc)
	}

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	vs, err := nxc.CreateVS(context.TODO(), nvsr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to create virtual server")
	}

	s.Stop()

	// output.Show(n)
	Output().Show(vs)
}

func getAppSvcPort(defaultPort string) int32 {
	port := input.GetInput("App Svc Port:", "", defaultPort, input.ValidPort)
	p, err := strconv.Atoi(port)
	if err != nil {
		status.Error(err, "Invalid port")
	}

	return int32(p)
}
