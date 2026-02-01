package http

import (
	nethttp "net/http"
	"payment-gateway/internal/ports/output"

	"github.com/jackc/pgx/v5/pgxpool"

	"payment-gateway/internal/adapters/input/http/handlers"
)

type Handlers struct {
	health *handlers.HealthHandler
}

func NewHandlers(log output.Logger, pool *pgxpool.Pool) *Handlers {
	return &Handlers{
		health: handlers.NewHealthHandler(log, pool),
	}
}

func (h *Handlers) Health(w nethttp.ResponseWriter, r *nethttp.Request) {
	h.health.Handle(w, r)
}
