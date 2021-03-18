package common

type Logger interface {
	Prefix() string
	SetPrefix(p string)
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

type LabstackLogger struct {
}

func NewLabstackLogger() Logger {
	return &LabstackLogger{}
}
