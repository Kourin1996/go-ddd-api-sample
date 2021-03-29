package log

import (
	"github.com/labstack/gommon/log"
)

type NormalLogFunc func(i ...interface{})
type FormatLogFunc func(format string, args ...interface{})
type JSONLogFunc func(j log.JSON)

func getPrint(l *log.Logger) NormalLogFunc {
	return l.Print
}

func getPrintf(l *log.Logger) FormatLogFunc {
	return l.Printf
}

func getPrintj(l *log.Logger) JSONLogFunc {
	return l.Printj
}

func getDebug(l *log.Logger) NormalLogFunc {
	return l.Debug
}

func getDebugf(l *log.Logger) FormatLogFunc {
	return l.Debugf
}

func getDebugj(l *log.Logger) JSONLogFunc {
	return l.Debugj
}

func getInfo(l *log.Logger) NormalLogFunc {
	return l.Info
}

func getInfof(l *log.Logger) FormatLogFunc {
	return l.Infof
}

func getInfoj(l *log.Logger) JSONLogFunc {
	return l.Infoj
}

func getWarn(l *log.Logger) NormalLogFunc {
	return l.Warn
}

func getWarnf(l *log.Logger) FormatLogFunc {
	return l.Warnf
}

func getWarnj(l *log.Logger) JSONLogFunc {
	return l.Warnj
}

func getError(l *log.Logger) NormalLogFunc {
	return l.Error
}

func getErrorf(l *log.Logger) FormatLogFunc {
	return l.Errorf
}

func getErrorj(l *log.Logger) JSONLogFunc {
	return l.Errorj
}

func getPanic(l *log.Logger) NormalLogFunc {
	return l.Panic
}

func getPanicf(l *log.Logger) FormatLogFunc {
	return l.Panicf
}

func getPanicj(l *log.Logger) JSONLogFunc {
	return l.Panicj
}

func getFatal(l *log.Logger) NormalLogFunc {
	return l.Fatal
}

func getFatalf(l *log.Logger) FormatLogFunc {
	return l.Fatalf
}

func getFatalj(l *log.Logger) JSONLogFunc {
	return l.Fatalj
}
