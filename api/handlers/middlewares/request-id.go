package middlewares

import (
	"net/http"

	reqContext "github.com/eugeneradionov/stocks/api/context"
	"github.com/google/uuid"
)

const requestIDHeader = "X-Request-Id"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.New().String()
			}

			reqContext.WithRequestID(ctx, requestID)

			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
