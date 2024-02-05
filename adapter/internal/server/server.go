package server

import (
	"log"
	"net/http"
	"pixie_adapter/internal/pixie"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	router     *chi.Mux
	connection *pixie.Connection
}

func New(connection *pixie.Connection) *Server {
	r := chi.NewRouter()
	s := Server{
		router:     r,
		connection: connection,
	}
	r.Use(middleware.Logger)
	r.Get("/", s.handleGetState)
	r.Put("/", s.handleSetState)
	r.Post("/disable", s.handleDisable)
	return &s
}

func (s *Server) Handler() http.Handler {
	return s.router
}

func Run(port int, connection *pixie.Connection) {
	s := New(connection)
	log.Printf("Starting server")
	http.ListenAndServe(":"+strconv.Itoa(port), s.Handler())
}
