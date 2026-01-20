package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/VKappaKV/fantasy-ranker-backend/internal/config"
	httpx "github.com/VKappaKV/fantasy-ranker-backend/internal/http"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cfg := config.Load()

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           httpx.NewRouter(cfg),
		ReadHeaderTimeout: 5 * time.Second,
	}

	logger.Info("starting api", "addr", cfg.HTTPAddr, "env", cfg.Env)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("server error", "err", err)
		os.Exit(1)
	}
}
