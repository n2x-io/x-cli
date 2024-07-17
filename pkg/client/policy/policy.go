package policy

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/client/node"
	"n2x.dev/x-cli/pkg/client/subnet"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-cli/pkg/vars"
	"n2x.dev/x-lib/pkg/ipnet"
	"n2x.dev/x-lib/pkg/utils/msg"
)

type nfObjectType string

const (
	endpointObject  nfObjectType = "Endpoint Object"
	ipNetCIDRObject nfObjectType = "IPNet CIDR"
)

type nfEndpoint struct {
	objectType nfObjectType
	af         ipnet.AddressFamily
	addr       string
}

func getNetworkFilter(np *topology.Policy) *topology.Filter {
	var objects []string

	nfilters := make(map[string]int)

	for i, nf := range np.NetworkFilters {
		obj := fmt.Sprintf("%d: src %s | dst %s | %d/%s | %s",
			nf.Index, nf.SrcIPNet, nf.DstIPNet, nf.DstPort, nf.Proto, nf.Policy)
		objects = append(objects, obj)
		nfilters[obj] = i
	}

	sort.Strings(objects)

	if len(objects) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	objID := input.GetSelect("Network Filter", "", objects, survey.Required)

	i := nfilters[objID]

	return np.NetworkFilters[i]
}

func setNetworkFilter(nf *topology.Filter) {
	nf.Index = getNetworkFilterIndex(nf.Index)

	nf.Description = input.GetInput("Description (optional):", "", nf.Description, nil)

	af := getNetworkFilterAddressFamily()

	src := getNetworkFilterEndpoint(false, nf.SrcIPNet, af)
	nf.SrcIPNet = src.addr
	dst := getNetworkFilterEndpoint(true, nf.DstIPNet, af)
	nf.DstIPNet = dst.addr

	nf.Proto = getProto(af)

	if nf.Proto == topology.Protocol_TCP || nf.Proto == topology.Protocol_UDP {
		nf.DstPort = getDstPort(nf.DstPort)
	} else {
		nf.DstPort = 0
	}

	nf.Policy = subnet.GetSecurityPolicy("Security Policy:")
}

func getNetworkFilterIndex(defIdx uint32) uint32 {
	index := input.GetInput("Policy Index:", "", fmt.Sprintf("%d", defIdx), input.ValidUint)

	i, err := strconv.Atoi(index)
	if err != nil {
		status.Error(err, "Invalid index")
	}

	return uint32(i)
}

func getNetworkFilterAddressFamily() ipnet.AddressFamily {
	opts := []string{
		ipnet.AddressFamilyIPv4.String(),
		ipnet.AddressFamilyIPv6.String(),
	}

	afOpt := input.GetSelect("Protocol:", "", opts, survey.Required)

	switch afOpt {
	case ipnet.AddressFamilyIPv4.String():
		return ipnet.AddressFamilyIPv4
	case ipnet.AddressFamilyIPv6.String():
		return ipnet.AddressFamilyIPv6
	}

	return ipnet.AddressFamilyUnspec
}

func getNetworkFilterEndpoint(dst bool, def string, af ipnet.AddressFamily) *nfEndpoint {
	var nfEndp *nfEndpoint

	vars.NodeID = ""

	endpMsg := "Source"
	if dst {
		endpMsg = "Destination"
	}

	var endpObj string
	endpOpts := []string{string(endpointObject), string(ipNetCIDRObject)}

	inputText := fmt.Sprintf("Select %s:", endpMsg)

	endpObj = input.GetSelect(inputText, "", endpOpts, survey.Required)

	if endpObj == string(endpointObject) {
		n := node.GetNodeBySubnet(false)
		e := node.GetEndpoint(n)
		nfEndp = &nfEndpoint{
			objectType: endpointObject,
			af:         af,
		}
		switch af {
		case ipnet.AddressFamilyIPv4:
			nfEndp.addr = fmt.Sprintf("%s/32", e.IPv4)
		case ipnet.AddressFamilyIPv6:
			nfEndp.addr = fmt.Sprintf("%s/128", e.IPv6)
		}
	} else { // endpObj == ipNetCIDRObject
		var ok bool
		var addr string

		inputMsg := fmt.Sprintf("%s %s CIDR:", endpMsg, af.String())

		for !ok {
			addr = input.GetInput(inputMsg, ipNetHelp(), def, input.ValidIPNetCIDR)

			switch af {
			case ipnet.AddressFamilyIPv4:
				if strings.Contains(addr, ":") {
					msg.Error("Invalid IPv4 address")
					continue
				}
			case ipnet.AddressFamilyIPv6:
				if strings.Contains(addr, ".") {
					msg.Error("Invalid IPv6 address")
					continue
				}
			}
			ok = true
		}

		nfEndp = &nfEndpoint{
			objectType: ipNetCIDRObject,
			addr:       addr,
			af:         af,
		}
	}

	return nfEndp
}

func getProto(af ipnet.AddressFamily) topology.Protocol {
	var proto topology.Protocol

	var protocols []string
	switch af {
	case ipnet.AddressFamilyIPv4:
		protocols = []string{
			topology.Protocol_TCP.String(),
			topology.Protocol_UDP.String(),
			topology.Protocol_ICMPv4.String(),
			topology.Protocol_ANY.String(),
		}
	case ipnet.AddressFamilyIPv6:
		protocols = []string{
			topology.Protocol_TCP.String(),
			topology.Protocol_UDP.String(),
			topology.Protocol_ICMPv6.String(),
			topology.Protocol_ANY.String(),
		}
	}

	p := input.GetSelect("Protocol:", "", protocols, survey.Required)

	switch p {
	case topology.Protocol_ANY.String():
		proto = topology.Protocol_ANY
	case topology.Protocol_TCP.String():
		proto = topology.Protocol_TCP
	case topology.Protocol_UDP.String():
		proto = topology.Protocol_UDP
	case topology.Protocol_ICMPv4.String():
		proto = topology.Protocol_ICMPv4
	case topology.Protocol_ICMPv6.String():
		proto = topology.Protocol_ICMPv6
	default:
		proto = topology.Protocol_ANY
	}

	return proto
}

func getDstPort(dstPort uint32) uint32 {
	helpText := "Destination TCP or UDP port"
	defaultPort := fmt.Sprintf("%d", dstPort)

	dstP := input.GetInput("Destination Port:", helpText, defaultPort, input.ValidPort)

	p, err := strconv.Atoi(dstP)
	if err != nil {
		status.Error(err, "Invalid port")
	}
	return uint32(p)
}

func ipNetHelp() string {
	text := `CIDR notation and prefix length, like '192.168.1.14/32', '10.32.4.0/24',
'fd77:f:a45:2d2:1::/128' or '2001:db8::/32' as defined in RFC 4632 and RFC 4291`

	return text
}
