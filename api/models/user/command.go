package user

import "github.com/Kourin1996/go-crud-api-sample/api/models/common"

type CreateUserCommand struct {
	Username string          `json:"username"`
	Password common.Password `json:"-"`
}
