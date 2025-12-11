package http

import (
	_ "gocrud/internal/docs" // важливо!
	h "gocrud/internal/hendlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes(r *chi.Mux, handler *h.UserHandler) {

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.CreateUser)       // POST /users
		r.Get("/{id}", handler.GetUser)       // GET /users/{id}
		r.Put("/{id}", handler.UpdateUser)    // PUT /users/{id}
		r.Delete("/{id}", handler.DeleatUser) // когда допишешь Delete
	})

	r.Get("/swagger/*", httpSwagger.WrapHandler)

}
