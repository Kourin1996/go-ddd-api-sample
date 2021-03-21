package services

import (
	"github.com/Kourin1996/go-crud-api-sample/api/models/book"
	"github.com/stretchr/testify/mock"
)

type MockBookService struct {
	mock.Mock
}

func (m *MockBookService) Create(dto *book.CreateBookDto) (*book.Book, error) {
	ret := m.Called(dto)
	return ret.Get(0).(*book.Book), ret.Error(1)
}

func (m *MockBookService) Get(hashId string) (*book.Book, error) {
	ret := m.Called(hashId)
	return ret.Get(0).(*book.Book), ret.Error(1)
}
