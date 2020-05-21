package candles

import (
	"net/http"

	"github.com/eugeneradionov/stocks/api/handlers/common"
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/eugeneradionov/stocks/api/services"
	"github.com/go-chi/chi"
)

func Get(w http.ResponseWriter, r *http.Request) {
	var (
		symbolName = chi.URLParam(r, "symbolName")

		req models.CandleRequest
	)

	err := common.ProcessRequestBody(w, r, &req)
	if err != nil {
		return
	}

	symbol, extErr := services.Get().Candles().Get(symbolName, req.Resolution, req.From.Time(), req.To.Time())
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, symbol)
}
