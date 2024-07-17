package client

import (
	"n2x.dev/x-cli/pkg/client/account"
	"n2x.dev/x-cli/pkg/client/alert"
	"n2x.dev/x-cli/pkg/client/auth"
	"n2x.dev/x-cli/pkg/client/vs"

	// "n2x.dev/x-cli/pkg/client/command"
	"n2x.dev/x-cli/pkg/client/iam/acl"
	// "n2x.dev/x-cli/pkg/client/iam/role"
	// "n2x.dev/x-cli/pkg/client/iam/sg"
	"n2x.dev/x-cli/pkg/client/iam/user"
	"n2x.dev/x-cli/pkg/client/k8s"
	"n2x.dev/x-cli/pkg/client/network"
	"n2x.dev/x-cli/pkg/client/node"
	"n2x.dev/x-cli/pkg/client/ops/project"
	"n2x.dev/x-cli/pkg/client/ops/workflow"
	"n2x.dev/x-cli/pkg/client/ops/workflow/tasklog"
	"n2x.dev/x-cli/pkg/client/policy"
	"n2x.dev/x-cli/pkg/client/subnet"
	"n2x.dev/x-cli/pkg/client/tenant"
)

func Auth() auth.Interface {
	return &auth.API{}
}

func Account() account.Interface {
	return &account.API{}
}

func Tenant() tenant.Interface {
	return &tenant.API{}
}

func Network() network.Interface {
	return &network.API{}
}

func Subnet() subnet.Interface {
	return &subnet.API{}
}

func Node() node.Interface {
	return &node.API{}
}

func VS() vs.Interface {
	return &vs.API{}
}

func NetworkPolicy() policy.Interface {
	return &policy.API{}
}

func ACL() acl.Interface {
	return &acl.API{}
}

/*
func Role() role.Interface {
	return &role.API{}
}

func SecurityGroup() sg.Interface {
	return &sg.API{}
}
*/

func User() user.Interface {
	return &user.API{}
}

func Project() project.Interface {
	return &project.API{}
}

func Workflow() workflow.Interface {
	return &workflow.API{}
}

func TaskLog() tasklog.Interface {
	return &tasklog.API{}
}

func Alert() alert.Interface {
	return &alert.API{}
}

func Kubernetes() k8s.Interface {
	return &k8s.API{}
}

// func Command() command.Interface {
// 	return &command.API{}
// }
