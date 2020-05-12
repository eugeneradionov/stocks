package symbols

import exterrors "github.com/eugeneradionov/ext-errors"

type Service interface {
	Consume() exterrors.ExtError
}
