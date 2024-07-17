package event

import "n2x.dev/x-cli/pkg/client/event/output"

func Output() output.Interface {
	return &output.API{}
}
