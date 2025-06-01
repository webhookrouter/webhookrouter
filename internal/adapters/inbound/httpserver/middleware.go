package httpserver

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func loggerMiddleware(logger zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()

			defer func() {
				dur := time.Since(t1)
				logger.Debug().
					Str("method", r.Method).
					Str("url", r.URL.String()).
					Int("status", ww.Status()).
					Dur("duration", dur).
					Msg("request completed")
			}()

			next.ServeHTTP(ww, r)
		})

	}
}
