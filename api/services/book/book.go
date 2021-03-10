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

func (service *BookService) Create(book *book.CreateBookCommand) (*book.BookModel, error) {
	return service.bookRepo.Create(book)
}

func (service *BookService) Get(id int32) (*book.BookModel, error) {
	return service.bookRepo.Get(id)
}

func (service *BookService) Update(id int32, book *book.UpdateBookCommand) (*book.BookModel, error) {
	return service.bookRepo.Update(id, book)
}

func (service *BookService) Delete(id int32) error {
	return service.bookRepo.Delete(id)
}
