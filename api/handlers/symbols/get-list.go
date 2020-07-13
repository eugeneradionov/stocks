package symbols

import (
	"net/http"

	"github.com/eugeneradionov/stocks/api/handlers/common"
	"github.com/eugeneradionov/stocks/api/logger"
	"github.com/eugeneradionov/stocks/api/services"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	var (
		query = r.URL.Query()
		ctx   = r.Context()
	)

	limit, offset, extErr := common.GetLimitAndOffset(query)
	if extErr != nil {
		logger.LogExtErr(ctx, extErr, "failed to get limit and offset")
		common.SendExtError(w, extErr)

		return
	}

	symbols, extErr := services.Get().Symbols().Get(limit, offset)
	if extErr != nil {
		logger.LogExtErr(ctx, extErr, "failed to get symbols list")
		common.SendExtError(w, extErr)

		return
	}

	common.SendResponse(w, http.StatusOK, symbols)
}
