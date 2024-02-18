package app

import (
	"fmt"
	"log"
	"os"
	"pixie_adapter/internal/controller/http"
	"pixie_adapter/internal/usecase"
	"pixie_adapter/pkg/pixie"

	"github.com/MyrtIO/myrtio-go/serial"
)

type Application struct {
	service *http.Service
}

// New initializes a new Application.
func New(portPath string, baudRate int) *Application {
	if portPath == "" {
		paths, err := serial.Discover()
		if err != nil {
			log.Panic(err)
		}
		if len(paths) == 0 {
			fmt.Println("Serial devices is not found")
			os.Exit(1)
		}
		portPath = paths[0]
	}

	p := pixie.NewConnection(portPath, baudRate)
	u := usecase.New()
	s := http.NewService(u, p)
	return &Application{
		service: s,
	}
}

// Start is a method of the Application struct that starts the application.
func (app *Application) Start() error {
	return app.service.Start()
}

// SetPort sets the port for the Application.
func (app *Application) SetPort(port int) {
	app.service.SetPort(port)
}
