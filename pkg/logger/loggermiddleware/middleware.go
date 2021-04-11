package loggermiddleware

import (
	"context"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

type LogApplier func(ctx context.Context, logger zerolog.Logger) zerolog.Logger

func RequestIDApply(ctx context.Context, logger zerolog.Logger) zerolog.Logger {
	requestID := middleware.GetReqID(ctx)
	if len(requestID) == 0 {
		return logger
	}
	return logger.With().Str("requestId", requestID).Logger()
}

func ContextApplying(logAppliers ...LogApplier) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := log.With().Logger()
			for _, apply := range logAppliers {
				logger = apply(r.Context(), logger)
			}
			ctx := logger.WithContext(r.Context())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
