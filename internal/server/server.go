package server

import (
	"encoding/json"
	"log"
	"net/http"
	"pixie_adapter/internal/dto"
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
	return &s
}

func (s *Server) Handler() http.Handler {
	return s.router
}

func (s *Server) handleSetState(w http.ResponseWriter, r *http.Request) {
	var req dto.ClockStateRequest
	d := json.NewDecoder(r.Body)
	err := d.Decode(&req)
	if err != nil {
		log.Printf("Error on parsing request: %v\n", err)
		w.WriteHeader(400)
	}
	tx := s.connection.Get()
	if tx == nil {
		w.WriteHeader(503)
		return
	}
	_, err = pixie.SetColor(tx, req.Color[0], req.Color[1], req.Color[2])
	if err != nil {
		log.Printf("Error on setting color: %v\n", err)
		w.WriteHeader(500)
		return
	}
	_, err = pixie.SetBrightness(tx, req.Brightness)
	if err != nil {
		log.Printf("Error on setting brightness: %v\n", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func (s *Server) handleGetState(w http.ResponseWriter, r *http.Request) {
	tx := s.connection.Get()
	if tx == nil {
		w.WriteHeader(503)
		return
	}
	colorValue, err := pixie.GetColor(tx)
	if err != nil {
		log.Printf("Error on reading color: %v\n", err)
		w.WriteHeader(500)
		return
	}
	brightness, err := pixie.GetBrightness(tx)
	if err != nil {
		log.Printf("Error on reading brightness: %v\n", err)
		w.WriteHeader(500)
		return
	}
	var color dto.ClockColor = colorValue
	jsonResponse(w, 200, dto.ClockStateResponse{
		Color:      color.Ints(),
		Brightness: brightness,
	})
}

func Run(port int, connection *pixie.Connection) {
	s := New(connection)
	log.Printf("Starting server")
	http.ListenAndServe(":"+strconv.Itoa(port), s.Handler())
}
