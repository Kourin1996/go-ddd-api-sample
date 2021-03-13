package auth

type IAuthService interface {
	SignUp(*SignUpDto) (*AuthResult, error)
	SignIn(*SignInDto) (*AuthResult, error)
}
