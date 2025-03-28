package factory

import (
	"database/sql"

	"github.com/marquescript/go-events/internal/infra/database"
	"github.com/marquescript/go-events/internal/infra/http/handlers"
	"github.com/marquescript/go-events/internal/service"
)

type UserFactory struct {
	Handler *handlers.UserHandler
}

func NewUserFactory(db *sql.DB) *UserFactory {
	userDB := database.NewUser(db)
	userService := service.NewUserService(userDB)
	userHandler := handlers.NewUserHandler(userService)

	return &UserFactory{
		Handler: userHandler,
	}
}
