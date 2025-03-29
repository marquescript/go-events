package service

import (
	"github.com/marquescript/go-events/internal/entity"
	internalErrors "github.com/marquescript/go-events/internal/errors"
	"github.com/marquescript/go-events/internal/infra/database"
)

type UserService struct {
	UserDB database.UserInterface
}

func NewUserService(userDB database.UserInterface) *UserService {
	return &UserService{
		UserDB: userDB,
	}
}

func (s *UserService) Create(name, email, password string) error {
	user, err := entity.NewUser(name, email, password)
	if err != nil {
		return err
	}

	user.GenerateHash()

	err = s.UserDB.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) FindByID(id int64) (*entity.User, error) {
	user, err := s.UserDB.FindByID(id)
	if err != nil {
		return nil, internalErrors.NewNotFoundError("User not found")
	}
	return user, nil
}

func (s *UserService) FindByEmail(email string) (*entity.User, error) {
	user, err := s.UserDB.FindByEmail(email)
	if err != nil {
		return nil, internalErrors.NewNotFoundError("User not found")
	}
	return user, nil
}
