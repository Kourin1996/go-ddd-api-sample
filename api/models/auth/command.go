package auth

type SignUpCommand struct {
	Username string
	Password string
}

type SignInCommand struct {
	Username string
	Password string
}
