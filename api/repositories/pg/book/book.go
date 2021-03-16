package book

import (
	"errors"

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
	_, err := r.DB.Model(book).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepository) Get(id int64) (*book.Book, error) {
	book := book.NewEmptyBook()

	err := r.DB.Model(book).Column("book.*").Relation("User").Where("book.id = ?", id).Limit(1).Select()
	if errors.Is(pg.ErrNoRows, err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepository) GetBooks(query *book.GetBookQuery) ([]*book.Book, error) {
	books := make([]*book.Book, query.Limit)
	for i, _ := range books {
		books[i] = book.NewEmptyBook()
	}

	err := r.DB.Model(&books).Column("book.*").Relation("User").Offset(query.Offset).Limit(query.Limit).Select()
	if errors.Is(pg.ErrNoRows, err) {
		return []*book.Book{}, nil
	}
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) Update(id int64, updateBook *book.UpdateBook) (*book.Book, error) {
	book := book.NewEmptyBook()

	_, err := r.DB.Model(updateBook).Where("id = ?", id).Returning("*").UpdateNotZero(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepository) Delete(id int64) error {
	_, err := r.DB.Model(&book.Book{}).Where("id = ?", id).Delete()
	return err
}
