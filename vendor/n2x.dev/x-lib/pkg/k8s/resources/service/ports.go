package service

import (
	corev1 "k8s.io/api/core/v1"
	"n2x.dev/x-lib/pkg/errors"
)

const (
	nodePortMin int = 30000
	nodePortMax int = 32767
)

func (a *API) GetNodePort() (int, error) {
	svcList, err := a.List("")
	if err != nil {
		return 0, errors.Wrapf(err, "[%v] function a.List()", errors.Trace())
	}

	var portAvailable bool
	var nodePort int
	for p := nodePortMin; p <= nodePortMax; p++ {
		portAvailable = true
		for _, svc := range svcList.Items {
			if svc.Spec.Type != corev1.ServiceTypeNodePort {
				continue
			}
			for _, port := range svc.Spec.Ports {
				if port.NodePort == int32(p) {
					portAvailable = false
					break
				}
			}
			if !portAvailable {
				break
			}
		}
		if portAvailable {
			nodePort = p
			break
		}
	}

	if !portAvailable || nodePort == 0 {
		return 0, errors.New("no external port available")
	}

	return nodePort, nil
}

/*
func getExternalPort(externalIPv4 string) (int, error) {
	clientset, err = config.NewClient(a.KubeConfig)
	if err != nil {
		return 0, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
	}

	listOpts := metav1.ListOptions{
		//FieldSelector: fmt.Sprintf("spec.loadBalancerIP=%s", externalIPv4),
	}

	svcList, err := clientset.CoreV1().Services("").List(context.TODO(), listOpts)
	if err != nil {
		return 0, errors.Wrapf(err, "[%v] function clientset.CoreV1().Services().List()", errors.Trace())
	}

	fmt.Printf("There are %d services in the cluster using the externalIP %s\n", len(svcList.Items), externalIPv4)

	var portAvailable bool
	var externalPort int
	for p := sxExternalPortMin; p < sxExternalPortMax; p++ {
		portAvailable = true
		for _, svc := range svcList.Items {
			if svc.Spec.LoadBalancerIP != externalIPv4 {
				continue
			}
			for _, port := range svc.Spec.Ports {
				if port.Port == int32(p) {
					portAvailable = false
					break
				}
			}
			if !portAvailable {
				break
			}
		}
		if portAvailable {
			externalPort = p
			break
		}
	}

	if !portAvailable || externalPort == 0 {
		return 0, errors.New("no external port available")
	}

	return externalPort, nil
}
*/
