package status

import (
	"fmt"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/utils/colors"
	"n2x.dev/x-lib/pkg/utils/msg"
)

func Error(err error, message string) {
	if err == nil {
		return
	}

	if len(message) > 0 {
		fmt.Printf("\n%s\n", colors.DarkRed(message))
	}

	s, ok := status.FromError(errors.Cause(err))
	if ok {
		switch s.Code() {
		case codes.Unauthenticated:
			msg.Errorf("NOT AUTHORIZED. Please login or check your IAM permissions.")
		case codes.Unknown:
			msg.Errorf("%s", s.Message())
		default:
			msg.Errorf("Code: %v | Status: %s", s.Code(), s.Message())
		}
	} else {
		msg.Errorf("Cause: %v", errors.Cause(err))
	}

	os.Exit(1)
}
