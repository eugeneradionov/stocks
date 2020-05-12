package stocks

import exterrors "github.com/eugeneradionov/ext-errors"

type Service interface {
	ConsumeAll() exterrors.ExtError
}
