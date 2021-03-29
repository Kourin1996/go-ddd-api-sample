package services

import (
	"github.com/Kourin1996/go-crud-api-sample/api/models/auth"
	"github.com/stretchr/testify/mock"
)

var _ auth.IAuthService = &MockAuthService{}

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) SignUp(dto *auth.SignUpDto) (*auth.AuthResult, error) {
	ret := m.Called(dto)
	return ret.Get(0).(*auth.AuthResult), ret.Error(1)
}

func (m *MockAuthService) SignIn(dto *auth.SignInDto) (*auth.AuthResult, error) {
	ret := m.Called(dto)
	return ret.Get(0).(*auth.AuthResult), ret.Error(1)
}
