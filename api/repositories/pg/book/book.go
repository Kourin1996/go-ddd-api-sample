package book

import (
	"fmt"

	"github.com/Kourin1996/go-crud-api-sample/api/models/book"
	repositories "github.com/Kourin1996/go-crud-api-sample/api/repositories/pg"
	"github.com/go-pg/pg/v10"
)

type BookRepository struct {
	repositories.Repository
}

func NewRepository(db *pg.DB) book.IBookRepository {
	return &BookRepository{repositories.Repository{DB: db}}
}

func (r *BookRepository) CreateBook(command *book.CreateBookCommand) (*book.BookModel, error) {
	entity := ToEntityFromCreateBookCommand(command)

	fmt.Printf("CreateBook: %+v\n", entity)
	_, err := r.DB.Model(entity).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Res CreateBook: %+v\n", entity)

	return ToBookModel(entity), nil
}

func (r *BookRepository) GetBook(id int32) (*book.BookModel, error) {
	entity := new(BookEntity)

	err := r.DB.Model(entity).Where("id = ?", id).Limit(1).Select()
	if err != nil {
		return nil, err
	}

	return ToBookModel(entity), nil
}

func (r *BookRepository) UpdateBook(id int32, command *book.UpdateBookCommand) (*book.BookModel, error) {
	entity := ToEntityFromUpdateBookCommand(id, command)

	fmt.Printf("UpdateBook: %+v\n", entity)
	_, err := r.DB.Model(entity).WherePK().Returning("*").UpdateNotZero()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Res UpdateBook: %+v\n", entity)

	return ToBookModel(entity), nil
}

func (r *BookRepository) DeleteBook(id int32) error {
	_, err := r.DB.Model(&BookEntity{}).Where("id = ?", id).Delete()
	return err
}
