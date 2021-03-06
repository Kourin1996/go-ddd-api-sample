package book

import (
	"github.com/Kourin1996/go-crud-api-sample/api/models/book"
)

type BookEntity struct {
	tableName   struct{} `pg:"books"`
	ID          *int
	Name        string
	Description string
	Price       int64
}

func ToEntity(b *book.Book) *BookEntity {
	return &BookEntity{
		Name:        b.Name,
		Description: b.Description,
		Price:       b.Price,
	}
}

func ToModel(b *BookEntity) *book.Book {
	return &book.Book{
		ID:          *b.ID,
		Name:        b.Name,
		Description: b.Description,
		Price:       b.Price,
	}
}
