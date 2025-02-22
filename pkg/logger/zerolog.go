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
	// Config log rotation
	logFile := &lumberjack.Logger{
		Filename:   "logs/app.log", // Folder where logs will save
		MaxSize:    10,             // MB
		MaxBackups: 3,              // Matains 3 older files
		MaxAge:     30,             // Days before to delete old logs
		Compress:   true,           // Compress old logs
	}

	// Config the console output with JSON format
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	multiWriter := io.MultiWriter(consoleWriter, logFile) // Salida a ambos

	// Config global level of logs
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log = zerolog.New(multiWriter).With().
		Timestamp().
		Str("service", serviceName).
		Str("hostname", getHostname()).
		Int("pid", os.Getpid()).
		Logger()

	// Allows config the log level
	logLevel := os.Getenv("LOG_LEVEL")
	if level, err := zerolog.ParseLevel(logLevel); err == nil {
		zerolog.SetGlobalLevel(level)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

// GetLogger return the logger configured
func GetLogger() zerolog.Logger {
	return log
}

// getHostname obtain the host name
func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}
