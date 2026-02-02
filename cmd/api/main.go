package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/VKappaKV/fantasy-ranker-backend/internal/config"
	httpx "github.com/VKappaKV/fantasy-ranker-backend/internal/http"
	"github.com/VKappaKV/fantasy-ranker-backend/internal/storage"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cfg := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var db *storage.DB
	if cfg.DBURL != "" {
		pool, err := storage.NewPool(ctx, cfg.DBURL)
		if err != nil {
			logger.Error("db connection failed", "err", err)
			os.Exit(1)
		}
		db = &storage.DB{Pool: pool}
		defer pool.Close()
		logger.Info("connected to db")
	} else {
		logger.Warn("no db url provided, running without db")
	}

	handler := httpx.NewRouter(cfg, db)

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func () {
		logger.Info("starting api", "addr",cfg.HTTPAddr, "env", cfg.Env)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server error", "err", err)
			stop()
		}
	}()

	<- ctx.Done() // wait for interrupt

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(shutdownCtx)
	logger.Info("server stopped")

}
