package repositories

import (
	"github.com/go-pg/pg/v10"
)

type Repository struct {
	DB *pg.DB
}
