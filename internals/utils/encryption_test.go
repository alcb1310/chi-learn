package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"chi-learn/internals/utils"
)

func TestEncryptPasssword(t *testing.T) {
	p := "test"
	_, err := utils.EncryptPasssword(p)
	assert.Nil(t, err)
}

func TestComparePassword(t *testing.T) {
	p := "test"
	hash, err := utils.EncryptPasssword(p)
	assert.Nil(t, err)

	ok, err := utils.ComparePassword(string(hash), p)
	assert.Nil(t, err)
	assert.True(t, ok)
}
