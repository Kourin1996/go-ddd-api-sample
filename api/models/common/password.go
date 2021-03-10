package common

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/Kourin1996/go-crud-api-sample/api/constants"
)

type Password struct {
	HashedPassword []byte
}

func NewPassword(rawPassword string) (Password, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(rawPassword), constants.BCRYPT_COST)
	if err != nil {
		return Password{}, err
	}
	return Password{HashedPassword: hashed}, nil
}

func (p *Password) Match(password string) bool {
	err := bcrypt.CompareHashAndPassword(p.HashedPassword, []byte(password))
	return err == nil
}
