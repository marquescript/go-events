package service

import (
	"time"

	"github.com/marquescript/go-events/internal/entity"
)

type EventServiceInterface interface {
	Create(description, address string, date time.Time, userID int64) error
	FindByID(userID, id int64) (*entity.Event, error)
	FindAll(userID int64, page, limit int, sort string) ([]*entity.Event, error)
	Update(userId, id int64, description, address string, date time.Time) error
	Delete(userId, id int64) error
}

type UserServiceInterface interface {
	Create(name, email, password string) error
	FindByID(id int64) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
