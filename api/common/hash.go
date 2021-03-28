package common

import (
	"github.com/speps/go-hashids"
)

func newHashIDObject(key, salt string, length int) (*hashids.HashID, error) {
	hd := hashids.NewData()
	hd.Salt = key + salt
	hd.MinLength = length
	return hashids.NewWithData(hd)
}

func EncodeHashID(id int64, key string, salt string, length int) (string, error) {
	h, err := newHashIDObject(key, salt, length)
	if err != nil {
		return "", err
	}
	return h.Encode([]int{int(id)})
}

func DecodeHashID(hashId string, key string, salt string, length int) (int64, error) {
	h, err := newHashIDObject(key, salt, length)
	if err != nil {
		return 0, err
	}
	res := h.Decode(hashId)
	if len(res) == 0 {
		return 0, err
	}
	return int64(res[0]), nil
}
