package book

import (
	"github.com/Kourin1996/go-crud-api-sample/api/models/book"
)

type BookService struct {
	bookRepo book.IBookRepository
}

func NewBookService(bookRepo book.IBookRepository) book.IBookService {
	return &BookService{bookRepo: bookRepo}
}

func (service *BookService) CreateBook(book *book.CreateBookCommand) (*book.BookResult, error) {
	return service.bookRepo.CreateBook(book)
}

func (service *BookService) GetBook(id int32) (*book.BookResult, error) {
	return service.bookRepo.GetBook(id)
}

func (service *BookService) UpdateBook(id int32, book *book.UpdateBookCommand) (*book.BookResult, error) {
	return service.bookRepo.UpdateBook(id, book)
}

func (service *BookService) DeleteBook(id int32) error {
	return service.bookRepo.DeleteBook(id)
}
