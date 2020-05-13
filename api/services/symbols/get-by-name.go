package symbols

import (
	"fmt"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/go-pg/pg/v9"
	"github.com/pkg/errors"
)

func (srv service) GetByName(name string) (symbol models.Symbol, extErr exterrors.ExtError) {
	symbol, err := srv.symbolsDAO.GetByName(name)
	if err != nil {
		if err == pg.ErrNoRows {
			return symbol, exterrors.NewNotFoundError(
				fmt.Errorf("symbol with name: '%s' not found", name), "symbol_name",
			)
		}

		return symbol, exterrors.NewInternalServerErrorError(
			errors.Wrapf(err, "get symbol with name: '%s'", name),
		)
	}

	return symbol, nil
}
