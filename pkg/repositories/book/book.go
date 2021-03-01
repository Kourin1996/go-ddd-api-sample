package book

import (
	"github.com/Kourin1996/go-crud-api-sample/pkg/models"
	"github.com/Kourin1996/go-crud-api-sample/pkg/repositories"
	"github.com/go-pg/pg/v10"
)

type IBookRepository interface {
	CreateBook(book *models.Book) error
	GetBook(id int) (*models.Book, error)
	UpdateBook(id int, newData *models.Book) error
	DeleteBook(id int) error
}

type BookRepository struct {
	repositories.Repository
}

func NewRepository(db *pg.DB) IBookRepository {
	return &BookRepository{repositories.Repository{DB: db}}
}

func (r *BookRepository) CreateBook(book *models.Book) error {
	_, err := r.DB.Model(book).Returning("*").Insert()
	return err
}

func (r *BookRepository) GetBook(id int) (*models.Book, error) {
	book := new(models.Book)
	err := r.DB.Model(book).Where("id = ?", id).Limit(1).Select()
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepository) UpdateBook(id int, book *models.Book) error {
	book.ID = id

	_, err := r.DB.Model(book).WherePK().Returning("*").Update()
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) DeleteBook(id int) error {
	_, err := r.DB.Model(&models.Book{}).Where("id = ?", id).Delete()
	return err
}
