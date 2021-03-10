package auth

import "github.com/Kourin1996/go-crud-api-sample/api/models/user"

type IAuthService interface {
	SignUp(*SignUpCommand) (*user.UserModel, error)
	SignIn(*SignInCommand) (*user.UserModel, error)
}
