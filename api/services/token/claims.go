package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Claims struct {
	SessionID uuid.UUID
	TokenID   uuid.UUID
	UserID    uuid.UUID

	jwt.StandardClaims
}

func (c *Claims) TTL() int64 {
	return c.StandardClaims.ExpiresAt - time.Now().Unix()
}
