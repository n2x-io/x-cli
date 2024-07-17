package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
	"n2x.dev/x-lib/pkg/logging"
	"n2x.dev/x-lib/pkg/utils/msg"
	"n2x.dev/x-lib/pkg/version"
	"n2x.dev/x-lib/pkg/xlog"
)

const (
	// versionStable string = "stable"
	versionDev string = "dev"

	defaultAPIServerAuthServer    string = "https://n2x.network"
	defaultAPIServerAuthServerDev string = "https://dev.n2x.network"
	defaultAPIServerEndpoint      string = "n2x.network:443"
	defaultAPIServerEndpointDev   string = "dev.n2x.network:443"

	defaultServiceID string = "n2x"
)

func Init() {
	hostID, err := os.Hostname()
	if err != nil {
		msg.Error(err)
		os.Exit(1)
	}

	n2xID := fmt.Sprintf("__cli:%s:%d:%d", hostID, os.Getegid(), time.Now().Unix())

	viper.Set("n2x.id", n2xID)

	viper.Set("n2x.app", version.CLI_NAME)

	viper.Set("host.id", hostID)

	SetDefaults()

	logging.LogLevel = xlog.GetLogLevel(viper.GetString("loglevel"))
	if logging.LogLevel == -1 {
		logging.LogLevel = xlog.INFO
	}

	logging.Interactive = true
}

func SetDefaults() {
	var isDev bool

	if viper.GetString("version.branch") == versionDev {
		isDev = true
	}

	if os.Getenv("N2X_VERSION") == versionDev {
		viper.Set("version.branch", versionDev)
		isDev = true
	}

	if isDev {
		viper.Set("apiserver.authServer", defaultAPIServerAuthServerDev)
		viper.Set("apiserver.endpoint", defaultAPIServerEndpointDev)
	} else {
		viper.Set("apiserver.authServer", defaultAPIServerAuthServer)
		viper.Set("apiserver.endpoint", defaultAPIServerEndpoint)
	}

	viper.Set("serviceID", defaultServiceID)
}
