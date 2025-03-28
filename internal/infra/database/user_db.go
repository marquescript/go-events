package database

import (
	"database/sql"

	"github.com/marquescript/go-events/internal/entity"
)

type UserDB struct {
	DB *sql.DB
}

func NewUser(db *sql.DB) *UserDB {
	return &UserDB{
		DB: db,
	}
}

func (u *UserDB) Create(user *entity.User) error {
	_, err := u.DB.Exec("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.DB.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDB) FindByID(id int64) (*entity.User, error) {
	var user entity.User
	err := u.DB.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
