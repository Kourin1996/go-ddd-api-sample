package book

import "github.com/Kourin1996/go-crud-api-sample/api/models/common"

const MODEL_NAME = "Book"

type Book struct {
	tableName struct{} `pg:"books"`
	common.BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

func NewEmptyBook() *Book {
	b := &Book{}
	b.Model = MODEL_NAME
	return b
}

type UpdateBook struct {
	tableName struct{} `pg:"books"`
	common.BaseModel
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Price       *int64  `json:"price"`
}

func NewEmptyUpdateBook() *UpdateBook {
	b := &UpdateBook{}
	b.Model = MODEL_NAME
	return b
}
