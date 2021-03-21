package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertError(t *testing.T, success bool, err error) {
	if success {
		assert.NoError(t, err)
	} else {
		assert.Error(t, err)
	}
}
