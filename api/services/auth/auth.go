package auth

import (
	"fmt"

	"github.com/Kourin1996/go-crud-api-sample/api/models/auth"
	jwtToken "github.com/Kourin1996/go-crud-api-sample/api/models/jwt"
	"github.com/Kourin1996/go-crud-api-sample/api/models/user"
)

type AuthService struct {
	userService user.IUserService
}

func NewAuthService(userService user.IUserService) auth.IAuthService {
	return &AuthService{userService: userService}
}

func (s *AuthService) SignUp(dto *auth.SignUpDto) (*auth.AuthResult, error) {
	currentUser, err := s.userService.GetByUsername(dto.Username)
	if err != nil {
		return nil, err
	}
	if currentUser != nil {
		return nil, fmt.Errorf("Username is already registered")
	}

	newUser, err := user.NewUser(dto.Username, dto.Password)
	if err != nil {
		return nil, err
	}

	newUser, err = s.userService.Create(newUser)
	if err != nil {
		return nil, err
	}
	token, err := jwtToken.EncodeJWTToken(newUser)
	if err != nil {
		return nil, err
	}

	return &auth.AuthResult{Token: token}, nil
}

func (s *AuthService) SignIn(dto *auth.SignInDto) (*auth.AuthResult, error) {
	currentUser, err := s.userService.GetByUsername(dto.Username)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, fmt.Errorf("Username or password is wrong")
	}

	if !currentUser.MatchPassword(dto.Password) {
		return nil, fmt.Errorf("Username or password is wrong")
	}

	token, err := jwtToken.EncodeJWTToken(currentUser)
	if err != nil {
		return nil, err
	}

	return &auth.AuthResult{Token: token}, nil
}
