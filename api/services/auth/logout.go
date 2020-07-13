package auth

import (
	"context"

	exterrors "github.com/eugeneradionov/ext-errors"
	reqContext "github.com/eugeneradionov/stocks/api/context"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
	"github.com/pkg/errors"
)

func (srv service) Logout(ctx context.Context) exterrors.ExtError {
	var jwtToken = reqContext.GetAccessToken(ctx)

	tx, err := postgres.GetDB().NewTXContext(ctx)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "start tx"))
	}
	defer tx.RollbackTx("logout user")

	claims, err := srv.tokenSrv.Revoke(jwtToken)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "revoke token"))
	}

	lsDAO := srv.loginSessionDAO.WithTx(tx)

	_, err = lsDAO.DisableByTokenID(claims.TokenID)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "disable session"))
	}

	err = tx.Commit()
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "commit tx"))
	}

	return nil
}
