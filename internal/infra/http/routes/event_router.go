package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marquescript/go-events/internal/infra/factory"
)

func RegisterEventRoutes(r chi.Router, eventFactory *factory.EventFactory) {
	r.Route("/events", func(r chi.Router) {
		r.Post("/", eventFactory.Handler.CreateEvent)
		r.Get("/{id}", eventFactory.Handler.FindEvent)
		r.Get("/", eventFactory.Handler.FindAllEvents)
		r.Put("/{eventId}", eventFactory.Handler.UpdateEvent)
		r.Delete("/{eventId}", eventFactory.Handler.DeleteEvent)
	})
}
