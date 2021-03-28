package common

import (
	"context"
	"time"

	"github.com/Kourin1996/go-crud-api-sample/api/common"
	"github.com/Kourin1996/go-crud-api-sample/api/constants"
	"github.com/Kourin1996/go-crud-api-sample/api/models/errors"
	"github.com/go-pg/pg/v10/orm"
)

type BaseModel struct {
	Model     string    `json:"-" pg:"-"`
	ID        int64     `json:"-" pg:",pk"`
	HashId    string    `json:"hash_id" pg:"-"`
	CreatedAt time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt time.Time `json:"updated_at" pg:"updated_at"`
	DeletedAt time.Time `json:"-" pg:"deleted_at"`
}

func (m *BaseModel) BeforeUpdate(db orm.DB) error {
	m.UpdatedAt = time.Now()
	if len(m.HashId) > 0 && m.ID == 0 {
		return m.SetHashId(m.HashId)
	}
	return nil
}

func (m *BaseModel) AfterScan(ctx context.Context) error {
	return m.SetId(m.ID)
}

func (m *BaseModel) SetId(id int64) error {
	hashId, err := common.EncodeHashID(m.ID, m.Model, constants.HASHIDS_SALT, constants.HASHIDS_LENGTH)
	if err != nil {
		return errors.NewInvalidDataError(err)
	}

	m.ID = id
	m.HashId = hashId
	return nil
}

func (m *BaseModel) SetHashId(hashId string) error {
	id, err := common.DecodeHashID(hashId, m.Model, constants.HASHIDS_SALT, constants.HASHIDS_LENGTH)
	if err != nil {
		return errors.NewInvalidDataError(err)
	}

	m.ID = id
	m.HashId = hashId
	return nil
}
