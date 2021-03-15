package book

import "github.com/Kourin1996/go-crud-api-sample/api/models/jwt"

type IBookService interface {
	Get(hashId string) (*Book, error)
	Create(*jwt.TokenData, *CreateBookDto) (*Book, error)
	Update(*jwt.TokenData, string, *UpdateBookDto) (*Book, error)
	Delete(*jwt.TokenData, string) error
}
