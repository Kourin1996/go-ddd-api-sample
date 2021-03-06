package pg

import (
	"github.com/go-pg/pg/v10"
)

type Repository struct {
	DB *pg.DB
}

type Config struct {
	Address  string
	User     string
	Password string
	Database string
}

func NewDb(config Config) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     config.Address,
		User:     config.User,
		Password: config.Password,
		Database: config.Database,
	})
	return db
}
