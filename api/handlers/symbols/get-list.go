package symbols

import (
	"net/http"

	"github.com/eugeneradionov/stocks/api/handlers/common"
	"github.com/eugeneradionov/stocks/api/services"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query()

	limit, offset, extErr := common.GetLimitAndOffset(query)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	symbols, extErr := services.Get().Symbols().Get(limit, offset)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, symbols)
}
