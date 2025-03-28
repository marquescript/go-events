package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEvent(t *testing.T) {
	event, err := NewEvent("Casamento final do mes", "Não definido", time.Now(), 1)
	assert.Nil(t, err)
	assert.NotNil(t, event.ID)
	assert.Equal(t, "Casamento final do mes", event.Description)
	assert.Equal(t, "Não definido", event.Address)
}
