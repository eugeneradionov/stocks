package loginssession

import (
	"context"

	"github.com/eugeneradionov/stocks/api/models"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
	query_utils "github.com/eugeneradionov/stocks/api/store/repo/postgres/query-utils"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type loginSessionDAO struct {
	q *postgres.DBQuery
}

func NewLoginSessionDAO() DAO {
	return &loginSessionDAO{q: postgres.GetDB().QueryContext(context.Background())}
}

func (dao loginSessionDAO) WithTx(tx *postgres.DBQuery) DAO {
	return &loginSessionDAO{q: tx}
}

func (dao loginSessionDAO) GetSessionByRefreshToken(
	refreshToken string,
	relations ...string,
) (session models.LoginSession, err error) {
	q := dao.q.Model(&session).
		Where(`refresh_token = ?`, refreshToken).
		Column("login_session.*")
	q = query_utils.WithRelations(q, relations...)

	err = q.First()

	return session, err
}

func (dao loginSessionDAO) UpdateSession(id, tokenID uuid.UUID, refreshToken string) (err error) {
	_, err = dao.q.Model(&models.LoginSession{}).
		Where(`"login_session"."id" = ?`, id).
		Set(`token_id = ?`, tokenID).
		Set(`refresh_token = ?`, refreshToken).
		Returning(`*`).
		Update()

	return errors.Wrap(err, "update session")
}

func (dao loginSessionDAO) Insert(ls models.LoginSession) (_ models.LoginSession, err error) {
	_, err = dao.q.Model(&ls).
		Returning("*").
		Insert()

	return ls, err
}

func (dao loginSessionDAO) DisableByTokenID(tokenID uuid.UUID) (sess models.LoginSession, err error) {
	_, err = dao.q.Model(&sess).
		Where(`"login_session"."token_id" = ?`, tokenID).
		Set("active = false").
		Returning("*").
		Update()

	return sess, err
}
