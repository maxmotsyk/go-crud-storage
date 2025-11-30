package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	router     *chi.Mux
	httpServer *http.Server
}

func CreatServer() *Server {
	r := chi.NewRouter()

	SetupRouter(&r)

	return &Server{
		router: ,
	}
}
