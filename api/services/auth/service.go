package auth

import (
	"github.com/eugeneradionov/stocks/api/services/password"
	"github.com/eugeneradionov/stocks/api/services/token"
	logins_session "github.com/eugeneradionov/stocks/api/store/repo/postgres/logins-session"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres/users"
)

type service struct {
	tokenSrv    token.Service
	passwordSrv password.Service

	loginSessionDAO logins_session.DAO
	usersDAO        users.DAO
}

func New(
	tokenSrv token.Service,
	passwordSrv password.Service,
	usersDAO users.DAO,
	loginSessionDAO logins_session.DAO,
) Service {
	return &service{
		tokenSrv:        tokenSrv,
		passwordSrv:     passwordSrv,
		usersDAO:        usersDAO,
		loginSessionDAO: loginSessionDAO,
	}
}
