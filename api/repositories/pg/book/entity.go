package book

import (
	"github.com/Kourin1996/go-crud-api-sample/api/models"
)

type BookEntity struct {
	tableName   struct{} `pg:"books"`
	ID          *int
	Name        string
	Description string
	Price       int64
}

func ToEntity(b *models.Book) *BookEntity {
	return &BookEntity{
		Name:        b.Name,
		Description: b.Description,
		Price:       b.Price,
	}
}

func ToModel(b *BookEntity) *models.Book {
	return &models.Book{
		ID:          *b.ID,
		Name:        b.Name,
		Description: b.Description,
		Price:       b.Price,
	}
}
