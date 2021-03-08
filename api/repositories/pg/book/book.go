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

func (r *BookRepository) CreateBook(book *book.Book) (*book.Book, error) {
	entity := ToEntity(book)

	_, err := r.DB.Model(entity).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return ToModel(entity), nil
}

func (r *BookRepository) GetBook(id int) (*book.Book, error) {
	entity := new(BookEntity)

	err := r.DB.Model(entity).Where("id = ?", id).Limit(1).Select()
	if err != nil {
		return nil, err
	}
	return ToModel(entity), nil
}

func (r *BookRepository) UpdateBook(id int, book *book.Book) (*book.Book, error) {
	entity := ToEntity(book)
	fmt.Printf("UpdateBook: %+v %s %d\n", book, book.Description, len(book.Description))

	entity.ID = id

	_, err := r.DB.Model(entity).WherePK().Returning("*").UpdateNotZero()
	if err != nil {
		return nil, err
	}
	return ToModel(entity), nil
}

func (r *BookRepository) DeleteBook(id int) error {
	_, err := r.DB.Model(&book.Book{}).Where("id = ?", id).Delete()
	return err
}
