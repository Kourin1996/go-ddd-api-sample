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

func (r *BookRepository) Create(book *book.Book) (*book.Book, error) {
	fmt.Printf("Create Book: %+v\n", book)
	_, err := r.DB.Model(book).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Created Book: %+v\n", book)

	return book, nil
}

func (r *BookRepository) Get(id int64) (*book.Book, error) {
	book := book.NewEmptyBook()

	err := r.DB.Model(book).Where("id = ?", id).Limit(1).Select()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Got Book: %d %+v\n", id, book)

	return book, nil
}

func (r *BookRepository) Update(id int64, updateBook *book.UpdateBook) (*book.Book, error) {
	book := book.NewEmptyBook()

	fmt.Printf("UpdateBook: %+v %+v\n", updateBook, book)
	_, err := r.DB.Model(updateBook).Where("id = ?", id).Returning("*").UpdateNotZero(book)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Res UpdatedBook: %+v\n", book)

	return book, nil
}

func (r *BookRepository) Delete(id int64) error {
	fmt.Printf("Delete: %d\n", id)
	_, err := r.DB.Model(&book.Book{}).Where("id = ?", id).Delete()
	return err
}
