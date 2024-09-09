package app

import (
	"context"
	"log"
	"os"
	"pixie_adapter/internal/config"
	"pixie_adapter/internal/controller/mqtt"
	"pixie_adapter/internal/repository"
	"pixie_adapter/internal/worker"
	"pixie_adapter/pkg/pixie"

	"github.com/MyrtIO/myrtio-go/serial"
)

// Application represents the core of the service
type Application struct {
	mqtt   *mqtt.Controller
	runner *worker.Runner
	repos  *repository.Repository
}

// New initializes a new Application.
func New(configPath string) *Application {
	config, err := config.New(configPath)
	if err != nil {
		log.Panic(err)
	}
	if config == nil || config.Host == "" {
		log.Panicf("Config is not loaded: %s", configPath)
	}

	portPath := config.Serial.Port
	if portPath == "" {
		paths, err := serial.Discover()
		if err != nil {
			log.Panic(err)
		}
		if len(paths) == 0 {
			log.Println("Serial devices is not found")
			os.Exit(1)
		}
		portPath = paths[0]
		log.Printf("Serial device is found at %s", portPath)
	}
	baudRate := config.BaudRate
	if baudRate == 0 {
		baudRate = 28800
	}

	conn := pixie.NewConnection(portPath, baudRate)
	repos := repository.New(conn)
	mqtt := mqtt.New(config, repos)

	runner := worker.NewRunner(
		worker.NewTimeSync(repos),
	)

	return &Application{
		mqtt:   mqtt,
		runner: runner,
		repos:  repos,
	}
}

// Start is a method of the Application struct that starts the application.
func (a *Application) Start() error {
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)

	defer cancel()

	if a.repos.System().IsConnected() {
		log.Println("Clock is connected")
	} else {
		log.Println("Clock is not connected")
	}

	a.runner.Run(ctx.Done())
	err := a.mqtt.Start(ctx)
	if err != nil {
		return err
	}

	return nil
}
