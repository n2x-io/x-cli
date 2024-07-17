package vs

import (
	"context"

	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
	"n2x.dev/x-lib/pkg/ipnet"
)

func (api *API) Update() {
	vs := getVS(false)

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	s.Stop()

	uvsr := &topology.UpdateVSRequest{
		AccountID:   vs.AccountID,
		VSID:        vs.VSID,
		Name:        input.GetInput("Name:", "", vs.Name, input.ValidName),
		Description: input.GetInput("Description:", "", vs.Description, nil),
		LocationID:  vs.LocationID,
		Cname:       input.GetInput("Custom DNS CNAME:", "Fully Qualified Domain Name", vs.Cname, input.ValidFQDN),
		ReqAuth:     input.GetConfirm("Authentication:", vs.ReqAuth),
		Tags:        input.GetMultiInput("Tags:", "Tags separated by comma", vs.Tags, input.ValidTags),
	}

	if len(uvsr.Cname) > 0 {
		if err := ipnet.VSCNAMEIsValid(uvsr.Cname, vs.LocationID); err != nil {
			status.Error(err, "Unable to modify virtual server")
		}
	}

	s = output.Spinner()

	vs, err := nxc.UpdateVS(context.TODO(), uvsr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to update virtual server")
	}

	s.Stop()

	// output.Show(n)
	Output().Show(vs)
}
