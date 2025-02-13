package entities

import (
	"time"

	"github.com/google/uuid"
)

type URL struct {
	ID        uuid.UUID
	ShortURL  string
	LongURL   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
