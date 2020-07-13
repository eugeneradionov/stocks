package models

import (
	"time"

	"github.com/google/uuid"
)

type UserStatus string

const (
	ActiveUser   UserStatus = "Active"
	InactiveUser UserStatus = "Inactive"
	BlockedUser  UserStatus = "Blocked"
	DeletedUser  UserStatus = "Deleted"
)

type User struct {
	ID uuid.UUID `json:"id"`

	Email    string `json:"email"`
	Password string `json:"-"`
	Salt     string `json:"-"`
	Name     string `json:"name"`

	Status UserStatus `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
