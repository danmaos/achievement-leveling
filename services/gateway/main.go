package main

import (
	"fmt"
	"log"
	"net/http"

	"achievement-leveling/gateway/config"
	"achievement-leveling/gateway/handler"
	mw "achievement-leveling/gateway/middleware"
	"achievement-leveling/gateway/oauth"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.Load()

	googleOAuth := oauth.NewGoogleOAuth(cfg.GoogleClientID, cfg.GoogleClientSecret, cfg.GoogleRedirectURL)
	authHandler := handler.NewAuthHandler(cfg, googleOAuth)
	proxy := handler.NewProxy(cfg.AchievementSvcURL)

	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(mw.CORS(cfg.FrontendURL))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok","service":"gateway"}`))
	})

	// Public auth routes
	r.Get("/auth/google/login", authHandler.GoogleLogin)
	r.Get("/auth/google/callback", authHandler.GoogleCallback)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(mw.JWTAuth(cfg.JWTSecret))
		r.Get("/auth/me", authHandler.Me)
		// Proxy all /api/* to achievement service
		r.Handle("/api/*", proxy)
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("gateway listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
