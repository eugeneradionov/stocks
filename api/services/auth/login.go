package auth

import (
	"context"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/api/models/auth"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
	"github.com/go-pg/pg/v9"
	"github.com/pkg/errors"
)

func (srv service) Login(ctx context.Context, req auth.LoginReq) (resp auth.LoginResp, extErr exterrors.ExtError) {
	tx, err := postgres.GetDB().NewTXContext(ctx)
	if err != nil {
		return resp, exterrors.NewInternalServerErrorError(errors.Wrap(err, "start tx"))
	}
	defer tx.RollbackTx("login user")

	usersDAO := srv.usersDAO.WithTx(tx)

	user, err := usersDAO.GetByEmail(req.Email)
	if err != nil {
		if err == pg.ErrNoRows { // internal error
			return resp, exterrors.NewUnprocessableEntityError(ErrInvalidEmailOrPassword, "email")
		}

		return resp, exterrors.NewInternalServerErrorError(errors.Wrap(err, "get user by email"))
	}

	if user == nil {
		return resp, exterrors.NewUnprocessableEntityError(ErrInvalidEmailOrPassword, "email")
	}

	if !srv.passwordSrv.CompareHashAndPassword(req.Password, user.Salt, user.Password) {
		return resp, exterrors.NewUnprocessableEntityError(ErrInvalidEmailOrPassword, "email")
	}

	lsDAO := srv.loginSessionDAO.WithTx(tx)

	sess, extErr := srv.createSession(lsDAO, user.ID)
	if extErr != nil {
		return resp, extErr
	}

	tkn, extErr := srv.createToken(sess)
	if extErr != nil {
		return resp, extErr
	}

	err = tx.Commit()
	if err != nil {
		return resp, exterrors.NewInternalServerErrorError(errors.Wrap(err, "commit tx"))
	}

	resp = auth.LoginResp{
		User:  *user,
		Token: tkn,
	}

	return resp, nil
}
