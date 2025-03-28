package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marquescript/go-events/internal/infra/factory"
)

func RegisterUserRoutes(r chi.Router, userFactory *factory.UserFactory) {
	r.Post("/users", userFactory.Handler.CreateUser)
	r.Post("/sign-in", userFactory.Handler.GetJWT)
}
