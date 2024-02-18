package http

import (
	"fmt"
	"net/http"
	"pixie_adapter/internal/interfaces"
	"pixie_adapter/pkg/pixie"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Service struct {
	router    *chi.Mux
	usecase   interfaces.Usecase
	provider  interfaces.TransportProvider
	port      int
	closePing chan struct{}
}

func NewService(usecase interfaces.Usecase, provider interfaces.TransportProvider) *Service {
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

func (s *Service) Start() error {
	ticker := time.NewTicker(5 * time.Second)
	s.closePing = make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				tx := s.provider.Get()
				if tx != nil {
					pixie.Ping(tx)
				}
			case <-s.closePing:
				ticker.Stop()
				return
			}
		}
	}()
	addr := ":" + strconv.Itoa(s.port)
	fmt.Println("Starting server on " + addr)
	return http.ListenAndServe(addr, s.Handler())
}

func (s *Service) Handler() http.Handler {
	return s.router
}

func (s *Service) SetPort(port int) {
	s.port = port
}
