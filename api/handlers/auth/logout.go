package auth

import (
	"net/http"

	"github.com/eugeneradionov/stocks/api/handlers/common"
	"github.com/eugeneradionov/stocks/api/logger"
	"github.com/eugeneradionov/stocks/api/services"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()

	extErr := services.Get().Auth().Logout(ctx)
	if extErr != nil {
		logger.LogExtErr(ctx, extErr, "failed to logout user")
		common.SendExtError(w, extErr)

		return
	}

	common.SendResponse(w, http.StatusNoContent, nil)
}
