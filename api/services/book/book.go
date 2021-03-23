package book

import (
	"fmt"

	"github.com/Kourin1996/go-crud-api-sample/api/common"
	"github.com/Kourin1996/go-crud-api-sample/api/constants"
	"github.com/Kourin1996/go-crud-api-sample/api/models/book"
	"github.com/Kourin1996/go-crud-api-sample/api/models/jwt"
	"github.com/Kourin1996/go-crud-api-sample/api/models/user"
)

type BookService struct {
	bookRepo book.IBookRepository
}

func NewBookService(bookRepo book.IBookRepository) book.IBookService {
	return &BookService{bookRepo: bookRepo}
}

func (s *BookService) Get(hashId string) (*book.Book, error) {
	b := book.NewEmptyBook()
	if err := b.SetHashId(hashId); err != nil {
		return nil, err
	}

	return s.bookRepo.Get(b.ID)
}

func (s *BookService) GetBooks(dto *book.GetBooksDto) ([]*book.Book, error) {
	query := &book.GetBookQuery{Offset: 0, Limit: 10}
	if dto.Number != nil {
		query.Limit = *dto.Number
	}
	if dto.Page != nil && (*dto.Page) >= 1 {
		query.Offset = query.Limit * (*dto.Page - 1)
	}

	return s.bookRepo.GetBooks(query)
}

func (s *BookService) Create(tokenData *jwt.TokenData, dto *book.CreateBookDto) (*book.Book, error) {
	userId, err := common.DecodeHashID(tokenData.HashId, user.MODEL_NAME, constants.HASHIDS_SALT, constants.HASHIDS_LENGTH)
	if err != nil {
		return nil, err
	}

	book := book.NewEmptyBook()
	book.Name = dto.Name
	book.Description = dto.Description
	book.Price = dto.Price
	book.UserId = userId

	return s.bookRepo.Create(book)
}

func (s *BookService) Update(tokenData *jwt.TokenData, hashId string, dto *book.UpdateBookDto) (*book.Book, error) {
	userId, err := common.DecodeHashID(tokenData.HashId, user.MODEL_NAME, constants.HASHIDS_SALT, constants.HASHIDS_LENGTH)
	if err != nil {
		return nil, err
	}

	book := book.NewEmptyUpdateBook()
	book.SetHashId(hashId)
	book.Name = dto.Name
	book.Description = dto.Description
	book.Price = dto.Price

	b, err := s.bookRepo.Get(book.ID)
	if err != nil {
		return nil, err
	}
	if b.UserId != userId {
		return nil, fmt.Errorf("Cannot update data")
	}

	return s.bookRepo.Update(book.ID, book)
}

func (s *BookService) Delete(tokenData *jwt.TokenData, hashId string) error {
	userId, err := common.DecodeHashID(tokenData.HashId, user.MODEL_NAME, constants.HASHIDS_SALT, constants.HASHIDS_LENGTH)
	if err != nil {
		return err
	}

	b := book.NewEmptyBook()
	b.SetHashId(hashId)

	b, err = s.bookRepo.Get(b.ID)
	if err != nil {
		return err
	}
	if b.UserId != userId {
		return fmt.Errorf("Cannot delete data")
	}

	return s.bookRepo.Delete(b.ID)
}
