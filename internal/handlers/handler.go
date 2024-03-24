package handlers

import (
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

func Handler(r *chi.Mux) {

	// Global middleware
	r.Use(chimiddle.StripSlashes)

	// User route
	r.Route("/user", func(r chi.Router) {
		r.Get("/", ListUser)
		r.Put("/", PutUser)
		r.Get("/{id}", GetUser)
		r.Delete("/{id}", DeleteUser)
	})

}
