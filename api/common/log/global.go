package log

var (
	global = NewLoggerManager()
)

func SetGlobalLogger(l ILoggerManager) {
	global = l
}

func Print(i ...interface{}) {
	logNormal(PRINT, global, getPrint, i...)
}

func Printf(format string, args ...interface{}) {
	logFormat(PRINT, global, getPrintf, format, args)
}

func Printj(j map[string]interface{}) {
	logJSON(PRINT, global, getPrintj, j)
}

func Debug(i ...interface{}) {
	logNormal(DEBUG, global, getDebug, i...)
}

func Debugf(format string, args ...interface{}) {
	logFormat(DEBUG, global, getDebugf, format, args)
}

func Debugj(j map[string]interface{}) {
	logJSON(DEBUG, global, getDebugj, j)
}

func Info(i ...interface{}) {
	logNormal(INFO, global, getInfo, i...)
}

func Infof(format string, args ...interface{}) {
	logFormat(INFO, global, getInfof, format, args)
}

func Infoj(j map[string]interface{}) {
	logJSON(INFO, global, getInfoj, j)
}

func Warn(i ...interface{}) {
	logNormal(WARN, global, getWarn, i...)
}

func Warnf(format string, args ...interface{}) {
	logFormat(WARN, global, getWarnf, format, args)
}

func Warnj(j map[string]interface{}) {
	logJSON(WARN, global, getWarnj, j)
}

func Error(i ...interface{}) {
	logNormal(ERROR, global, getError, i...)
}

func Errorf(format string, args ...interface{}) {
	logFormat(ERROR, global, getErrorf, format, args)
}

func Errorj(j map[string]interface{}) {
	logJSON(ERROR, global, getErrorj, j)
}

func Panic(i ...interface{}) {
	logNormal(PANIC, global, getPanic, i...)
}

func Panicf(format string, args ...interface{}) {
	logFormat(PANIC, global, getPanicf, format, args)
}

func Panicj(j map[string]interface{}) {
	logJSON(PANIC, global, getPanicj, j)
}

func Fatal(i ...interface{}) {
	logNormal(FATAL, global, getFatal, i...)
}

func Fatalf(format string, args ...interface{}) {
	logFormat(FATAL, global, getFatalf, format, args)
}

func Fatalj(j map[string]interface{}) {
	logJSON(FATAL, global, getFatalj, j)
}
