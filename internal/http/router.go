package http

import (
	"github.com/go-chi/chi/v5"
	h "gocrud/internal/hendlers"
)

func SetupRoutes(r *chi.Mux, handler *h.UserHandler) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.CreateUser)    // POST /users
		r.Get("/{id}", handler.GetUser)    // GET /users/{id}
		r.Put("/{id}", handler.UpdateUser) // PUT /users/{id}
		// r.Delete("/{id}", h.DeleteUser) // когда допишешь Delete
	})
}
