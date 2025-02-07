// pkg/version/version.go
package version

import "fmt"

// Variables set during build time
var (
	Version   = "unknown" // Set via -ldflags during build
	GitCommit = "unknown" // Set via -ldflags during build
	BuildTime = "unknown" // Set via -ldflags during build
)

// Info holds version details.
type Info struct {
	Version   string `json:"version"`
	GitCommit string `json:"git_commit"`
	BuildTime string `json:"build_time"`
}

// GetVersionInfo returns version details as a struct.
func GetVersionInfo() Info {
	return Info{
		Version:   Version,
		GitCommit: GitCommit,
		BuildTime: BuildTime,
	}
}

// GetVersionString returns a formatted string with version information.
func GetVersionString() string {
	return fmt.Sprintf("Version: %s, GitCommit: %s, BuildTime: %s", Version, GitCommit, BuildTime)
}
