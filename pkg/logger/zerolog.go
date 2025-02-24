package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log zerolog.Logger

// InitLogger init global logger
func InitLogger(serviceName string) {
	var multiWriter io.Writer
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	// Config log rotation
	saveLogsInFiles := os.Getenv("LOG_TO_FILE")
	if saveLogsInFiles == "true" {
		logFile := &lumberjack.Logger{
			Filename:   "logs/app.log",
			MaxSize:    10,   // MB
			MaxBackups: 3,    // Matains 3 older files
			MaxAge:     30,   // Days before to delete old logs
			Compress:   true, // Compress old logs
		}
		multiWriter = io.MultiWriter(consoleWriter, logFile)
	} else {
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
