package models

import (
	"time"

	"github.com/google/uuid"
)

type Symbol struct {
	ID uuid.UUID `json:"id"`

	Description   string `json:"description"`
	DisplaySymbol string `json:"displaySymbol"`
	Symbol        string `json:"symbol"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
