package logger

import (
	lgr "github.com/eugeneradionov/logger"
	"github.com/eugeneradionov/stocks/api/config"
	"go.uber.org/zap"
)

var logger *zap.Logger

func Get() *zap.Logger {
	return logger
}

func Load(cfg *config.Config) (err error) {
	logger, err = lgr.Load(lgr.LogPreset(cfg.LogPreset))
	return err
}
