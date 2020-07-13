package users

import (
	"context"

	"github.com/eugeneradionov/stocks/api/models"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
)

type usersDAO struct {
	q *postgres.DBQuery
}

func NewUsersDAO() DAO {
	return &usersDAO{q: postgres.GetDB().QueryContext(context.Background())}
}

func (dao usersDAO) WithTx(tx *postgres.DBQuery) DAO {
	return &usersDAO{q: tx}
}

func (dao usersDAO) GetByEmail(email string) (_ *models.User, err error) {
	var user = &models.User{}

	err = dao.q.Model(user).
		Where("email = ?", email).
		First()

	return user, err
}

func (dao usersDAO) Insert(user models.User) (_ models.User, err error) {
	_, err = dao.q.Model(&user).
		Returning("*").
		Insert()

	return user, err
}
