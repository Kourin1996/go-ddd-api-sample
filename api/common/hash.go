package common

import (
	"fmt"

	"github.com/speps/go-hashids"
)

func newHashIDObject(key, salt string, length int) (*hashids.HashID, error) {
	hd := hashids.NewData()
	hd.Salt = key + salt
	hd.MinLength = length
	return hashids.NewWithData(hd)
}

func EncodeHashID(id int, key string, salt string, length int) (string, error) {
	h, err := newHashIDObject(key, salt, length)
	if err != nil {
		return "", err
	}
	return h.Encode([]int{id})
}

func DecodeHashID(hashId string, key string, salt string, length int) (int, error) {
	h, err := newHashIDObject(key, salt, length)
	if err != nil {
		return 0, err
	}
	res := h.Decode(hashId)
	if len(res) == 0 {
		return 0, fmt.Errorf("Failed to decode hash id")
	}
	return res[0], nil
}