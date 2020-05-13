package symbols

import (
	"net/http"

	"github.com/eugeneradionov/stocks/api/handlers/common"
	"github.com/eugeneradionov/stocks/api/services"
	"github.com/go-chi/chi"
)

func GetByName(w http.ResponseWriter, r *http.Request) {
	var symbolName = chi.URLParam(r, "symbolName")

	symbol, extErr := services.Get().Symbols().GetByName(symbolName)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, symbol)
}
