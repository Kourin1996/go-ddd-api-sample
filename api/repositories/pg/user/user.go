package book

import (
	"errors"

	"github.com/Kourin1996/go-crud-api-sample/api/models/user"
	repositories "github.com/Kourin1996/go-crud-api-sample/api/repositories/pg"
	"github.com/go-pg/pg/v10"
)

type UserRepository struct {
	repositories.Repository
}

func NewUserRepository(db *pg.DB) user.IUserRepository {
	return &UserRepository{repositories.Repository{DB: db}}
}

func (r *UserRepository) Create(u *user.User) (*user.User, error) {
	_, err := r.DB.Model(u).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Get(id int64) (*user.User, error) {
	u := user.NewEmptyUser()

	err := r.DB.Model(u).Where("id = ?", id).Limit(1).Select()
	if errors.Is(pg.ErrNoRows, err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) GetByUsername(username string) (*user.User, error) {
	u := user.NewEmptyUser()

	err := r.DB.Model(u).Where("username = ?", username).Limit(1).Select()
	if errors.Is(pg.ErrNoRows, err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return u, nil
}
