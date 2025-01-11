package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	r *chi.Mux
}

func (s Server) Start() error {
	return http.ListenAndServe(":3003", s.r)
}

func NewAPI() *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	return &Server{r}
}
