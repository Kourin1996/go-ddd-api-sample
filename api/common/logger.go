package common

import (
	"io"

	"github.com/labstack/gommon/log"
)

type Logger interface {
	Print(i ...interface{})
	Printf(format string, args ...interface{})
	Printj(j map[string]interface{})
	Debug(i ...interface{})
	Debugf(format string, args ...interface{})
	Debugj(j map[string]interface{})
	Info(i ...interface{})
	Infof(format string, args ...interface{})
	Infoj(j map[string]interface{})
	Warn(i ...interface{})
	Warnf(format string, args ...interface{})
	Warnj(j map[string]interface{})
	Error(i ...interface{})
	Errorf(format string, args ...interface{})
	Errorj(j map[string]interface{})
	Fatal(i ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalj(j map[string]interface{})
	Panic(i ...interface{})
	Panicf(format string, args ...interface{})
	Panicj(j map[string]interface{})
}

type NormalLogFunc func(i ...interface{})
type FormatLogFunc func(format string, args ...interface{})
type JsonLogFunc func(j log.JSON)

type LogLevel uint8

const (
	PRINT LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	OFF
	PANIC
	FATAL
)

type LoggerOption struct {
	Prefix string
	Level  LogLevel
	Header string
	Output io.Writer
}

type LabstackLogger struct {
	options []*LoggerOption
	loggers []*log.Logger
}

func NewLabstackLogger(options ...*LoggerOption) Logger {
	loggers := make([]*log.Logger, len(options))
	for i, o := range options {
		loggers[i] = log.New(o.Prefix)
		loggers[i].SetLevel(log.Lvl(o.Level))

		if len(o.Header) > 0 {
			loggers[i].SetHeader(o.Header)
		}
		if o.Output != nil {
			loggers[i].SetOutput(o.Output)
		}
	}

	return &LabstackLogger{options: options, loggers: loggers}
}

func (l *LabstackLogger) Print(i ...interface{}) {
	l.logNormal(PRINT, getPrint, i...)
}

func (l *LabstackLogger) Printf(format string, args ...interface{}) {
	l.logFormat(PRINT, getPrintf, format, args)
}

func (l *LabstackLogger) Printj(j map[string]interface{}) {
	l.logJson(PRINT, getPrintj, j)
}

func (l *LabstackLogger) Debug(i ...interface{}) {
	l.logNormal(DEBUG, getDebug, i...)
}

func (l *LabstackLogger) Debugf(format string, args ...interface{}) {
	l.logFormat(DEBUG, getDebugf, format, args)
}

func (l *LabstackLogger) Debugj(j map[string]interface{}) {
	l.logJson(DEBUG, getDebugj, j)
}

func (l *LabstackLogger) Info(i ...interface{}) {
	l.logNormal(INFO, getInfo, i...)
}

func (l *LabstackLogger) Infof(format string, args ...interface{}) {
	l.logFormat(INFO, getInfof, format, args)
}

func (l *LabstackLogger) Infoj(j map[string]interface{}) {
	l.logJson(INFO, getInfoj, j)
}

func (l *LabstackLogger) Warn(i ...interface{}) {
	l.logNormal(WARN, getWarn, i...)
}

func (l *LabstackLogger) Warnf(format string, args ...interface{}) {
	l.logFormat(WARN, getWarnf, format, args)
}

func (l *LabstackLogger) Warnj(j map[string]interface{}) {
	l.logJson(WARN, getWarnj, j)
}

func (l *LabstackLogger) Error(i ...interface{}) {
	l.logNormal(ERROR, getError, i...)
}

func (l *LabstackLogger) Errorf(format string, args ...interface{}) {
	l.logFormat(ERROR, getErrorf, format, args)
}

func (l *LabstackLogger) Errorj(j map[string]interface{}) {
	l.logJson(ERROR, getErrorj, j)
}

func (l *LabstackLogger) Panic(i ...interface{}) {
	l.logNormal(PANIC, getPanic, i...)
}

func (l *LabstackLogger) Panicf(format string, args ...interface{}) {
	l.logFormat(PANIC, getPanicf, format, args)
}

func (l *LabstackLogger) Panicj(j map[string]interface{}) {
	l.logJson(PANIC, getPanicj, j)
}

func (l *LabstackLogger) Fatal(i ...interface{}) {
	l.logNormal(FATAL, getFatal, i...)
}

func (l *LabstackLogger) Fatalf(format string, args ...interface{}) {
	l.logFormat(FATAL, getFatalf, format, args)
}

func (l *LabstackLogger) Fatalj(j map[string]interface{}) {
	l.logJson(FATAL, getFatalj, j)
}

func (l *LabstackLogger) getLoggers(level LogLevel) []*log.Logger {
	loggers := make([]*log.Logger, 0, len(l.loggers))
	for i, o := range l.options {
		if o.Level <= level {
			loggers = append(loggers, l.loggers[i])
		}
	}
	return loggers
}

func (l *LabstackLogger) logNormal(level LogLevel, f func(*log.Logger) NormalLogFunc, i ...interface{}) {
	loggers := l.getLoggers(level)
	for _, l := range loggers {
		f(l)(i...)
	}
}

func (l *LabstackLogger) logFormat(level LogLevel, f func(*log.Logger) FormatLogFunc, format string, args ...interface{}) {
	loggers := l.getLoggers(level)
	for _, l := range loggers {
		f(l)(format, args...)
	}
}

func (l *LabstackLogger) logJson(level LogLevel, f func(*log.Logger) JsonLogFunc, json map[string]interface{}) {
	loggers := l.getLoggers(level)
	for _, l := range loggers {
		f(l)(log.JSON(json))
	}
}

// Get Log Receiver from logger
func getPrint(l *log.Logger) NormalLogFunc {
	return l.Print
}

func getPrintf(l *log.Logger) FormatLogFunc {
	return l.Printf
}

func getPrintj(l *log.Logger) JsonLogFunc {
	return l.Printj
}

func getDebug(l *log.Logger) NormalLogFunc {
	return l.Debug
}

func getDebugf(l *log.Logger) FormatLogFunc {
	return l.Debugf
}

func getDebugj(l *log.Logger) JsonLogFunc {
	return l.Debugj
}

func getInfo(l *log.Logger) NormalLogFunc {
	return l.Info
}

func getInfof(l *log.Logger) FormatLogFunc {
	return l.Infof
}

func getInfoj(l *log.Logger) JsonLogFunc {
	return l.Infoj
}

func getWarn(l *log.Logger) NormalLogFunc {
	return l.Warn
}

func getWarnf(l *log.Logger) FormatLogFunc {
	return l.Warnf
}

func getWarnj(l *log.Logger) JsonLogFunc {
	return l.Warnj
}

func getError(l *log.Logger) NormalLogFunc {
	return l.Error
}

func getErrorf(l *log.Logger) FormatLogFunc {
	return l.Errorf
}

func getErrorj(l *log.Logger) JsonLogFunc {
	return l.Errorj
}

func getPanic(l *log.Logger) NormalLogFunc {
	return l.Panic
}

func getPanicf(l *log.Logger) FormatLogFunc {
	return l.Panicf
}

func getPanicj(l *log.Logger) JsonLogFunc {
	return l.Panicj
}

func getFatal(l *log.Logger) NormalLogFunc {
	return l.Fatal
}

func getFatalf(l *log.Logger) FormatLogFunc {
	return l.Fatalf
}

func getFatalj(l *log.Logger) JsonLogFunc {
	return l.Fatalj
}
