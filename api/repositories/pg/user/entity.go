package user

import "time"

type UserEntity struct {
	tableName struct{} `pg:"users"`
	Id        int64
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
