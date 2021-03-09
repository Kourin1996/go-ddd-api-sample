package book

import (
	"time"

	"github.com/Kourin1996/go-crud-api-sample/api/models/book"
)

// Case1: use pointer
type BookEntity struct {
	tableName   struct{} `pg:"books"`
	ID          int32
	Name        *string
	Description *string
	Price       *int64
	//todo add
	UserId    int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToEntityFromCreateBookCommand(command *book.CreateBookCommand) *BookEntity {
	return &BookEntity{
		Name:        &command.Name,
		Description: &command.Description,
		Price:       &command.Price,
	}
}

func ToEntityFromUpdateBookCommand(ID int32, command *book.UpdateBookCommand) *BookEntity {
	return &BookEntity{
		ID:          ID,
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
	}
}

func ToBookModel(entity *BookEntity) *book.BookModel {
	return &book.BookModel{
		ID:          entity.ID,
		Name:        *entity.Name,
		Description: *entity.Description,
		Price:       *entity.Price,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

// Case2: Create new UpdateBookEntity
