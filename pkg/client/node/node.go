package node

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/resource"
	tenant_pb "n2x.dev/x-api-go/grpc/resources/tenant"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/subnet"
	"n2x.dev/x-cli/pkg/client/tenant"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-cli/pkg/vars"
	"n2x.dev/x-lib/pkg/n2x"
	"n2x.dev/x-lib/pkg/utils/msg"
)

func GetNodeByTenant(edit bool, connected *bool) *topology.Node {
	nl := nodesByTenant()

	return getNode(nl, edit, connected)
}

func GetNodeBySubnet(edit bool) *topology.Node {
	nl := nodesBySubnet()

	return getNode(nl, edit, n2x.Bool(true))
}

func GetEndpoint(n *topology.Node) *topology.Endpoint {
	var eID string
	var endpoints []string

	for endpointID := range n.Endpoints {
		endpoints = append(endpoints, endpointID)
	}

	sort.Strings(endpoints)

	if len(endpoints) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	eID = input.GetSelect("Endpoints:", "", endpoints, survey.Required)

	return n.Endpoints[eID]
}

func getNode(nl map[string]*topology.Node, edit bool, connected *bool) *topology.Node {
	if len(nl) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var nodeOptID string
	nodesOpts := make([]string, 0)
	nodes := make(map[string]*topology.Node)

	var getAll, getConnected bool
	if connected != nil {
		getConnected = *connected
	} else {
		getAll = true
	}

	for nodeName, n := range nl {
		if !getAll {
			if getConnected && len(n.Cfg.SubnetID) == 0 {
				continue
			}
			if !getConnected && len(n.Cfg.SubnetID) > 0 {
				continue
			}
		}
		// nodeOptID = nodeName
		nodeOptID = fmt.Sprintf("[%s] %s", nodeName, n.Cfg.Description)
		nodesOpts = append(nodesOpts, nodeOptID)
		nodes[nodeOptID] = n
	}

	sort.Strings(nodesOpts)

	if edit {
		nodesOpts = append(nodesOpts, input.NewResource)
	}

	nodeOptID = input.GetSelect("Node:", "", nodesOpts, survey.Required)

	if nodeOptID == input.NewResource {
		return nil
	}

	vars.NodeID = nodes[nodeOptID].NodeID

	return nodes[nodeOptID]
}

func nodesByTenant() map[string]*topology.Node {
	t := tenant.GetTenant()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	lr := &topology.ListNodesRequest{
		Meta: &resource.ListRequest{},
		Tenant: &tenant_pb.TenantReq{
			AccountID: t.AccountID,
			TenantID:  t.TenantID,
		},
	}

	nodes := make(map[string]*topology.Node) // map[nodeName]*topology.Node

	for {
		nl, err := nxc.ListNodes(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list nodes by tenant")
		}

		for _, n := range nl.Nodes {
			if n.Cfg != nil {
				if len(n.Cfg.NodeName) > 0 {
					nodes[n.Cfg.NodeName] = n
				}
			}
		}

		if len(nl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = nl.Meta.NextPageToken
		} else {
			break
		}
	}

	s.Stop()

	return nodes
}

func nodesBySubnet() map[string]*topology.Node {
	s := subnet.GetSubnet(false)

	ss := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	lr := &topology.ListNodesRequest{
		Meta: &resource.ListRequest{},
		Tenant: &tenant_pb.TenantReq{
			AccountID: s.AccountID,
			TenantID:  s.TenantID,
		},
		Filter: &topology.NodeFilter{
			NetID:    s.NetID,
			SubnetID: s.SubnetID,
		},
	}

	nodes := make(map[string]*topology.Node) // map[nodeName]*topology.Node

	for {
		nl, err := nxc.ListNodes(context.TODO(), lr)
		if err != nil {
			ss.Stop()
			status.Error(err, "Unable to list nodes by subnet")
		}

		for _, n := range nl.Nodes {
			if n.Cfg != nil {
				if len(n.Cfg.NodeName) > 0 {
					nodes[n.Cfg.NodeName] = n
				}
			}
		}

		if len(nl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = nl.Meta.NextPageToken
		} else {
			break
		}
	}

	ss.Stop()

	return nodes
}

func FetchNode(nr *topology.NodeReq) *topology.Node {
	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	n, err := nxc.GetNode(context.TODO(), nr)
	if err != nil {
		status.Error(err, "Unable to get node")
	}

	return n
}
