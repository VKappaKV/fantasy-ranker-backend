package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/VKappaKV/fantasy-ranker-backend/internal/config"
	"github.com/VKappaKV/fantasy-ranker-backend/internal/http/handlers"
	"github.com/VKappaKV/fantasy-ranker-backend/internal/riot"
	"github.com/VKappaKV/fantasy-ranker-backend/internal/storage"
)

func NewRouter(cfg config.Config, db *storage.DB) http.Handler {
	r := chi.NewRouter()
	riotClient := riot.New(cfg.RiotAPIKey)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	
	r.Get("/health", handlers.Health())
	r.Get("/version", handlers.Version(cfg.Version))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/riot", func(r chi.Router) {
			r.Get("/account", handlers.RiotAccount(riotClient))
			r.Get("/matches", handlers.RiotMatches(riotClient))
		})
	})

	return r
}
