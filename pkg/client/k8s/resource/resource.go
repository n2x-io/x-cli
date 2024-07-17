package resource

const (
	KubernetesResourceTypeService int = iota
	KubernetesResourceTypeStatefulSet
	KubernetesResourceTypeDeployment
	KubernetesResourceTypeDaemonSet
)

type KubernetesResource struct {
	KubernetesResourceType int
	Namespace              string
	Name                   string
	Connected              bool

	NetStatus NetStatus

	Labels             KubernetesLabels
	ServiceAnnotations map[string]string
}

type NetStatus struct {
	TenantID string
	NetID    string
	SubnetID string
}

type KubernetesLabels struct {
	NodeType  string
	AccountID string
	TenantID  string
	// NodeID      string
	NodeGroupID string
	NetID       string
	SubnetID    string
}

func (r *KubernetesResource) ParseLabelsPod(labels map[string]string) {
	if labels == nil {
		return
	}

	for k, v := range labels {
		switch k {
		case "n2x-type":
			r.Labels.NodeType = v
		case "n2x-account":
			r.Labels.AccountID = v
		case "n2x-tenant":
			r.Labels.TenantID = v
			r.NetStatus.TenantID = v
		case "n2x-nodegroup":
			r.Labels.NodeGroupID = v
		case "n2x-network":
			// r.Labels.NetID = v
			r.NetStatus.NetID = v
		case "n2x-subnet":
			// r.Labels.SubnetID = v
			r.NetStatus.SubnetID = v
		}
	}

	if len(r.Labels.NodeType) > 0 &&
		len(r.Labels.AccountID) > 0 &&
		len(r.Labels.TenantID) > 0 &&
		len(r.Labels.NodeGroupID) > 0 {
		r.Connected = true
	}
}

func (r *KubernetesResource) ParseLabelsGateway(labels map[string]string) {
	if labels == nil {
		return
	}

	for k, v := range labels {
		switch k {
		case "n2x-type":
			r.Labels.NodeType = v
		case "n2x-account":
			r.Labels.AccountID = v
		case "n2x-tenant":
			r.Labels.TenantID = v
			r.NetStatus.TenantID = v
		// case "n2x-node":
		// 	r.Labels.NodeID = v
		case "n2x-nodegroup":
			r.Labels.NodeGroupID = v
		case "n2x-network":
			r.Labels.NetID = v
			r.NetStatus.NetID = v
		case "n2x-subnet":
			r.Labels.SubnetID = v
			r.NetStatus.SubnetID = v
		}
	}

	if len(r.Labels.NodeType) > 0 &&
		len(r.Labels.AccountID) > 0 &&
		len(r.Labels.TenantID) > 0 &&
		len(r.Labels.NodeGroupID) > 0 &&
		len(r.Labels.NetID) > 0 &&
		len(r.Labels.SubnetID) > 0 {
		r.Connected = true
	}
}

func (r *KubernetesResource) ParseServiceAnnotations(annotations map[string]string) {
	if annotations == nil {
		return
	}

	var hasTenant, hasNetwork, hasSubnet bool

	for k, v := range annotations {
		switch k {
		case "n2x.io/account":
			r.ServiceAnnotations[k] = v
		case "n2x.io/tenant":
			r.ServiceAnnotations[k] = v
			r.NetStatus.TenantID = v
			hasTenant = true
		case "n2x.io/network":
			r.ServiceAnnotations[k] = v
			r.NetStatus.NetID = v
			hasNetwork = true
		case "n2x.io/subnet":
			r.ServiceAnnotations[k] = v
			r.NetStatus.SubnetID = v
			hasSubnet = true
		case "n2x.io/dnsName":
			r.ServiceAnnotations[k] = v
		case "n2x.io/ipv4":
			r.ServiceAnnotations[k] = v
		}
	}

	if hasTenant && hasNetwork && hasSubnet {
		r.Connected = true
	}
}
