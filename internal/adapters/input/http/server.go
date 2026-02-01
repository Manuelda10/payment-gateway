package http

import (
	"net/http"
	"payment-gateway/internal/adapters/input/http/middleware"
	"payment-gateway/internal/ports/output"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Server struct {
	log    *zap.Logger
	router chi.Router
}

func NewServer(base *zap.Logger, appLogger output.Logger) *Server {
	r := chi.NewRouter()

	r.Use(middleware.WithRequestID)
	r.Use(middleware.AccessLog(appLogger))

	s := &Server{
		log:    base,
		router: r,
	}
	return s
}

func (s *Server) Router() http.Handler {
	return s.router
}

func (s *Server) HTTPServer(addr string) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           s.router,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

func (s *Server) Chi() chi.Router { return s.router }
