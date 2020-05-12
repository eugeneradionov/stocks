package stocks

import (
	"sync"

	"github.com/eugeneradionov/stocks/api/services/symbols"
)

type service struct {
	symbolsSrv symbols.Service
}

var (
	srv  service
	once = &sync.Once{}
)

func New(symbolsSrv symbols.Service) (_ Service, err error) {
	once.Do(func() {
		srv = service{
			symbolsSrv: symbolsSrv,
		}
	})

	return srv, err
}
