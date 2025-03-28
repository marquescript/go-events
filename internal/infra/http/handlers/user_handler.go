package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/marquescript/go-events/internal/dto"
	"github.com/marquescript/go-events/internal/infra/http/middlewares"
	"github.com/marquescript/go-events/internal/service"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// CreateUser godoc
//
//	@Summary		CreateUser
//	@Description	Create user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.UserDTO	true	"user request"
//	@Success		201
//	@Failure		500	{object}	Error
//	@Router			/users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		middlewares.HandlerError(w, err)
		return
	}

	err = h.UserService.Create(user.Name, user.Email, user.Password)
	if err != nil {
		middlewares.HandlerError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetJWT godoc
//
//	@Summary		GetJWT
//	@Description	Get jwt
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.UserMinDTO	true	"user credentials"
//	@Success		200		{object}	dto.JWTOutput
//	@Failure		404
//	@Failure		500
//	@Router			/sign-in [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExpiresIn").(int)

	var user dto.UserMinDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := h.UserService.FindByEmail(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub":    strconv.FormatInt(u.ID, 10),
		"expire": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken := dto.JWTOutput{
		AccessToken: tokenString,
		Payload: dto.JWTPayload{
			Sub: strconv.FormatInt(u.ID, 10),
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
