package common

import (
	"fmt"

	"github.com/Kourin1996/go-crud-api-sample/api/constants"
	"github.com/go-pg/pg/v10/orm"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	RawPassword       string `json:"-" pg:"-"`
	EncryptedPassword string `json:"-" pg:"password"`
}

func (p *Password) BeforeInsert(db orm.DB) error {
	if len(p.RawPassword) > 0 {
		return p.SetRawPassword(p.EncryptedPassword)
	}
	return nil
}

func (p *Password) BeforeUpdate(db orm.DB) error {
	if len(p.RawPassword) > 0 {
		return p.SetRawPassword(p.EncryptedPassword)
	}
	return nil
}

func (p *Password) SetRawPassword(rawPassword string) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), constants.BCRYPT_COST)
	if err != nil {
		return fmt.Errorf("Failed to encrypt password")
	}
	p.RawPassword = rawPassword
	p.EncryptedPassword = string(encryptedPassword)
	return nil
}

func (p *Password) MatchPassword(rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.EncryptedPassword), []byte(rawPassword))
	return err == nil
}
