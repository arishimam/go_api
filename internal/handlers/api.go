package handlers

import (
	"github.com/arishimam/go_api/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// starting this with uppercase means it can be imported elsewhere. Lowercase = private
func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
