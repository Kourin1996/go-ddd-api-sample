package auth

import (
	"github.com/Kourin1996/go-crud-api-sample/api/models/auth"
	"github.com/Kourin1996/go-crud-api-sample/api/models/user"
)

type AuthService struct {
	userRepo user.IUserRepository
}

func NewAuthService(userRepo user.IUserRepository) auth.IAuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) SignUp(command *auth.SignUpCommand) (*user.UserModel, error) {
	return nil, nil
}

func (s *AuthService) SignIn(command *auth.SignInCommand) (*user.UserModel, error) {
	return nil, nil
}
