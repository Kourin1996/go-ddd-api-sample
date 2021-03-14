package user

import (
	"fmt"

	"github.com/Kourin1996/go-crud-api-sample/api/models/common"
)

const MODEL_NAME = "User"

type User struct {
	tableName struct{} `pg:"users"`
	common.BaseModel
	common.Password
	Username string `json:"username"`
}

func NewEmptyUser() *User {
	u := &User{}
	u.Model = MODEL_NAME
	return u
}

func NewUser(
	username string,
	password string,
) (*User, error) {

	if len(username) < 6 || len(username) > 255 {
		return nil, fmt.Errorf("Invalid username")
	}
	if len(password) < 8 || len(password) > 255 {
		return nil, fmt.Errorf("Invalid password")
	}

	u := NewEmptyUser()
	u.Username = username

	err := u.SetRawPassword(password)
	if err != nil {
		return nil, err
	}

	return u, nil
}
