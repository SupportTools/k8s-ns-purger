package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetVersionInfo(t *testing.T) {
	// Set test values
	Version = "1.0.0"
	GitCommit = "abc123"
	BuildTime = "2025-01-27T12:00:00Z"

	// Expected output
	expected := Info{
		Version:   "1.0.0",
		GitCommit: "abc123",
		BuildTime: "2025-01-27T12:00:00Z",
	}

	// Check if GetVersionInfo returns the expected struct
	assert.Equal(t, expected, GetVersionInfo())
}

func TestGetVersionString(t *testing.T) {
	// Set test values
	Version = "1.0.0"
	GitCommit = "abc123"
	BuildTime = "2025-01-27T12:00:00Z"

	// Expected output
	expected := "Version: 1.0.0, GitCommit: abc123, BuildTime: 2025-01-27T12:00:00Z"

	// Check if GetVersionString returns the expected string
	assert.Equal(t, expected, GetVersionString())
}