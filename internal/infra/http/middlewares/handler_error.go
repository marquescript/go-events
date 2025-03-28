package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/marquescript/go-events/internal/errors"
)

func HandlerError(w http.ResponseWriter, err error) {
	var statusCode int
	var errorMessage string

	switch e := err.(type) {
	case *errors.NotFoundError:
		statusCode = http.StatusNotFound
		errorMessage = e.Message

	default:
		statusCode = http.StatusInternalServerError
		errorMessage = "Internal server error"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"message": errorMessage,
	})
}
