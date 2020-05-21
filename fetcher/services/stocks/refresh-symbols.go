package stocks

import (
	"time"

	"github.com/eugeneradionov/stocks/fetcher/config"
	"github.com/eugeneradionov/stocks/fetcher/logger"
	"go.uber.org/zap"
)

// RefreshSymbols spawns goroutine and gets symbols for exchange code with specified timeout
func (srv service) RefreshSymbols(cfg config.Symbols, exchangeCode string) {
	go func() {
		extErr := srv.GetSymbols(exchangeCode)
		if extErr != nil {
			logger.Get().Error("failed to refresh symbols", zap.Error(extErr))
		}

		time.Sleep(time.Duration(cfg.RefreshFreq))
	}()
}
