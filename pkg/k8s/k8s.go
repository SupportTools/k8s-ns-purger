// pkg/k8s/k8s.go
package k8s

import (
	"fmt"
	"os"

	"github.com/supporttools/k8s-ns-purger/pkg/config"
	"github.com/supporttools/k8s-ns-purger/pkg/logging"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var log = logging.SetupLogging()

// ConnectToK8s connects to a Kubernetes cluster by checking the environment and configuration settings.
func ConnectToK8s() (*kubernetes.Clientset, error) {
	var kubeConfig *rest.Config
	var err error
	var finalErr error

	// Attempt to use in-cluster configuration
	log.Debug("Attempting to connect using in-cluster configuration...")
	kubeConfig, err = rest.InClusterConfig()
	if err == nil {
		log.Info("Successfully obtained in-cluster configuration.")
		clientset, err := kubernetes.NewForConfig(kubeConfig)
		if err != nil {
			log.Errorf("Failed to create Kubernetes client from in-cluster config: %v", err)
			finalErr = err
		} else {
			log.Debug("Successfully created Kubernetes client using in-cluster configuration.")
			return clientset, nil
		}
	} else {
		log.Warnf("In-cluster configuration failed: %v", err)
		finalErr = err
	}

	// Fallback to KUBECONFIG
	kubeConfigPath := config.CFG.KubeConfig
	if kubeConfigPath == "" {
		kubeConfigPath = os.Getenv("KUBECONFIG")
	}

	if kubeConfigPath != "" {
		log.Debugf("Using KUBECONFIG: %s", kubeConfigPath)
		log.Debug("Attempting to build configuration from KUBECONFIG...")
		kubeConfig, err = clientcmd.BuildConfigFromFlags("", kubeConfigPath)
		if err == nil {
			log.Info("Successfully loaded configuration from KUBECONFIG.")
			clientset, err := kubernetes.NewForConfig(kubeConfig)
			if err != nil {
				log.Errorf("Failed to create Kubernetes client from KUBECONFIG: %v", err)
				finalErr = err
			} else {
				log.Debug("Successfully created Kubernetes client using KUBECONFIG.")
				return clientset, nil
			}
		} else {
			log.Errorf("Failed to load configuration from KUBECONFIG (%s): %v", kubeConfigPath, err)
			finalErr = err
		}
	} else {
		log.Debug("No KUBECONFIG path is provided.")
	}

	// All connection attempts failed
	log.Errorf("All attempts to configure Kubernetes client failed: %v", finalErr)
	return nil, fmt.Errorf("failed to configure Kubernetes client: %v", finalErr)
}
