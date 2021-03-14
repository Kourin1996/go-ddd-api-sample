package auth

import (
	"fmt"
	"time"

	"github.com/Kourin1996/go-crud-api-sample/api/constants"
	"github.com/Kourin1996/go-crud-api-sample/api/models/auth"
	"github.com/Kourin1996/go-crud-api-sample/api/models/user"
	"github.com/dgrijalva/jwt-go"
)

func createJWT(user *user.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["hash_id"] = user.HashId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(constants.JWT_SECRET))
}

type AuthService struct {
	userService user.IUserService
}

func NewAuthService(userService user.IUserService) auth.IAuthService {
	return &AuthService{userService: userService}
}

func (s *AuthService) SignUp(dto *auth.SignUpDto) (*auth.AuthResult, error) {
	fmt.Printf("AuthService::SignUp dto => %+v\n", dto)
	currentUser, err := s.userService.GetByUsername(dto.Username)
	fmt.Printf("AuthService::SignUp currentUser, error => %+v, %+v\n", currentUser, err)
	if err != nil {
		return nil, err
	}
	if currentUser != nil {
		return nil, fmt.Errorf("Username is already registered")
	}

	newUser, err := user.NewUser(dto.Username, dto.Password)
	fmt.Printf("AuthService::SignUp newUser, error => %+v, %+v\n", newUser, err)
	if err != nil {
		return nil, err
	}

	newUser, err = s.userService.Create(newUser)
	fmt.Printf("AuthService::SignUp newUser, error => %+v, %+v\n", newUser, err)
	if err != nil {
		return nil, err
	}
	token, err := createJWT(newUser)
	fmt.Printf("AuthService::SignUp token, error => %s, %+v\n", token, err)
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

	token, err := createJWT(currentUser)
	if err != nil {
		return nil, err
	}

	return &auth.AuthResult{Token: token}, nil
}
