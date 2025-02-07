package logging

import (
	"bytes"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetupLogging(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("LOG_FORMAT", "json")

	// Capture the output
	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	// Initialize the logger
	logger := SetupLogging()

	// Check if the logger is not nil
	assert.NotNil(t, logger)

	// Check if the logger level is set correctly
	assert.Equal(t, logrus.DebugLevel, logger.GetLevel())

	// Check if the logger formatter is set correctly
	_, ok := logger.Formatter.(*logrus.JSONFormatter)
	assert.True(t, ok, "Expected JSONFormatter")

	// Reset environment variables
	os.Unsetenv("LOG_LEVEL")
	os.Unsetenv("LOG_FORMAT")
}

func TestCustomTextFormatter(t *testing.T) {
	formatter := &CustomTextFormatter{DisableTimestamp: true}
	entry := &logrus.Entry{
		Logger:  logrus.New(),
		Level:   logrus.InfoLevel,
		Message: "Test message",
	}

	// Format the entry
	formatted, err := formatter.Format(entry)
	assert.NoError(t, err)

	// Check the formatted output
	expected := "level=info msg=\"Test message\"\n"
	assert.Equal(t, expected, string(formatted))
}

func TestGetLogLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected logrus.Level
	}{
		{"debug", logrus.DebugLevel},
		{"warn", logrus.WarnLevel},
		{"warning", logrus.WarnLevel},
		{"error", logrus.ErrorLevel},
		{"fatal", logrus.FatalLevel},
		{"panic", logrus.PanicLevel},
		{"invalid", logrus.InfoLevel}, // Default case
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, getLogLevel(test.input))
	}
}
