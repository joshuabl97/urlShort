package handler

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

// LoggingMiddleware accepts a HandlerFunc and a *zerolog.Logger,
// and returns a new HandlerFunc that logs information about incoming requests.
func Logger(handler http.HandlerFunc, l *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the provided handler function
		handler(w, r)

		// Logging the request details
		l.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Str("remote_addr", r.RemoteAddr).
			Dur("duration", time.Since(start)).
			Msg("request handled")
	}
}
