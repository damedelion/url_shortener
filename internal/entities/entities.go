package entities

import "github.com/google/uuid"

type ShortURL struct {
	ID  uuid.UUID `json:"id"`
	URL string    `json:"short_url"`
}
