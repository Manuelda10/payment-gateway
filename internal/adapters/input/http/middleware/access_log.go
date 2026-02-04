package middleware

import (
	"net/http"
	"payment-gateway/internal/ports/output"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (w *statusRecorder) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *statusRecorder) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	n, err := w.ResponseWriter.Write(b)
	w.bytes += n
	return n, err
}

func AccessLog(logger output.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rec := &statusRecorder{ResponseWriter: w}

			next.ServeHTTP(rec, r)

			latencyMs := time.Since(start).Milliseconds()

			fields := []output.Field{
				output.F("method", r.Method),
				output.F("path", r.URL.Path),
				output.F("status", rec.status),
				output.F("bytes", rec.bytes),
				output.F("latency_ms", latencyMs),
				output.F("remote_ip", r.RemoteAddr),
				output.F("user_agent", r.UserAgent()),
			}

			switch {
			case rec.status >= 500:
				logger.Error(r.Context(), "http request", nil, fields...)
			case rec.status >= 400:
				logger.Warn(r.Context(), "http request", fields...)
			default:
				logger.Info(r.Context(), "http request", fields...)
			}
		})
	}
}
