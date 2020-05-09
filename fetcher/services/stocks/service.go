package stocks

import (
	"errors"
	"net/http"
	"sync"

	"github.com/eugeneradionov/stocks/fetcher/config"
	"github.com/eugeneradionov/stocks/fetcher/services/stocks/adapters"
	"github.com/eugeneradionov/stocks/fetcher/services/stocks/adapters/finnhub"
)

type service struct {
	httpClient *http.Client
	adapter    adapters.Interface
}

var (
	srv  service
	once = &sync.Once{}
)

func New(cfg config.Stocks, httpClient *http.Client) (_ Service, err error) {
	var adapter adapters.Interface

	once.Do(func() {
		adapter, err = NewAdapter(cfg, httpClient)
		if err != nil {
			return
		}

		srv = service{
			httpClient: httpClient,
			adapter:    adapter,
		}
	})

	return srv, err
}

func NewAdapter(cfg config.Stocks, httpClient *http.Client) (adapters.Interface, error) {
	switch cfg.Adapter {
	case "finnhub":
		return finnhub.NewFinnhub(cfg.Finnhub, httpClient), nil
	default:
		return nil, errors.New("invalid stocks adapter")
	}
}
