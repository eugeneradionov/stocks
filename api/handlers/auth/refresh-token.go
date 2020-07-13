package auth

import (
	"net/http"

	reqContext "github.com/eugeneradionov/stocks/api/context"
	"github.com/eugeneradionov/stocks/api/handlers/common"
	"github.com/eugeneradionov/stocks/api/logger"
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/eugeneradionov/stocks/api/services"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	var (
		req models.Token
		ctx = r.Context()
	)

	err := common.ProcessRequestBody(w, r, &req)
	if err != nil {
		return
	}

	token, extErr := services.Get().Token().Refresh(ctx, reqContext.GetAccessToken(ctx), req.Refresh)
	if extErr != nil {
		logger.LogExtErr(ctx, extErr, "refresh token")
		common.SendExtError(w, extErr)

		return
	}

	common.SendResponse(w, http.StatusOK, token)
}
