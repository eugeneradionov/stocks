package token

import (
	"context"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrRefreshNotFound  = errors.New("active refresh token not found")
	ErrForbiddenRefresh = errors.New("refresh is forbidden")
)

func (srv service) Refresh(ctx context.Context, accessToken, refreshToken string) (*models.Token, exterrors.ExtError) { // nolint
	access, err := srv.parseJWT(accessToken)
	if err != nil {
		return nil, exterrors.NewUnprocessableEntityError(ErrTokenInvalid, "access_token")
	}

	accessClaims := srv.parseClaims(access)
	if accessClaims == nil {
		return nil, exterrors.NewUnprocessableEntityError(ErrTokenClaimsInvalid, "access_token")
	}

	if access.Claims.Valid() == nil {
		err = srv.checkBlackList(accessClaims.TokenID)
		if err != nil {
			return nil, exterrors.NewUnprocessableEntityError(ErrForbiddenRefresh, "refresh_token")
		}
	}

	tx, err := postgres.GetDB().NewTXContext(ctx)
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "start db transaction"))
	}
	defer tx.RollbackTx("refresh JWT token")

	lsDAO := srv.loginSessionDAO.WithTx(tx)

	session, err := lsDAO.GetSessionByRefreshToken(refreshToken)
	if err != nil {
		return nil, exterrors.NewUnprocessableEntityError(ErrRefreshNotFound, "refresh_token")
	}

	if !session.Active || session.TokenID != accessClaims.TokenID {
		return nil, exterrors.NewUnprocessableEntityError(ErrTokenNotFound, "refresh_token")
	}

	claims := Claims{
		TokenID:   uuid.New(),
		SessionID: session.ID,
		UserID:    session.UserID,
	}

	accessNew, err := srv.Generate(&claims)
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "generate token"))
	}

	refreshNew, err := srv.GenerateRefresh()
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "generate refresh token"))
	}

	err = lsDAO.UpdateSession(session.ID, claims.TokenID, refreshNew)
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "update session token"))
	}

	_, err = srv.Revoke(accessToken)
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "revoke access token"))
	}

	err = tx.Commit()
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "commit tx"))
	}

	return &models.Token{Access: accessNew, Refresh: refreshNew}, nil
}
