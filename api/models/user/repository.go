package user

type IUserRepository interface {
	Create(*CreateUserCommand) (*UserModel, error)
	Get(int32) (*UserModel, error)
	GetByUsername(string) (*UserModel, error)
	Delete(int32) error
}
