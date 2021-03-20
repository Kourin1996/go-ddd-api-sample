package log

import (
	"io"

	"github.com/labstack/gommon/log"
)

type LoggerOption struct {
	Prefix string
	Level  LogLevel
	Header string
	Output io.Writer
}

type LoggerManager struct {
	options []*LoggerOption
	loggers []*log.Logger
}

func NewLoggerManager(options ...*LoggerOption) ILoggerManager {
	loggers := make([]*log.Logger, len(options))
	for i, o := range options {
		l := log.New(o.Prefix)
		l.SetLevel(log.Lvl(o.Level))

		if len(o.Header) > 0 {
			l.SetHeader(o.Header)
		}
		if o.Output != nil {
			l.SetOutput(o.Output)
		}
		loggers[i] = l
	}

	return &LoggerManager{options: options, loggers: loggers}
}

func (l *LoggerManager) GetLoggers(level LogLevel) []*log.Logger {
	loggers := make([]*log.Logger, 0, len(l.loggers))
	for i, o := range l.options {
		if o.Level <= level {
			loggers = append(loggers, l.loggers[i])
		}
	}
	return loggers
}

func logNormal(level LogLevel, l ILoggerManager, f func(*log.Logger) NormalLogFunc, i ...interface{}) {
	for _, l := range l.GetLoggers(level) {
		f(l)(i...)
	}
}

func logFormat(level LogLevel, l ILoggerManager, f func(*log.Logger) FormatLogFunc, format string, args ...interface{}) {
	for _, l := range l.GetLoggers(level) {
		f(l)(format, args...)
	}
}

func logJson(level LogLevel, l ILoggerManager, f func(*log.Logger) JsonLogFunc, json map[string]interface{}) {
	for _, l := range l.GetLoggers(level) {
		f(l)(log.JSON(json))
	}
}
