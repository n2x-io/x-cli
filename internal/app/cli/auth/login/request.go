package login

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"n2x.dev/x-api-go/grpc/resources/iam/auth"
	"n2x.dev/x-cli/pkg/input"
	"n2x.dev/x-cli/pkg/output"
	"n2x.dev/x-lib/pkg/utils/msg"
	"n2x.dev/x-lib/pkg/version"
)

func userAgent() string {
	return fmt.Sprintf("%s/%s", version.CLI_NAME, version.GetVersion())
}

func NewRequestWithOTP() *auth.OTPSigninRequest {
	req := &auth.OTPSigninRequest{
		UserEmail: viper.GetString("user.email"),
		Method:    auth.SigninMethod_SIGNIN_BY_EMAIL,
		UserAgent: userAgent(),
	}

	if len(req.UserEmail) == 0 {
		output.AuthenticationRequired()
		req.UserEmail = input.GetInput("Email:", "", "", input.ValidEmail)
	}

	return req
}

func NewRequestWithToken() *auth.LoginRequest {
	userToken := viper.GetString("token")

	if len(userToken) == 0 {
		msg.Error("Authorization token not found")
		os.Exit(1)
	}

	return &auth.LoginRequest{
		UserToken: userToken,
	}
}
