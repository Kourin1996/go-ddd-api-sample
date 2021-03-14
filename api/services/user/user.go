package user

import (
	"github.com/Kourin1996/go-crud-api-sample/api/common"
	"github.com/Kourin1996/go-crud-api-sample/api/constants"
	"github.com/Kourin1996/go-crud-api-sample/api/models/user"
)

type UserService struct {
	userRepo user.IUserRepository
}

func NewUserService(userRepo user.IUserRepository) user.IUserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Create(data *user.User) (*user.User, error) {
	u, err := s.userRepo.Create(data)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserService) GetByHashId(hashId string) (*user.User, error) {
	id, err := common.DecodeHashID(hashId, user.MODEL_NAME, constants.HASHIDS_SALT, constants.HASHIDS_LENGTH)
	if err != nil {
		return nil, err
	}

	u, err := s.userRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UserService) GetByUsername(username string) (*user.User, error) {
	u, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	return u, nil
}
