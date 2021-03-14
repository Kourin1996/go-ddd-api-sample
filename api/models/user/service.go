package user

type IUserService interface {
	Create(*User) (*User, error)
	GetByHashId(string) (*User, error)
	GetByUsername(string) (*User, error)
}
