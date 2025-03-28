package factory

import (
	"database/sql"

	"github.com/marquescript/go-events/internal/infra/database"
	"github.com/marquescript/go-events/internal/infra/http/handlers"
	"github.com/marquescript/go-events/internal/service"
)

type EventFactory struct {
	Handler *handlers.EventHandler
}

func NewEventFactory(db *sql.DB) *EventFactory {
	eventDB := database.NewEvent(db)
	userDB := database.NewUser(db)
	eventService := service.NewEventService(eventDB, userDB)
	eventHandler := handlers.NewEventHandler(eventService)
	return &EventFactory{
		Handler: eventHandler,
	}
}
