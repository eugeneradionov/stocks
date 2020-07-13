package users

import (
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
)

type DAO interface {
	WithTx(tx *postgres.DBQuery) DAO

	GetByEmail(email string) (*models.User, error)
	Insert(user models.User) (models.User, error)
}
