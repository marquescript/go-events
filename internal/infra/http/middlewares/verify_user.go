package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/jwtauth"
	"github.com/marquescript/go-events/internal/service"
)

type UserContext struct {
	ID   int64
	Name string
}

type key string

const UserContextKey key = "user"

func VerifyUserMiddleware(userService service.UserServiceInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			if claims == nil {
				http.Error(w, "Token inválido", http.StatusUnauthorized)
				return
			}
			userID, ok := claims["sub"].(string)
			if !ok || userID == "" {
				http.Error(w, "Token inválido", http.StatusUnauthorized)
				return
			}

			userIDInt, err := strconv.ParseInt(userID, 10, 64)
			if err != nil {
				http.Error(w, "ID inválido", http.StatusUnauthorized)
				return
			}

			user, err := userService.FindByID(userIDInt)
			if err != nil || user == nil {
				http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
				return
			}

			userCtx := UserContext{
				ID:   user.ID,
				Name: user.Name,
			}

			ctx := context.WithValue(r.Context(), UserContextKey, userCtx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
