package handlers

import (
	"encoding/json"
	"net/http"
	"payment-gateway/internal/ports/input"
	"payment-gateway/internal/ports/output"
)

type NiubizHandler struct {
	log output.Logger
	svc input.NiubizService
}

func NewNiubizHandler(log output.Logger, svc input.NiubizService) *NiubizHandler {
	return &NiubizHandler{log: log, svc: svc}
}

func (h *NiubizHandler) GetSecurityToken(w http.ResponseWriter, r *http.Request) {
	token, err := h.svc.GetSecurityToken(r.Context())
	if err != nil {
		h.log.Error(r.Context(), "niubiz: handler failed", err)
		w.WriteHeader(http.StatusBadGateway)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"code":    "NIUBIZ_TOKEN_ERROR",
			"message": err.Error(),
		})
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]any{"token": token})
}
