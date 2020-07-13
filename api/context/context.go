package context

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const (
	requestIDKey     contextKey = "request_id"
	userIDKey        contextKey = "user_id"
	accessTokenIDKey contextKey = "access_token_id"
	accessTokenKey   contextKey = "access_token"
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

func GetAccessToken(ctx context.Context) (value string) {
	value, ok := ctx.Value(accessTokenKey).(string)
	if !ok {
		return ""
	}

	return value
}

func WithAccessToken(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, accessTokenKey, value)
}

func GetUserID(ctx context.Context) uuid.UUID {
	value, _ := ctx.Value(userIDKey).(uuid.UUID)
	return value
}

func WithUserID(ctx context.Context, value uuid.UUID) context.Context {
	return context.WithValue(ctx, userIDKey, value)
}

func GetAccessTokenID(ctx context.Context) uuid.UUID {
	value, _ := ctx.Value(accessTokenIDKey).(uuid.UUID)
	return value
}

func WithAccessTokenID(ctx context.Context, value uuid.UUID) context.Context {
	return context.WithValue(ctx, accessTokenIDKey, value)
}
