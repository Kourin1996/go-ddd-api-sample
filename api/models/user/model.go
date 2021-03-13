package user

import (
	"github.com/Kourin1996/go-crud-api-sample/api/models/common"
)

type User struct {
	common.BaseModel
	common.Password
	Username string `json:"username"`
}
