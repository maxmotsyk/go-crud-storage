package http

import (
	h "gocrud/internal/hendlers"
	"gocrud/internal/stor"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router     *chi.Mux
	httpServer *http.Server
}

func CreatServer(storage *stor.Storage) *Server {
	r := chi.NewRouter()
	userHendler := h.NewUserHandler(storage)

	SetupRoutes(r, userHendler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &Server{
		router:     r,
		httpServer: srv,
	}
}

func (s *Server) Listen() error {
	return s.httpServer.ListenAndServe()
}
