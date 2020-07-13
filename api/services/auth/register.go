package auth

import (
	"context"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/eugeneradionov/stocks/api/models/auth"
	"github.com/eugeneradionov/stocks/api/services/token"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
	logins_session "github.com/eugeneradionov/stocks/api/store/repo/postgres/logins-session"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres/users"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (srv service) Register(
	ctx context.Context,
	req auth.RegistrationReq,
) (resp auth.RegistrationResp, extErr exterrors.ExtError) {
	tx, err := postgres.GetDB().NewTXContext(ctx)
	if err != nil {
		return resp, exterrors.NewInternalServerErrorError(errors.Wrap(err, "start tx"))
	}
	defer tx.RollbackTx("register user")

	usersDAO := srv.usersDAO.WithTx(tx)

	existingUser, err := usersDAO.GetByEmail(req.Email)
	if err != nil {
		if err != pg.ErrNoRows { // internal error
			return resp, exterrors.NewInternalServerErrorError(errors.Wrap(err, "get user by email"))
		}
	}

	if existingUser != nil {
		return resp, exterrors.NewUnprocessableEntityError(ErrEmailExists, "email")
	}

	user, extErr := srv.createUser(usersDAO, req)
	if extErr != nil {
		return resp, extErr
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

	resp = auth.RegistrationResp{
		User:  user,
		Token: tkn,
	}

	return resp, nil
}

func (srv service) createUser(usersDAO users.DAO, req auth.RegistrationReq) (user models.User, _ exterrors.ExtError) {
	salt, err := srv.passwordSrv.GenerateSalt()
	if err != nil {
		return user, exterrors.NewInternalServerErrorError(errors.Wrap(err, "generate password salt"))
	}

	encrPassword, err := srv.passwordSrv.EncryptPassword(req.Password, salt)
	if err != nil {
		return user, exterrors.NewInternalServerErrorError(errors.Wrap(err, "encrypt password"))
	}

	user = models.User{
		ID:       uuid.New(),
		Email:    req.Email,
		Password: encrPassword,
		Salt:     salt,
		Name:     req.Name,
		Status:   models.ActiveUser,
	}

	user, err = usersDAO.Insert(user)
	if err != nil {
		return user, exterrors.NewInternalServerErrorError(errors.Wrap(err, "insert user"))
	}

	return user, nil
}

func (srv service) createSession(
	lsDAO logins_session.DAO,
	userID uuid.UUID,
) (sess models.LoginSession, _ exterrors.ExtError) {
	refreshToken, err := srv.tokenSrv.GenerateRefresh()
	if err != nil {
		return sess, exterrors.NewInternalServerErrorError(errors.Wrap(err, "generate refresh token"))
	}

	sess = models.LoginSession{
		ID:           uuid.New(),
		TokenID:      uuid.New(),
		UserID:       userID,
		RefreshToken: refreshToken,
		Active:       true,
	}

	sess, err = lsDAO.Insert(sess)
	if err != nil {
		return sess, exterrors.NewInternalServerErrorError(errors.Wrap(err, "insert login session"))
	}

	return sess, nil
}

func (srv service) createToken(sess models.LoginSession) (tkn models.Token, extError exterrors.ExtError) {
	claims := &token.Claims{
		SessionID: sess.ID,
		TokenID:   sess.TokenID,
		UserID:    sess.UserID,
	}

	jwtToken, err := srv.tokenSrv.Generate(claims)
	if err != nil {
		return tkn, exterrors.NewInternalServerErrorError(errors.Wrap(err, "generate JWT token"))
	}

	tkn = models.Token{
		Access:  jwtToken,
		Refresh: sess.RefreshToken,
	}

	return tkn, nil
}
