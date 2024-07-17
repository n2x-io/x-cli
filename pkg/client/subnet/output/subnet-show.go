package output

import (
	"fmt"
	"strconv"

	"n2x.dev/x-api-go/grpc/resources/topology"
	np_output "n2x.dev/x-cli/pkg/client/policy/output"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/output/table"
	"n2x.dev/x-lib/pkg/ipnet"
	"n2x.dev/x-lib/pkg/utils/colors"
)

func (api *API) Show(s *topology.Subnet) {
	output.SectionHeader("Subnet Details")
	output.TitleT1("Subnet Information")

	t := table.New()

	// t.AddRow(colors.Black("Account ID"), colors.DarkWhite(s.AccountID))
	t.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(s.TenantID))
	t.AddRow(colors.Black("Network ID"), colors.DarkWhite(s.NetID))
	t.AddRow(colors.Black("Subnet ID"), colors.DarkWhite(s.SubnetID))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(s.Description))

	t.Render()
	fmt.Println()

	if s.IPAM != nil {
		output.SubTitleT2("IP Address Management (IPAM)")

		t := table.New()

		t.AddRow(colors.Black("Network CIDR"), colors.DarkWhite(s.IPAM.NetworkCIDR))
		t.AddRow(colors.Black("Subnet CIDR"), colors.DarkWhite(s.IPAM.SubnetCIDR))
		t.AddRow(colors.Black("IPv4 Addresses Available"), colors.DarkWhite(strconv.Itoa(int(s.IPAM.TotalAvailable))))
		t.AddRow(colors.Black("IPv4 Addresses Leased"), colors.DarkWhite(strconv.Itoa(int(s.IPAM.TotalLeased))))

		t.Render()
		fmt.Println()

		if s.IPAM.TotalLeased > 0 {
			t := table.New()
			// t.Header(colors.Black("IPv4 ADDRESS"), colors.Black("IPv6 ADDRESS"), colors.Black("ENDPOINT"))
			t.Header(colors.Black("IPv4 Address"), colors.Black("IPv6 Address"), colors.Black("Endpoint"))
			// t.Header(output.TableHeader("IPv4 Address"), output.TableHeader("IPv6 Address"), output.TableHeader("Endpoint"))

			for ipv4, leaseEndpointMap := range s.IPAM.Leased {
				ipv6, err := ipnet.GetIPv6(ipv4)
				if err != nil {
					ipv6 = "-"
				}

				l := 0
				for endpoint := range leaseEndpointMap.Endpoints {
					if l == 0 {
						t.AddRow(colors.DarkWhite(ipv4), ipv6, colors.DarkWhite(output.Fit(endpoint, 32)))
					} else {
						t.AddRow("", "", colors.DarkWhite(endpoint))
					}
					l++
				}
			}

			t.Render()
			fmt.Println()
		}
	}

	if s.NetworkPolicy != nil {
		output.SubTitleT2("Security Policy")
		np_output.ShowNetworkPolicy(s.NetworkPolicy)
	}
}
