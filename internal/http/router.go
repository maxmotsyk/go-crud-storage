package http

import (
	"github.com/go-chi/chi/v5"
)

func SetupRouter(r *chi.Mux) {

	r.Route("/users", func(r chi.Router) {
		r.With()
	})

}
