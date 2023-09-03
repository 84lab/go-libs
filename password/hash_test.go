package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HashPassword(t *testing.T) {
	rawPassword := "password"
	wrongRawPassword := "wrong_password"
	hashedPassword, err := HashPassword(rawPassword)
	assert.Nil(t, err)

	isValid := CheckPassword(hashedPassword, wrongRawPassword)
	assert.Equal(t, false, isValid)

	isValid = CheckPassword(hashedPassword, rawPassword)
	assert.Equal(t, true, isValid)
}
