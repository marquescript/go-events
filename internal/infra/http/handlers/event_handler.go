package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/marquescript/go-events/internal/dto"
	"github.com/marquescript/go-events/internal/infra/http/middlewares"
	"github.com/marquescript/go-events/internal/service"
	pkg "github.com/marquescript/go-events/pkg/utils"
)

type EventHandler struct {
	EventService service.EventServiceInterface
}

func NewEventHandler(eventService service.EventServiceInterface) *EventHandler {
	return &EventHandler{
		EventService: eventService,
	}
}

// CreateEvent godoc
//
//	@Summary		CreateEvent
//	@Description	Create a new event
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.EventDTO	true	"event request"
//	@Success		201
//	@Failure		400	{string}	string	"Invalid request body or parameters"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/events [post]
//
//	@Security		ApiKeyAuth
func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event dto.EventDTO
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := pkg.ParseDate(event.Date)
	if err != nil {
		http.Error(w, "formato de data inválido. Use o formato YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	userID, err := pkg.ParseID(event.UserID)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	err = h.EventService.Create(
		event.Description,
		event.Address,
		date,
		userID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// FindEvent godoc
//
//	@Summary		FindEvent
//	@Description	Find an event by ID
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Event ID"
//	@Success		200	{object}	dto.EventDTO
//	@Failure		400	{string}	string	"Invalid ID"
//	@Failure		401	{string}	string	"Unauthorized"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/events/{id} [get]
//
//	@Security		ApiKeyAuth
func (h *EventHandler) FindEvent(w http.ResponseWriter, r *http.Request) {
	userCtx, ok := r.Context().Value(middlewares.UserContextKey).(middlewares.UserContext)
	if !ok {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}

	id := chi.URLParam(r, "id")
	eventID, err := pkg.ParseID(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	event, err := h.EventService.FindByID(userCtx.ID, eventID)
	if err != nil {
		middlewares.HandlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

// CreateEvent godoc
//
//	@Summary		FindAllEvents
//	@Description	Find all events
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			page	query		string	false	"page number"
//	@Param			limit	query		string	false	"limit"
//	@Success		201		{array}		entity.Event
//	@Failure		400		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/events [get]
//
//	@Security		ApiKeyAuth
func (h *EventHandler) FindAllEvents(w http.ResponseWriter, r *http.Request) {
	userCtx, ok := r.Context().Value(middlewares.UserContextKey).(middlewares.UserContext)
	if !ok {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "asc"
	}

	events, err := h.EventService.FindAll(userCtx.ID, page, limit, sort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

// UpdateEvent godoc
//
//	@Summary		UpdateEvent
//	@Description	Update an existing event
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			eventId	path	string				true	"Event ID"
//	@Param			request	body	dto.EventUpdateDTO	true	"Event update request"
//	@Success		204
//	@Failure		400	{string}	string	"Invalid request body or parameters"
//	@Failure		401	{string}	string	"Unauthorized"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/events/{eventId} [put]
//
//	@Security		ApiKeyAuth
func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	userCtx, ok := r.Context().Value(middlewares.UserContextKey).(middlewares.UserContext)
	if !ok {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}
	eventId := chi.URLParam(r, "eventId")

	eventIdParsed, err := pkg.ParseID(eventId)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var event dto.EventUpdateDTO
	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := pkg.ParseDate(event.Date)
	if err != nil {
		http.Error(w, "formato de data inválido. Use o formato YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	err = h.EventService.Update(userCtx.ID, eventIdParsed, event.Description, event.Address, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteEvent godoc
//
//	@Summary		DeleteEvent
//	@Description	Delete an event by ID
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			eventId	path	string	true	"Event ID"
//	@Success		204
//	@Failure		400	{string}	string	"Invalid ID"
//	@Failure		401	{string}	string	"Unauthorized"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/events/{eventId} [delete]
//
//	@Security		ApiKeyAuth
func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	userCtx, ok := r.Context().Value(middlewares.UserContextKey).(middlewares.UserContext)
	if !ok {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}
	eventId := chi.URLParam(r, "eventId")

	eventIdParsed, err := pkg.ParseID(eventId)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.EventService.Delete(userCtx.ID, eventIdParsed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
