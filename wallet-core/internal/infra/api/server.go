package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router   chi.Router
	Handlers map[string]http.HandlerFunc
	Port     string
}

func NewServer(port string) *Server {
	return &Server{
		Router:   chi.NewRouter(),
		Handlers: make(map[string]http.HandlerFunc),
		Port:     port,
	}
}

func (s *Server) AddRoute(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *Server) Start() {
	s.Router.Use(middleware.Logger)

	for path, handler := range s.Handlers {
		s.Router.Post(path, handler)
	}

	log.Println("server started at port", s.Port)
	err := http.ListenAndServe(s.Port, s.Router)
	if err != nil {
		log.Fatal("failed to start server", err)
	}
}
