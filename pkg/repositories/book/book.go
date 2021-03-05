package book

import (
	"github.com/Kourin1996/go-crud-api-sample/pkg/models"
	"github.com/Kourin1996/go-crud-api-sample/pkg/repositories"
	"github.com/go-pg/pg/v10"
)

//todo move to interface layer
type IBookRepository interface {
	CreateBook(book *models.Book) (*models.Book, error)
	GetBook(id int) (*models.Book, error)
	UpdateBook(id int, newData *models.Book) (*models.Book, error)
	DeleteBook(id int) error
}

type BookRepository struct {
	repositories.Repository
}

func NewRepository(db *pg.DB) IBookRepository {
	return &BookRepository{repositories.Repository{DB: db}}
}

func (r *BookRepository) CreateBook(book *models.Book) (*models.Book, error) {
	entity := ToEntity(book)

	_, err := r.DB.Model(entity).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return ToModel(entity), nil
}

func (r *BookRepository) GetBook(id int) (*models.Book, error) {
	entity := new(BookEntity)

	err := r.DB.Model(entity).Where("id = ?", id).Limit(1).Select()
	if err != nil {
		return nil, err
	}
	return ToModel(entity), nil
}

func (r *BookRepository) UpdateBook(id int, book *models.Book) (*models.Book, error) {
	entity := ToEntity(book)
	entity.ID = &id

	_, err := r.DB.Model(entity).WherePK().Returning("*").Update()
	if err != nil {
		return nil, err
	}
	return ToModel(entity), nil
}

func (r *BookRepository) DeleteBook(id int) error {
	_, err := r.DB.Model(&models.Book{}).Where("id = ?", id).Delete()
	return err
}
