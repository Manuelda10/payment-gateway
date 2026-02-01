package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"payment-gateway/config"
	httpapi "payment-gateway/internal/adapters/input/http"
	zaplogger "payment-gateway/internal/adapters/output/logging/zap"
	pg "payment-gateway/internal/adapters/output/persistence/postgres"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	if os.Getenv("APP_ENV") == "dev" {
		fmt.Println("APP_ENV is development")
		_ = godotenv.Load()
	}
	fmt.Println("Starting Payment Gateway API...")
	cfg := config.Load()

	zlog, err := zaplogger.NewBase(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = zlog.Sync() }()
	appLogger := zaplogger.New(zlog)

	if cfg.DBDsnEmpty() {
		zlog.Fatal("DB_DSN is required")
	}

	ctx := context.Background()

	pool, err := pg.NewPool(ctx, cfg.DBDSN)
	if err != nil {
		zlog.Fatal("failed to create db pool", zap.Error(err))
	}
	defer pool.Close()

	srv := httpapi.NewServer(zlog, appLogger)
	h := httpapi.NewHandlers(appLogger, pool)

	httpapi.RegisterRoutes(srv.Chi(), h)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	zlog.Info("starting api", zap.String("addr", addr))

	if err := srv.HTTPServer(addr).ListenAndServe(); err != nil {
		zlog.Fatal("server error", zap.Error(err))
	}
}
