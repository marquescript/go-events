package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john.doe@example.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user.ID)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "john.doe@example.com", "123456")
	user.GenerateHash()
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
