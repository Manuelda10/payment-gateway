package http

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *Handlers) {
	r.Get("/health", h.Health)
}
