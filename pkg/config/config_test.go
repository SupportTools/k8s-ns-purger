package config

import (
	"os"
	"testing"
	"time"
)

func TestLoadConfiguration(t *testing.T) {
	os.Setenv("DEBUG", "true")
	os.Setenv("KUBECONFIG", "/path/to/kubeconfig")
	os.Setenv("LABEL_SELECTOR", "test-selector")
	os.Setenv("PURGE_AGE", "6h")
	os.Setenv("INTERVAL", "30m")

	LoadConfiguration()

	if !CFG.Debug {
		t.Errorf("Expected Debug to be true, got %v", CFG.Debug)
	}
	if CFG.KubeConfig != "/path/to/kubeconfig" {
		t.Errorf("Expected KubeConfig to be '/path/to/kubeconfig', got %v", CFG.KubeConfig)
	}
	if CFG.LabelSelector != "test-selector" {
		t.Errorf("Expected LabelSelector to be 'test-selector', got %v", CFG.LabelSelector)
	}
	if CFG.PurgeAge != 6*time.Hour {
		t.Errorf("Expected PurgeAge to be 6h, got %v", CFG.PurgeAge)
	}
	if CFG.Interval != 30*time.Minute {
		t.Errorf("Expected Interval to be 30m, got %v", CFG.Interval)
	}
}

func TestGetEnvOrDefault(t *testing.T) {
	os.Setenv("TEST_ENV", "value")
	defer os.Unsetenv("TEST_ENV")

	value := getEnvOrDefault("TEST_ENV", "default")
	if value != "value" {
		t.Errorf("Expected 'value', got %v", value)
	}

	value = getEnvOrDefault("NON_EXISTENT_ENV", "default")
	if value != "default" {
		t.Errorf("Expected 'default', got %v", value)
	}
}

func TestParseEnvBool(t *testing.T) {
	os.Setenv("BOOL_ENV", "true")
	defer os.Unsetenv("BOOL_ENV")

	value := parseEnvBool("BOOL_ENV", false)
	if !value {
		t.Errorf("Expected true, got %v", value)
	}

	value = parseEnvBool("NON_EXISTENT_BOOL_ENV", true)
	if !value {
		t.Errorf("Expected true, got %v", value)
	}
}

func TestParseEnvDuration(t *testing.T) {
	os.Setenv("DURATION_ENV", "5h")
	defer os.Unsetenv("DURATION_ENV")

	value := parseEnvDuration("DURATION_ENV", "2h")
	if value != 5*time.Hour {
		t.Errorf("Expected 5h, got %v", value)
	}

	value = parseEnvDuration("NON_EXISTENT_DURATION_ENV", "2h")
	if value != 2*time.Hour {
		t.Errorf("Expected 2h, got %v", value)
	}
}
