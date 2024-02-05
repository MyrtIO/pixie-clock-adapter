package app

import (
	"net/http"
	"pixie_adapter/internal/interfaces"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Service struct {
	router   *chi.Mux
	usecase  interfaces.Usecase
	provider interfaces.TransportProvider
}

func New(usecase interfaces.Usecase, provider interfaces.TransportProvider) *Service {
	r := chi.NewRouter()
	s := Service{
		router:   r,
		usecase:  usecase,
		provider: provider,
	}
	s.Setup()
	return &s
}

func (s *Service) Setup() {
	s.router.Use(middleware.Logger)
	s.router.Get("/", s.GetState)
	s.router.Put("/", s.SetState)
	s.router.Post("/disable", s.DisableLights)
}

func (s *Service) Start(port int) {
	http.ListenAndServe(":"+strconv.Itoa(port), s.Handler())
}

func (s *Service) Handler() http.Handler {
	return s.router
}
