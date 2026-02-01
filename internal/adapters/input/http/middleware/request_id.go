package middleware

import (
	"net/http"
	"payment-gateway/internal/shared/requestctx"
	"strings"

	"github.com/google/uuid"
)

const HeaderRequestID = "X-Request-Id"

func WithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rid := strings.TrimSpace(r.Header.Get(HeaderRequestID))
		if rid == "" {
			rid = uuid.NewString()
		}

		w.Header().Set(HeaderRequestID, rid)

		ctx := requestctx.WithRequestID(r.Context(), rid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
