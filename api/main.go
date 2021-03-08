package main

import (
	"log"

	"github.com/Kourin1996/go-crud-api-sample/api/common"
	"github.com/Kourin1996/go-crud-api-sample/api/controllers"
	"github.com/Kourin1996/go-crud-api-sample/api/repositories/pg"
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

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("error happened on loading config: %s", err.Error())
		return
	}

	db := pg.NewDb(config.Db)
	if err = controllers.Start(config.Api, db); err != nil {
		log.Fatalf("error happened on starting API: %s", err.Error())
		return
	}
}
