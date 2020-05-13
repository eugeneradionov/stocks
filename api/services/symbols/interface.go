package symbols

import (
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/api/models"
)

type Service interface {
	Get(limit, offset int) (symbols []models.Symbol, extErr exterrors.ExtError)
	GetByName(name string) (symbol models.Symbol, extErr exterrors.ExtError)

	Consume() exterrors.ExtError
}
