package helper

import (
	"github.com/Kourin1996/go-crud-api-sample/api/common"
	"github.com/Kourin1996/go-crud-api-sample/api/constants"
)

func EncodeId(model string, id int64) string {
	hashId, _ := common.EncodeHashID(id, model, constants.HASHIDS_SALT, constants.HASHIDS_LENGTH)
	return hashId
}

func DecodeHashId(model string, hashId string) int64 {
	id, _ := common.DecodeHashID(hashId, model, constants.HASHIDS_SALT, constants.HASHIDS_LENGTH)
	return id
}
