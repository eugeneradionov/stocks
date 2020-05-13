package context

import (
	"context"

	"github.com/eugeneradionov/stocks/api/logger"
)

type contextKey string

const (
	requestIDKey contextKey = "request_id"
)

func GetRequestID(ctx context.Context) (value string) {
	value, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		logger.Get().Error("get request id from context")
		return ""
	}

	return value
}

func WithRequestID(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, requestIDKey, value)
}
