package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/VKappaKV/fantasy-ranker-backend/internal/config"
	"github.com/VKappaKV/fantasy-ranker-backend/internal/http/handlers"
)

func NewRouter(cfg config.Config) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * 1e9))

	r.Get("/health", handlers.Health())
	r.Get("/version", handlers.Version(cfg.Version))

	// future:
	// r.Route("/api/v1", func(r chi.Router) { ... })

	return r
}
