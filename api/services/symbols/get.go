package symbols

import (
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/pkg/errors"
)

func (srv service) Get(limit, offset int) (symbols []models.Symbol, extErr exterrors.ExtError) {
	symbols, err := srv.symbolsDAO.Get(limit, offset)
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "get symbols"))
	}

	return symbols, nil
}
