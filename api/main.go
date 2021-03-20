package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Kourin1996/go-crud-api-sample/api/common"
	"github.com/Kourin1996/go-crud-api-sample/api/common/log"
	"github.com/Kourin1996/go-crud-api-sample/api/controllers"
	"github.com/Kourin1996/go-crud-api-sample/api/repositories/pg"
	gopg "github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Db  pg.Config
	Api controllers.Config
}

func LoadConfig() (Config, error) {
	var config Config

	err := godotenv.Load()
	if err != nil {
		return config, err
	}

	config.Db.Address = common.LoadEnv("DB_ADDRESS", ":5432")
	config.Db.User = common.LoadEnv("DB_USER", "")
	config.Db.Password = common.LoadEnv("DB_PASSWORD", "")
	config.Db.Database = common.LoadEnv("DB_DATABASE", "")
	config.Api.Address = common.LoadEnv("API_ADDRESS", ":8080")

	return config, err
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *gopg.QueryEvent) (context.Context, error) {
	query, err := q.FormattedQuery()
	if err == nil {
		log.Debug(string(query))
	}
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *gopg.QueryEvent) error {
	return nil
}

func main() {
	errorLogFile, err := os.Create("error.log")
	if err != nil {
		fmt.Println("error")
	}

	loggerOptions := []*log.LoggerOption{
		{
			Prefix: "Debug",
			Level:  log.DEBUG,
			Header: "[Debug]",
		},
		{
			Prefix: "Error",
			Level:  log.ERROR,
			Header: "[Error]",
			Output: errorLogFile,
		},
	}
	logger := log.NewLoggerManager(loggerOptions...)
	log.SetGlobalLogger(logger)

	log.Print("test")
	log.Debugf("This is debug test 1+1=%d", 1+1)
	log.Errorj(map[string]interface{}{
		"Msg":   "This is error test",
		"Hello": "World",
	})

	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("error happened on loading config: %s", err.Error())
		return
	}

	db := pg.NewDb(config.Db)
	db.AddQueryHook(&dbLogger{})

	if err = controllers.Start(config.Api, db); err != nil {
		log.Fatalf("error happened on starting API: %s", err.Error())
		return
	}
}
