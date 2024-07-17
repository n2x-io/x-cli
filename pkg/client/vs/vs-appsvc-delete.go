package vs

import (
	"context"
	"fmt"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"n2x.dev/x-api-go/grpc/resources/topology"
	"n2x.dev/x-cli/pkg/grpc"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-cli/pkg/status"
)

func (api *API) DeleteVSAppSvc() {
	vs := getVS(false)

	dvsasr := &topology.DeleteVSAppSvcRequest{
		AccountID: vs.AccountID,
		VSID:      vs.VSID,
	}

	var appSvcOptID string
	appSvcsOpts := make([]string, 0)
	appSvcs := make(map[string]*topology.AppSvc)

	for _, as := range vs.AppSvcs {
		appSvcOptID = fmt.Sprintf("[%s] port %d", as.NodeName, as.RSPort)
		appSvcsOpts = append(appSvcsOpts, appSvcOptID)
		appSvcs[appSvcOptID] = as
	}

	sort.Strings(appSvcsOpts)

	appSvcOptID = input.GetSelect("Remove node appSvc?", "", appSvcsOpts, survey.Required)

	dvsasr.AppSvcID = appSvcs[appSvcOptID].AppSvcID

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	vs, err := nxc.DeleteVSAppSvc(context.TODO(), dvsasr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to modifiy virtual server")
	}

	s.Stop()

	// output.Show(n)
	Output().Show(vs)
}
