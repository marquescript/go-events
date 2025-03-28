package database

import (
	"testing"

	"github.com/marquescript/go-events/internal/entity"
	"github.com/marquescript/go-events/test/setup"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestCreateUser(t *testing.T) {
	db := setup.SetupTestDB(t)
	defer db.Close()

	user, _ := entity.NewUser("John Doe", "john.doe@example.com", "password123")
	userDB := NewUser(db)

	err := userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", user.ID).Scan(&userFound.ID, &userFound.Name, &userFound.Email, &userFound.Password)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
}

func TestFindByEmail(t *testing.T) {
	db := setup.SetupTestDB(t)
	defer db.Close()

	user, _ := entity.NewUser("John Doe", "john.doe@example.com", "password123")
	userDB := NewUser(db)

	err := userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
}
