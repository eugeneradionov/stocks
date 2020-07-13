package loginssession

import (
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
	"github.com/google/uuid"
)

type DAO interface {
	WithTx(tx *postgres.DBQuery) DAO

	GetSessionByRefreshToken(refreshToken string, relations ...string) (session models.LoginSession, err error)
	UpdateSession(id, tokenID uuid.UUID, refreshToken string) (err error)
	Insert(ls models.LoginSession) (_ models.LoginSession, err error)
	DisableByTokenID(tokenID uuid.UUID) (sess models.LoginSession, err error)
}
