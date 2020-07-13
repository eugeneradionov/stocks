package auth

import (
	"net/http"

	"github.com/eugeneradionov/stocks/api/handlers/common"
	"github.com/eugeneradionov/stocks/api/logger"
	"github.com/eugeneradionov/stocks/api/models/auth"
	"github.com/eugeneradionov/stocks/api/services"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		req auth.LoginReq
	)

	err := common.ProcessRequestBody(w, r, &req)
	if err != nil {
		return
	}

	resp, extErr := services.Get().Auth().Login(ctx, req)
	if extErr != nil {
		logger.LogExtErr(ctx, extErr, "failed to register user")
		common.SendExtError(w, extErr)

		return
	}

	common.SendResponse(w, http.StatusOK, resp)
}
