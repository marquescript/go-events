package database

import (
	"database/sql"
	"fmt"

	"github.com/marquescript/go-events/internal/entity"
)

type EventDB struct {
	DB *sql.DB
}

func NewEvent(db *sql.DB) *EventDB {
	return &EventDB{
		DB: db,
	}
}

func (e *EventDB) Create(event *entity.Event) error {
	_, err := e.DB.Exec("INSERT INTO events (id, description, date, address, user_id) VALUES (?, ?, ?, ?, ?)", event.ID, event.Description, event.Date, event.Address, event.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (e *EventDB) FindAll(userId int64, page, limit int, sort string) ([]*entity.Event, error) {
	if sort == "" || sort != "desc" && sort != "asc" {
		sort = "asc"
	}
	query := fmt.Sprintf("SELECT * FROM events WHERE user_id = ? ORDER BY date %s LIMIT ? OFFSET ?", sort)
	rows, err := e.DB.Query(query, userId, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []*entity.Event{}
	for rows.Next() {
		var event entity.Event
		err := rows.Scan(&event.ID, &event.Description, &event.Date, &event.Address, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}

func (e *EventDB) FindByID(userID, id int64) (*entity.Event, error) {
	var event entity.Event
	err := e.DB.QueryRow("SELECT * FROM events WHERE id = ? AND user_id = ?", id, userID).Scan(&event.ID, &event.Description, &event.Date, &event.Address, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e *EventDB) Update(event *entity.Event) error {
	_, err := e.DB.Exec("UPDATE events SET description = ?, date = ?, address = ?, user_id = ? WHERE id = ?", event.Description, event.Date, event.Address, event.UserID, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (e *EventDB) Delete(id int64) error {
	_, err := e.DB.Exec("DELETE FROM events WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
