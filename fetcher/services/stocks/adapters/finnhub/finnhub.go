package finnhub

import (
	"net/http"

	"github.com/eugeneradionov/stocks/fetcher/config"
)

// https://finnhub.io/api/v1/stock/symbol?exchange=US&token=bqqhj5nrh5rcj5177tag
type Finnhub struct {
	cfg        config.Finnhub
	httpClient *http.Client
}

func NewFinnhub(fhCfg config.Finnhub, httpClient *http.Client) Finnhub {
	return Finnhub{
		cfg:        fhCfg,
		httpClient: httpClient,
	}
}
