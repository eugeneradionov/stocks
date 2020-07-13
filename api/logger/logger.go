package logger

import (
	"context"
	"net/http"

	exterrors "github.com/eugeneradionov/ext-errors"
	lgr "github.com/eugeneradionov/logger"
	"github.com/eugeneradionov/stocks/api/config"
	reqContext "github.com/eugeneradionov/stocks/api/context"
	"go.uber.org/zap"
)

const requestIDKey = "request_id"

var logger *zap.Logger

func Get() *zap.Logger {
	return logger
}

func Load(cfg *config.Config) (err error) {
	logger, err = lgr.Load(lgr.LogPreset(cfg.LogPreset))
	return err
}

func WithCtxValue(ctx context.Context) *zap.Logger {
	return logger.With(zapFieldsFromContext(ctx)...)
}

func LogExtErr(ctx context.Context, extErr exterrors.ExtError, msg string, fields ...zap.Field) {
	fields = append(fields, zap.Error(extErr))

	if extErr.HTTPCode() >= http.StatusInternalServerError {
		WithCtxValue(ctx).Error(msg, fields...)
	} else {
		WithCtxValue(ctx).Info(msg, fields...)
	}
}

func zapFieldsFromContext(ctx context.Context) []zap.Field {
	return []zap.Field{
		zap.String(requestIDKey, reqContext.GetRequestID(ctx)),
	}
}
