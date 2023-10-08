package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassword2Hash(t *testing.T) {
	hashedPwd, _ := Password2Hash("password")
	validateResult := ValidatePasswordAndHash("password", hashedPwd)
	assert.Equal(t, validateResult, true)
}
