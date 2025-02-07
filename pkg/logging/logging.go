package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
	once   sync.Once
)

// Default settings
const (
	defaultLogLevel  = "info"
	defaultLogFormat = "text"
)

// SetupLogging initializes the logger with the appropriate settings.
func SetupLogging() *logrus.Logger {
	once.Do(func() {
		logger = logrus.New()
		logger.SetOutput(os.Stdout)
		logger.SetReportCaller(true)

		// Get log level from environment
		logLevelStr := os.Getenv("LOG_LEVEL")
		if logLevelStr == "" {
			logLevelStr = defaultLogLevel
		}
		logLevel := getLogLevel(logLevelStr)
		logger.SetLevel(logLevel)

		// Get log format from environment
		logFormat := strings.ToLower(os.Getenv("LOG_FORMAT"))
		if logFormat == "" {
			logFormat = defaultLogFormat
		}

		switch logFormat {
		case "json":
			logger.SetFormatter(&logrus.JSONFormatter{
				DisableTimestamp: false,
				PrettyPrint:      false,
			})
		default:
			logger.SetFormatter(&CustomTextFormatter{
				DisableTimestamp: false,
			})
		}

		logger.Debugf("Logger initialized with level: %s and format: %s", logLevel.String(), logFormat)
	})

	return logger
}

// CustomTextFormatter formats log entries in a more concise way.
type CustomTextFormatter struct {
	DisableTimestamp bool
}

func (f *CustomTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	caller := ""
	function := ""
	if entry.HasCaller() {
		file := filepath.Base(entry.Caller.File)
		caller = fmt.Sprintf("%s:%d", file, entry.Caller.Line)
		function = entry.Caller.Function
	}

	timestamp := ""
	if !f.DisableTimestamp {
		timestamp = entry.Time.Format("2006-01-02 15:04:05") + " "
	}

	// More readable log formatting
	logMessage := fmt.Sprintf("%slevel=%s msg=\"%s\"", timestamp, entry.Level.String(), entry.Message)
	if function != "" {
		logMessage += fmt.Sprintf(" func=%s", function)
	}
	if caller != "" {
		logMessage += fmt.Sprintf(" caller=%s", caller)
	}
	logMessage += "\n"

	return []byte(logMessage), nil
}

// getLogLevel returns the logrus log level based on the input string.
func getLogLevel(level string) logrus.Level {
	switch strings.ToLower(level) {
	case "debug":
		return logrus.DebugLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}
