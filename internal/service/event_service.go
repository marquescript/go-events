package service

import (
	"errors"
	"time"

	"github.com/marquescript/go-events/internal/entity"
	internalErrors "github.com/marquescript/go-events/internal/errors"
	"github.com/marquescript/go-events/internal/infra/database"
)

type EventService struct {
	EventDB database.EventInterface
	UserDB  database.UserInterface
}

func NewEventService(
	eventDB database.EventInterface,
	userDB database.UserInterface,
) *EventService {
	return &EventService{
		EventDB: eventDB,
		UserDB:  userDB,
	}
}

func (s *EventService) Create(description, address string, date time.Time, userID int64) error {
	p, err := entity.NewEvent(description, address, date, userID)
	if err != nil {
		return err
	}
	err = s.EventDB.Create(p)
	if err != nil {
		return err
	}
	return nil
}

func (s *EventService) FindByID(userID, id int64) (*entity.Event, error) {
	event, err := s.EventDB.FindByID(userID, id)
	if err != nil {
		return nil, internalErrors.NewNotFoundError("Event not found")
	}

	return event, nil
}

func (s *EventService) FindAll(userID int64, page, limit int, sort string) ([]*entity.Event, error) {
	events, err := s.EventDB.FindAll(userID, page, limit, sort)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *EventService) Update(userId, id int64, description, address string, date time.Time) error {
	event, err := s.EventDB.FindByID(userId, id)
	if err != nil {
		return errors.New("event not found")
	}
	if event.UserID != userId {
		return errors.New("this event does not belong to the user")
	}

	s.updateIfNotEmpty(&event.Description, &description)
	s.updateIfNotEmpty(&event.Address, &address)

	if !date.IsZero() {
		event.Date = date
	}

	err = s.EventDB.Update(event)
	if err != nil {
		return err
	}
	return nil
}

func (s *EventService) Delete(userId, id int64) error {
	event, err := s.EventDB.FindByID(userId, id)
	if err != nil {
		return errors.New("event not found")
	}
	if event.UserID != userId {
		return errors.New("this event does not belong to the user")
	}
	err = s.EventDB.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *EventService) updateIfNotEmpty(dst *string, src *string) {
	if *src != "" {
		*dst = *src
	}
}
