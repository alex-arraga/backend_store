package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/alex-arraga/backend_store/pkg/utils"
)

var log zerolog.Logger

// InitLogger init global logger
func InitLogger(serviceName string) {
	var multiWriter io.Writer
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	appEnv := utils.GetAppEnv()
	if appEnv == "prod" {
		logFile := &lumberjack.Logger{
			Filename:   "logs/app.log",
			MaxSize:    10,   // MB
			MaxBackups: 3,    // Matains 3 older files
			MaxAge:     30,   // Days before to delete old logs
			Compress:   true, // Compress old logs
		}

		// If appEnv is "prod" (production) logs will are saved in files and output will be JSON format
		multiWriter = io.MultiWriter(os.Stdout, logFile)
	} else {
		// Otherwise, it is assumed that it is a “dev” environment and the logs will not be saved to files, the output will be Stdout
		multiWriter = consoleWriter
	}

	// Config the output of logs
	log = zerolog.New(multiWriter).
		With().Timestamp().
		Str("service", serviceName).
		Str("hostname", getHostname()).
		Int("pid", os.Getpid()).
		Logger()

	// Config the log level
	logLevel := os.Getenv("LOG_LEVEL")
	if level, err := zerolog.ParseLevel(logLevel); err == nil {
		zerolog.SetGlobalLevel(level)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

// UseLogger return the logger configured
func UseLogger() *zerolog.Logger {
	return &log
}

// getHostname obtain the host name
func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}
