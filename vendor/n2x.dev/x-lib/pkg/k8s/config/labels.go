package config

import "n2x.dev/x-api-go/grpc/resources/topology"

type AppLabel string

const (
	AppLabelController AppLabel = "n2x-controller"
	AppLabelNode       AppLabel = "n2x-node"
)

func (a AppLabel) String() string {
	return string(a)
}

func NodeLabels(ni *topology.NodeInstance) map[string]string {
	return map[string]string{
		"n2x-app":       AppLabelNode.String(),
		"n2x-type":      ni.Node.Type.String(),
		"n2x-account":   ni.Node.AccountID,
		"n2x-tenant":    ni.Node.TenantID,
		"n2x-nodegroup": ni.Node.NodeGroupID,
		"n2x-network":   ni.Node.Cfg.NetID,
		"n2x-subnet":    ni.Node.Cfg.SubnetID,
		// "n2x-node":      ni.Node.NodeID,
	}
}
