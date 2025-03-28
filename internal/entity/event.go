package entity

import (
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Address     string    `json:"address"`
	UserID      int64     `json:"userId"`
}

func NewEvent(description, address string, date time.Time, userID int64) (*Event, error) {
	return &Event{
		Description: description,
		Address:     address,
		Date:        date,
		UserID:      userID,
	}, nil
}
