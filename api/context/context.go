package context

import (
	"context"
)

type contextKey string

const (
	requestIDKey contextKey = "request_id"
)

func GetRequestID(ctx context.Context) (value string) {
	value, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return ""
	}

	return value
}

func WithRequestID(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, requestIDKey, value)
}
