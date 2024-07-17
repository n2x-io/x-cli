package node

import (
	"context"
	"fmt"
	"os"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/subnet"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/utils/colors"
	"n2x.dev/x-lib/pkg/utils/msg"
)

func (api *API) AddNode() {
	s := subnet.GetSubnet(false)
	if s == nil {
		msg.Alert("No subnet found.")
		msg.Alert("Please, configure at least one before adding nodes.")
		os.Exit(1)
	}

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	/*
			fmt.Printf(`
		This command does NOT create a new node, just generates YAML
		config that you can use as a template for your convenience.

		The configuration comes with some important features disabled
		that you may want to enable (relay, kubernetes controller..).

		It is highly recommended to review the management section and
		customize it according to your security standards or criteria.

		If you feel you need some support, don't hesitate to contact us
		at %s ;-)

		`, colors.White("n2x.io/discord"))
	*/

	nodeName := input.GetInput("Node Name:", "", "", input.ValidID)
	desc := input.GetInput("Description:", "", "", nil)

	nnr := &topology.NewNodeRequest{
		AccountID:   s.AccountID,
		TenantID:    s.TenantID,
		NetID:       s.NetID,
		SubnetID:    s.SubnetID,
		NodeName:    nodeName,
		Description: desc,
		Type:        topology.NodeType_GENERIC,
	}

	ss := output.Spinner()

	ni, err := nxc.CreateGenericNode(context.TODO(), nnr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to create node")
	}

	ss.Stop()

	Output().Show(ni.Node)

	fmt.Print(colors.Black("----- n2x-node.yml -----\n"))
	fmt.Printf("%s\n", colors.DarkWhite(ni.Config.YAML))
	fmt.Print(colors.Black("----- n2x-node.yml -----\n"))
	fmt.Println()
}

/*
func (api *API) GetInstallationWebhook() {
	n := GetNode(false)

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	nr := &topology.NodeReq{
		AccountID: n.AccountID,
		TenantID:  n.TenantID,
		NetID:     n.NetID,
		SubnetID:  n.SubnetID,
		NodeID:    n.NodeID,
	}

	s := output.Spinner()

	ni, err := nxc.CreateNodeInstallLinuxWebhook(context.TODO(), nr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to generate node config")
	}

	s.Stop()

	nodeName := n.Cfg.NodeName

	cmd1 := colors.DarkGreen(fmt.Sprintf("curl -s %s | sudo sh", ni.Config.WebhookURL))
	cmd2 := colors.DarkGreen(fmt.Sprintf("ssh <your_node> \"curl -s %s | sudo sh\"", ni.Config.WebhookURL))
	configFile := colors.DarkWhite("/etc/n2x/n2x-node.yml")

	fmt.Printf(`
The magic link generated is exclusively to connect the node %s to
the subnet %s. If you want to add more nodes, simply repeat this
process for them :-)

For security, these magic links expire in 24h and can be used once only,
but you can generate as many as you need.

Use the following commands on the machine %s to setup and
connect the node to your n2x network:

%s

or

%s

Once installed you can review (or modify) the configuration
at %s. It is recommended to review and customize
the management section according to your security standards or specific needs.

`,
		colors.DarkWhite(nodeName),
		colors.DarkWhite(n.SubnetID),
		colors.DarkWhite(nodeName),
		cmd1,
		cmd2,
		configFile,
	)

	// fmt.Printf("\n# Generated config, use it to create your n2x-node.yml file\n")
}
*/
