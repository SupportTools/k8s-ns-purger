package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// AppConfig defines the structure for application configuration loaded from environment variables.
type AppConfig struct {
	Debug         bool          `json:"debug"`         // Enable debug logging
	KubeConfig    string        `json:"kubeConfig"`    // Path to kubeconfig file
	LabelSelector string        `json:"labelSelector"` // Label selector for namespaces to delete
	PurgeAge      time.Duration `json:"purgeAge"`      // Duration after which namespaces should be purged
	Interval      time.Duration `json:"interval"`      // Interval between namespace purge checks
}

// CFG is the global configuration instance.
var CFG AppConfig

// LoadConfiguration loads the configuration from environment variables and validates it.
func LoadConfiguration() {
	// Load environment variables with defaults
	CFG.Debug = parseEnvBool("DEBUG", false)
	CFG.KubeConfig = getEnvOrDefault("KUBECONFIG", "")
	CFG.LabelSelector = getEnvOrDefault("LABEL_SELECTOR", "k8s-ns-purger=true")
	CFG.PurgeAge = parseEnvDuration("PURGE_AGE", "8h")
	CFG.Interval = parseEnvDuration("INTERVAL", "1h")

	// Log loaded configuration in debug mode
	if CFG.Debug {
		log.Printf("Configuration Loaded: %+v", CFG)
	}
}

// getEnvOrDefault retrieves the value of an environment variable or returns a default value if not set.
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Printf("Environment variable %s not set. Using default: %s", key, defaultValue)
	return defaultValue
}

// parseEnvBool parses an environment variable as a boolean, supporting common truthy/falsy values.
func parseEnvBool(key string, defaultValue bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not set. Using default: %t", key, defaultValue)
		return defaultValue
	}

	value = strings.ToLower(value)
	switch value {
	case "1", "true", "yes", "on", "enabled":
		return true
	case "0", "false", "no", "off", "disabled":
		return false
	default:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			log.Printf("Error parsing %s as bool: %v. Using default value: %t", key, err, defaultValue)
			return defaultValue
		}
		return boolValue
	}
}

// parseEnvDuration parses an environment variable as a time.Duration, returning a default value if parsing fails.
func parseEnvDuration(key, defaultValue string) time.Duration {
	value := getEnvOrDefault(key, defaultValue)
	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("Error parsing %s as duration: %v. Using default: %s", key, err, defaultValue)
		duration, _ = time.ParseDuration(defaultValue)
	}
	return duration
}
