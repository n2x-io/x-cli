package policy

/*
import (
	"fmt"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/subnet"
)

func (api *API) Export() {
	s := subnet.GetSubnet(false)

	fmt.Println()
	fmt.Println(networkPolicyExport(s))
	fmt.Println()
}

func networkPolicyExport(s *topology.Subnet) string {
	header := `# n2x networkPolicy
apiVersion: v1
kind: NetworkPolicy

tenant: ` + s.TenantID + `
network: ` + s.NetID + `
subnet: ` + s.SubnetID + `

enabled: true

networkPolicy:
  defaultPolicy: ` + s.NetworkPolicy.DefaultPolicy.String() + `
  networkFilters:`

	var nfilters string
	for _, nf := range s.NetworkPolicy.NetworkFilters {
		f := `
  - index: ` + fmt.Sprintf("%d", nf.Index) + `
    description: ` + nf.Description + `
    srcIPNet: ` + nf.SrcIPNet + `
    dstIPNet: ` + nf.DstIPNet + `
    proto: ` + nf.Proto.String() + `
    dstPort: ` + fmt.Sprintf("%d", nf.DstPort) + `
    policy: ` + nf.Policy.String()

		nfilters += f
	}

	return header + nfilters
}
*/
