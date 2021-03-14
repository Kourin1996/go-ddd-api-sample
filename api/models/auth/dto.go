package auth

type SignUpDto struct {
	Username string `json:"username" validate:"required,min=6,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type SignInDto struct {
	Username string `json:"username" validate:"required,min=6,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}
