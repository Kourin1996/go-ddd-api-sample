package common

import (
	"github.com/Kourin1996/go-crud-api-sample/api/common"
	"github.com/Kourin1996/go-crud-api-sample/api/constants"
)

type ID struct {
	Model string
	ID    int
}

func NewIDFromID(model string, id int) ID {
	return ID{Model: model, ID: id}
}

func NewIDFromHashId(model string, hashId string) (ID, error) {
	id, err := common.DecodeHashID(hashId, model, constants.HASHIDS_SALT, constants.HASHIDS_LENGTH)
	if err != nil {
		return ID{}, err
	}
	return ID{Model: model, ID: id}, nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	hashId, err := common.EncodeHashID(id.ID, id.Model, constants.HASHIDS_SALT, constants.HASHIDS_LENGTH)
	if err != nil {
		return nil, err
	}
	return []byte(`"` + hashId + `"`), nil
}
