package database

import (
	"fmt"
	"testing"
	"time"

	"github.com/marquescript/go-events/internal/entity"
	"github.com/marquescript/go-events/test/setup"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewEvent(t *testing.T) {
	db := setup.SetupTestDB(t)
	defer db.Close()

	event, err := entity.NewEvent("Evento de teste", "Endereço de teste", time.Date(2025, 3, 20, 10, 0, 0, 0, time.UTC), 1)
	assert.Nil(t, err)

	eventDB := NewEvent(db)
	err = eventDB.Create(event)
	assert.Nil(t, err)

	var eventFound entity.Event
	err = db.QueryRow("SELECT * FROM events WHERE id = ?", event.ID).Scan(&eventFound.ID, &eventFound.Description, &eventFound.Date, &eventFound.Address, &eventFound.UserID)
	assert.Nil(t, err)
	assert.Equal(t, event.ID, eventFound.ID)
	assert.Equal(t, event.Description, eventFound.Description)
}

func TestFindAllEvents(t *testing.T) {
	db := setup.SetupTestDB(t)
	defer db.Close()

	userDB := NewUser(db)
	user, err := entity.NewUser("Teste", "teste@teste.com", "123456")
	assert.Nil(t, err)

	err = userDB.Create(user)
	assert.Nil(t, err)

	eventDB := NewEvent(db)
	for i := 1; i <= 20; i++ {
		event, err := entity.NewEvent(fmt.Sprintf("Evento de teste %d", i), fmt.Sprintf("Endereço de teste %d", i), time.Date(2025, 3, 20, 10, 0, 0, 0, time.UTC), 1)
		assert.Nil(t, err)

		err = eventDB.Create(event)
		assert.Nil(t, err)
	}

	events, err := eventDB.FindAll(user.ID, 1, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 10, len(events))
	assert.Equal(t, "Evento de teste 1", events[0].Description)
}

func TestFindEventByID(t *testing.T) {
	db := setup.SetupTestDB(t)
	defer db.Close()

	eventDB := NewEvent(db)
	event, err := entity.NewEvent("Evento de teste", "Endereço de teste", time.Date(2025, 3, 20, 10, 0, 0, 0, time.UTC), 1)
	assert.Nil(t, err)

	err = eventDB.Create(event)
	assert.Nil(t, err)

	eventFound, err := eventDB.FindByID(1, event.ID)
	assert.Nil(t, err)
	assert.Equal(t, event.ID, eventFound.ID)
	assert.Equal(t, event.Description, eventFound.Description)
}

func TestUpdateEvent(t *testing.T) {
	db := setup.SetupTestDB(t)
	defer db.Close()

	eventDB := NewEvent(db)
	event, err := entity.NewEvent("Evento de teste", "Endereço de teste", time.Date(2025, 3, 20, 10, 0, 0, 0, time.UTC), 1)
	assert.Nil(t, err)

	err = eventDB.Create(event)
	assert.Nil(t, err)

	event.Description = "Evento de teste atualizado"
	event.Address = "Endereço de teste atualizado"

	err = eventDB.Update(event)
	assert.Nil(t, err)

	var eventFound entity.Event
	err = db.QueryRow("SELECT * FROM events WHERE id = ?", event.ID).Scan(&eventFound.ID, &eventFound.Description, &eventFound.Date, &eventFound.Address, &eventFound.UserID)
	assert.Nil(t, err)
	assert.Equal(t, event.ID, eventFound.ID)
	assert.Equal(t, event.Description, eventFound.Description)
	assert.Equal(t, event.Address, eventFound.Address)
}

func TestDeleteEvent(t *testing.T) {
	db := setup.SetupTestDB(t)
	defer db.Close()

	eventDB := NewEvent(db)
	event, err := entity.NewEvent("Evento de teste", "Endereço de teste", time.Date(2025, 3, 20, 10, 0, 0, 0, time.UTC), 1)
	assert.Nil(t, err)

	err = eventDB.Create(event)
	assert.Nil(t, err)

	err = eventDB.Delete(event.ID)
	assert.Nil(t, err)

	_, err = eventDB.FindByID(1, event.ID)
	assert.Error(t, err)
}
