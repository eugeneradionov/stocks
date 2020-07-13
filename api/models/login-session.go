package models

import (
	"time"

	"github.com/google/uuid"
)

type LoginSession struct {
	ID uuid.UUID `json:"id"`

	TokenID uuid.UUID `json:"-"`
	UserID  uuid.UUID `json:"-"`

	RefreshToken string `json:"refresh_token"`
	Active       bool   `json:"active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
