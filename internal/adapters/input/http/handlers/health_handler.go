package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"payment-gateway/internal/ports/output"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthHandler struct {
	log  output.Logger
	pool *pgxpool.Pool
}

func NewHealthHandler(log output.Logger, pool *pgxpool.Pool) *HealthHandler {
	return &HealthHandler{log: log, pool: pool}
}

func (h *HealthHandler) Handle(w http.ResponseWriter, r *http.Request) {
	dbStatus := "ok"

	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	if err := h.pool.Ping(ctx); err != nil {
		dbStatus = "down"
		h.log.Error(r.Context(), "db ping failed", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
		"db":     dbStatus,
	})
}
