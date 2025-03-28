package database

import "github.com/marquescript/go-events/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id int64) (*entity.User, error)
}

type EventInterface interface {
	Create(event *entity.Event) error
	FindByID(userID, id int64) (*entity.Event, error)
	FindAll(userId int64, page, limit int, sort string) ([]*entity.Event, error)
	Update(event *entity.Event) error
	Delete(id int64) error
}
