package auth

import (
	"context"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/api/models/auth"
)

type Service interface {
	Register(ctx context.Context, req auth.RegistrationReq) (auth.RegistrationResp, exterrors.ExtError)
	Login(ctx context.Context, req auth.LoginReq) (auth.LoginResp, exterrors.ExtError)
	Logout(ctx context.Context) exterrors.ExtError
}
