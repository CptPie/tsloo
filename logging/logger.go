package logging

import (
	"fmt"
	"os"

	clog "github.com/charmbracelet/log"
)

type Logger struct {
	consoleLogger clog.Logger
	fileLogger    clog.Logger
}

func New(logfile os.File) *Logger {
	console := clog.NewWithOptions(os.Stderr, clog.Options{
		ReportTimestamp: true,
	})
	file := clog.NewWithOptions(&logfile, clog.Options{
		ReportTimestamp: true,
	})
	return &Logger{
		consoleLogger: *console,
		fileLogger:    *file,
	}
}

func (l Logger) Debug(s string, v ...interface{}) {
	if len(v) != 0 {
		l.consoleLogger.Debug(fmt.Sprintf(s, v...))
		l.fileLogger.Debug(fmt.Sprintf(s, v...))
	} else {
		l.consoleLogger.Debug(s)
		l.fileLogger.Debug(s)
	}
}

func (l Logger) Info(s string, v ...interface{}) {
	if len(v) != 0 {
		l.consoleLogger.Info(fmt.Sprintf(s, v...))
		l.fileLogger.Info(fmt.Sprintf(s, v...))
	} else {
		l.consoleLogger.Info(s)
		l.fileLogger.Info(s)
	}
}

func (l Logger) Warn(s string, v ...interface{}) {
	if len(v) != 0 {
		l.consoleLogger.Warn(fmt.Sprintf(s, v...))
		l.fileLogger.Warn(fmt.Sprintf(s, v...))
	} else {
		l.consoleLogger.Warn(s)
		l.fileLogger.Warn(s)
	}
}

func (l Logger) Error(s string, v ...interface{}) {
	if len(v) != 0 {
		l.consoleLogger.Error(fmt.Sprintf(s, v...))
		l.fileLogger.Error(fmt.Sprintf(s, v...))
	} else {
		l.consoleLogger.Error(s)
		l.fileLogger.Error(s)
	}
}
