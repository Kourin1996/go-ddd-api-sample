package models

import (
	"time"

	"github.com/Kourin1996/go-crud-api-sample/api/models/book"
)

func NewFakeBookWithId(id int64, name string, description string, price int64, createdAt string, updatedAt string) *book.Book {
	ca, _ := time.Parse(time.RFC3339, createdAt)
	ua, _ := time.Parse(time.RFC3339, updatedAt)

	b, _ := book.NewBook(id, &book.CreateBookDto{
		Name:        name,
		Description: description,
		Price:       price,
	})
	b.SetId(id)
	b.CreatedAt = ca
	b.UpdatedAt = ua

	return b
}
