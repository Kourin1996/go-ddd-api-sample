package book

import (
	"github.com/Kourin1996/go-crud-api-sample/api/models"
	"github.com/Kourin1996/go-crud-api-sample/api/repositories/book"
)

// todo: move to other dir
type IBookService interface {
	CreateBook(book *models.Book) (*models.Book, error)
	GetBook(id int) (*models.Book, error)
	UpdateBook(id int, book *models.Book) (*models.Book, error)
	DeleteBook(id int) error
}

type BookService struct {
	bookRepo book.IBookRepository
}

func NewBookService(bookRepo book.IBookRepository) IBookService {
	return &BookService{bookRepo: bookRepo}
}

func (service *BookService) CreateBook(book *models.Book) (*models.Book, error) {
	return service.bookRepo.CreateBook(book)
}

func (service *BookService) GetBook(id int) (*models.Book, error) {
	return service.bookRepo.GetBook(id)
}

func (service *BookService) UpdateBook(id int, book *models.Book) (*models.Book, error) {
	return service.bookRepo.UpdateBook(id, book)
}

func (service *BookService) DeleteBook(id int) error {
	return service.bookRepo.DeleteBook(id)
}
