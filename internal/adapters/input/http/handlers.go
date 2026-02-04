package http

import (
	nethttp "net/http"
	"payment-gateway/internal/ports/input"
	"payment-gateway/internal/ports/output"

	"github.com/jackc/pgx/v5/pgxpool"

	"payment-gateway/internal/adapters/input/http/handlers"
)

type Handlers struct {
	health *handlers.HealthHandler
	niubiz *handlers.NiubizHandler
}

func NewHandlers(log output.Logger, pool *pgxpool.Pool, niubizSvc input.NiubizService) *Handlers {
	return &Handlers{
		health: handlers.NewHealthHandler(log, pool),
		niubiz: handlers.NewNiubizHandler(log, niubizSvc),
	}
}

func (h *Handlers) Health(w nethttp.ResponseWriter, r *nethttp.Request) {
	h.health.Handle(w, r)
}

func (h *Handlers) NiubizSecurityToken(w nethttp.ResponseWriter, r *nethttp.Request) {
	h.niubiz.GetSecurityToken(w, r)
}
