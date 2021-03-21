package helper

import "github.com/Kourin1996/go-crud-api-sample/api/common/log"

const DEFAULT_HEADER = "${time_rfc3339} ${level}"

func init() {
	lm := log.NewLoggerManager(&log.LoggerOption{
		Level:  log.PRINT,
		Header: DEFAULT_HEADER,
	})
	log.SetGlobalLogger(lm)
}
