package user

type IUserRepository interface {
	Create(*User) (*User, error)
	Get(int32) (*User, error)
	GetByUsername(string) (*User, error)
	Delete(int32) error
}
