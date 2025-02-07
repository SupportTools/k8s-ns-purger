package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/supporttools/k8s-ns-purger/pkg/config"
	"github.com/supporttools/k8s-ns-purger/pkg/logging"
	"github.com/supporttools/k8s-ns-purger/pkg/version"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	logger    = logging.SetupLogging()
	clientset *kubernetes.Clientset
)

func main() {
	logger.Infof("Starting k8s-ns-purger - %s", version.GetVersionString())

	// Load and validate configuration.
	logger.Info("Loading configuration...")
	config.LoadConfiguration()

	// Connect to Kubernetes.
	logger.Info("Connecting to Kubernetes...")
	kubeConfig, err := rest.InClusterConfig()
	if err != nil {
		logger.Fatalf("Failed to create Kubernetes config: %v", err)
	}

	clientset, err = kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		logger.Fatalf("Failed to connect to Kubernetes: %v", err)
	}
	logger.Info("Connected to Kubernetes successfully.")

	// Run main loop
	go func() {
		for {
			logger.Info("Deleting namespaces...")
			deleteNamespaces()
			logger.Infof("Sleeping for %v...", config.CFG.Interval)
			time.Sleep(config.CFG.Interval)
		}
	}()

	// Handle OS signals for graceful shutdown.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	logger.Info("Shutting down server...")
}

func deleteNamespaces() {
	nsList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{
		LabelSelector: config.CFG.LabelSelector,
	})
	if err != nil {
		logger.Errorf("Error fetching namespaces: %v", err)
		return
	}

	deleteTime := time.Now().UTC().Add(-config.CFG.PurgeAge)

	for _, ns := range nsList.Items {
		logger.Infof("Checking namespace: %s, created at %s", ns.Name, ns.CreationTimestamp.Time)
		if ns.CreationTimestamp.Time.Before(deleteTime) {
			logger.Infof("Deleting namespace: %s", ns.Name)
			if err := clientset.CoreV1().Namespaces().Delete(context.TODO(), ns.Name, metav1.DeleteOptions{}); err != nil {
				logger.Errorf("Error deleting namespace %s: %v", ns.Name, err)
			} else {
				logger.Infof("Successfully deleted namespace: %s", ns.Name)
			}
		}
	}
}
