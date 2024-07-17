package grpc

import (
	"os"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	auth_pb "n2x.dev/x-api-go/grpc/resources/iam/auth"
	"n2x.dev/x-api-go/grpc/rpc"
	"n2x.dev/x-cli/pkg/auth"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/grpc/client"
	"n2x.dev/x-lib/pkg/utils/msg"
)

type apiClientParameters struct {
	authKey    *auth_pb.AuthKey
	authSecret string
	grpcServer string
}

func getManagerAPIClientParams(reqAuth bool) *apiClientParameters {
	p := &apiClientParameters{}

	if reqAuth {
		authKey, err := auth.GetAuthKey()
		if err != nil {
			msg.Alert("Invalid or inexistent authorization key. Login to refresh your token.")
			os.Exit(1)
		}
		p.authKey = authKey
	} else {
		p.authKey = auth.GetNoAuthKey()
	}

	// p.authSecret = viper.GetString("controller.authSecret")
	p.grpcServer = viper.GetString("apiserver.endpoint")

	return p
}
func getControllerAPIClientParams() *apiClientParameters {
	p := &apiClientParameters{}

	authKey, err := auth.GetAuthKey()
	if err != nil {
		msg.Alert("Invalid or inexistent authorization key. Login to refresh your token.")
		os.Exit(1)
	}
	p.authKey = authKey

	// p.authSecret = viper.GetString("controller.authSecret")

	controllerEndpoint, err := auth.GetControllerEndpoint()
	if err != nil {
		msg.Alert("Unable to get controller endpoint.")
		msg.Alert("Invalid or inexistent api key. Login to refresh your token.")
		os.Exit(1)
	}

	p.grpcServer = controllerEndpoint

	return p
}

// manager

func GetManagerAPIClient(reqAuth bool) (rpc.ManagerAPIClient, *grpc.ClientConn) {
	p := getManagerAPIClientParams(reqAuth)

	nxc, conn, err := client.NewManagerAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetAccountAPIClient(reqAuth bool) (rpc.AccountAPIClient, *grpc.ClientConn) {
	p := getManagerAPIClientParams(reqAuth)

	nxc, conn, err := client.NewAccountAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetServicesAPIClient(reqAuth bool) (rpc.ServicesAPIClient, *grpc.ClientConn) {
	p := getManagerAPIClientParams(reqAuth)

	nxc, conn, err := client.NewServicesAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetBillingAPIClient(reqAuth bool) (rpc.BillingAPIClient, *grpc.ClientConn) {
	p := getManagerAPIClientParams(reqAuth)

	nxc, conn, err := client.NewBillingAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

// controller

func GetControllerAPIClient() (rpc.ControllerAPIClient, *grpc.ClientConn) {
	p := getControllerAPIClientParams()

	nxc, conn, err := client.NewControllerAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetNetworkAPIClient(controllerHost string) (rpc.NetworkAPIClient, *grpc.ClientConn) {
	serverEndpoint := controllerHost

	p := getControllerAPIClientParams()

	if len(serverEndpoint) == 0 {
		serverEndpoint = p.grpcServer
	}

	nxc, conn, err := client.NewNetworkAPIClient(serverEndpoint, p.authKey, p.authSecret)
	if err != nil {
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetIAMAPIClient() (rpc.IAMAPIClient, *grpc.ClientConn) {
	p := getControllerAPIClientParams()

	nxc, conn, err := client.NewIAMAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetTenantAPIClient() (rpc.TenantAPIClient, *grpc.ClientConn) {
	p := getControllerAPIClientParams()

	nxc, conn, err := client.NewTenantAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		// msg.Errorf("Unable to connect to gRPC server. Check configuration and connectivity: %v", err)
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetTopologyAPIClient() (rpc.TopologyAPIClient, *grpc.ClientConn) {
	p := getControllerAPIClientParams()

	nxc, conn, err := client.NewTopologyAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		// msg.Errorf("Unable to connect to gRPC server. Check configuration and connectivity: %v", err)
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetNStoreAPIClient() (rpc.NStoreAPIClient, *grpc.ClientConn) {
	p := getControllerAPIClientParams()

	nxc, conn, err := client.NewNStoreAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		// msg.Errorf("Unable to connect to gRPC server. Check configuration and connectivity: %v", err)
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetMonitoringAPIClient() (rpc.MonitoringAPIClient, *grpc.ClientConn) {
	p := getControllerAPIClientParams()

	nxc, conn, err := client.NewMonitoringAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		// msg.Errorf("Unable to connect to gRPC server. Check configuration and connectivity: %v", err)
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetOpsAPIClient() (rpc.OpsAPIClient, *grpc.ClientConn) {
	p := getControllerAPIClientParams()

	nxc, conn, err := client.NewOpsAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		// msg.Errorf("Unable to connect to gRPC server. Check configuration and connectivity: %v", err)
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}
