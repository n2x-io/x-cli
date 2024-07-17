package k8s

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/k8s/config"
)

func getKubeConfig() ([]byte, error) {
	kubeConfigFile := os.Getenv("KUBECONFIG")

	if len(kubeConfigFile) == 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function os.UserHomeDir()", errors.Trace())
		}

		kubeConfigFile = filepath.Join(homeDir, ".kube", "config")
	}

	if !fileExists(kubeConfigFile) {
		return nil, errors.New("missing kubeconfig file")
	}

	kubeConfig, err := ioutil.ReadFile(kubeConfigFile)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function ioutil.ReadFile()", errors.Trace())
	}

	if len(kubeConfig) == 0 {
		return nil, errors.New("invalid kubeconfig file")
	}

	if _, err := config.NewClient(kubeConfig); err != nil {
		return nil, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
	}

	return kubeConfig, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
