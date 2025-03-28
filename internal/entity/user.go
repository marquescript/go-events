package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) (*User, error) {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}

func (u *User) GenerateHash() {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(hash)
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
