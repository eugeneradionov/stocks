package token

import (
	"context"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/google/uuid"
)

type Service interface {
	Generate(*Claims) (string, error)
	GenerateRefresh() (string, error)
	Validate(token string) (*models.LoginSession, error)
	Revoke(token string) (*Claims, error)
	Refresh(ctx context.Context, token, refreshToken string) (*models.Token, exterrors.ExtError)
	AddToBlacklist(tokenID uuid.UUID, ttl int64) error
	RemoveFromBlacklist(tokenID uuid.UUID) error
	MaxExpirationTime() int64
}
