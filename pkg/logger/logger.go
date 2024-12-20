package logger

import (
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

// interface for logger
type LoggerI interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

// logger struct
type Logger struct {
	logger *log.Logger
}

var _ LoggerI = (*Logger)(nil)

func New(level string) *Logger {
	l := log.New(os.Stderr)

	// some options
	l.SetReportTimestamp(true)
	l.SetFormatter(log.JSONFormatter)

	switch strings.ToLower(level) {
	case "debug":
		l.SetLevel(log.DebugLevel)
	case "info":
		l.SetLevel(log.InfoLevel)
	case "warn":
		l.SetLevel(log.WarnLevel)
	case "error":
		l.SetLevel(log.ErrorLevel)
	case "fatal":
		l.SetLevel(log.FatalLevel)
	default:
		l.SetLevel(log.InfoLevel)
	}

	return &Logger{
		logger: l,
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}
