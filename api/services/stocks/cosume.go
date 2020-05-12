package stocks

import exterrors "github.com/eugeneradionov/ext-errors"

func (srv service) ConsumeAll() (extErr exterrors.ExtError) {
	extErr = srv.symbolsSrv.Consume()
	if extErr != nil {
		return extErr
	}

	return nil
}
