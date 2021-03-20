package log

import "github.com/labstack/gommon/log"

type ILoggerManager interface {
	GetLoggers(level LogLevel) []*log.Logger
}

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
