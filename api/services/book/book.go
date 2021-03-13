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

func (service *BookService) Create(dto *book.CreateBookDto) (*book.Book, error) {
	book := book.NewEmptyBook()
	book.Name = dto.Name
	book.Description = dto.Description
	book.Price = dto.Price

	return service.bookRepo.Create(book)
}

func (service *BookService) Get(hashId string) (*book.Book, error) {
	b := book.NewEmptyBook()
	b.SetHashId(hashId)

	return service.bookRepo.Get(b.ID)
}

func (service *BookService) Update(hashId string, dto *book.UpdateBookDto) (*book.Book, error) {
	book := book.NewEmptyUpdateBook()
	book.SetHashId(hashId)
	book.Name = dto.Name
	book.Description = dto.Description
	book.Price = dto.Price

	return service.bookRepo.Update(book.ID, book)
}

func (service *BookService) Delete(hashId string) error {
	b := book.NewEmptyBook()
	b.SetHashId(hashId)

	return service.bookRepo.Delete(b.ID)
}
