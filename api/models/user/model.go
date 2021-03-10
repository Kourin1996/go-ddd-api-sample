package user

import (
	"time"

	"github.com/Kourin1996/go-crud-api-sample/api/models/common"
)

type UserModel struct {
	ID        common.ID       `json:"id"`
	Username  string          `json:"username"`
	Password  common.Password `json:"-"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
