package helper

import (
	"encoding/json"
	"testing"
)

func MustMarshalJson(t *testing.T, v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}
