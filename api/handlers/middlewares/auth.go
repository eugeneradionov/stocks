package middlewares

import (
	"net/http"

	exterrors "github.com/eugeneradionov/ext-errors"
	reqContext "github.com/eugeneradionov/stocks/api/context"
	"github.com/eugeneradionov/stocks/api/handlers/common"
	"github.com/eugeneradionov/stocks/api/logger"
	"github.com/eugeneradionov/stocks/api/models/auth"
	"github.com/eugeneradionov/stocks/api/services"
	"go.uber.org/zap"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var accessToken = r.Header.Get(auth.AccessTokenHeader)

			loginSes, err := services.Get().Token().Validate(accessToken)
			if err != nil {
				logger.Get().Info("Token is invalid", zap.String("token", accessToken))
				common.SendExtError(w, exterrors.NewUnauthorizedError(err))

				return
			}

			ctx := reqContext.WithUserID(r.Context(), loginSes.UserID)
			ctx = reqContext.WithAccessTokenID(ctx, loginSes.TokenID)
			ctx = reqContext.WithAccessToken(ctx, accessToken)

			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
