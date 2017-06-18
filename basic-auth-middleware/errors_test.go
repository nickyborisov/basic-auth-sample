package basic_auth

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const (
	TestUserName = "test"
	TestPassword = "test"
)

func TestInvalidUserError(t *testing.T) {

	err := NewInvalidUserError(TestUserName)

	assert.IsType(t, InvalidUserError{}, err)
	assert.Equal(t, err.(InvalidUserError).User, TestUserName, "they should be equal")

}

func TestInvalidPasswordError(t *testing.T) {

	err := NewInvalidPasswordError(TestPassword)

	assert.IsType(t, InvalidPasswordError{}, err)
	assert.Equal(t, err.(InvalidPasswordError).Password, TestPassword, "they should be equal")

}