package application

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"pixie_adapter/internal/config"
	"pixie_adapter/internal/controller/mqtt"
	"pixie_adapter/internal/interfaces"
	"pixie_adapter/internal/repository"
	"pixie_adapter/internal/worker"
	"pixie_adapter/pkg/pixie"
	"syscall"

	"github.com/MyrtIO/myrtio-go/serial"
)

// Application represents the core of the service
type Application struct {
	mqtt   *mqtt.Controller
	runner *worker.Runner
	repos  *repository.Repository
}

// New initializes a new Application.
func New(configPath string) (*Application, error) {
	config, err := config.New(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	if config == nil || config.Host == "" {
		return nil, fmt.Errorf("config is not loaded: %s", configPath)
	}

	var conn interfaces.TransportProvider
	if config.Serial.Port != "" {
		conn = pixie.NewConnection(config.Serial.Port, config.Serial.BaudRate)
	} else {
		paths, err := serial.Discover()
		if err != nil {
			return nil, fmt.Errorf("failed to discover serial devices: %w", err)
		}
		if len(paths) == 0 {
			return nil, errors.New("no supported serial device is found")
		}
		selectedPort := paths[0]
		conn = pixie.NewConnection(selectedPort, config.Serial.BaudRate)
	}

	repos := repository.New(conn)
	mqtt := mqtt.New(config, repos)
	runner := worker.NewRunner(
		worker.NewTimeSync(repos),
	)

	return &Application{
		mqtt:   mqtt,
		runner: runner,
		repos:  repos,
	}, nil
}

// Start is a method of the Application struct that starts the application.
func (a *Application) Start() error {
	if a.repos.System().IsConnected() {
		log.Println("Clock is connected")
	} else {
		log.Println("Clock is not connected")
	}

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx,
		syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	defer cancel()

	a.runner.Run(ctx.Done())

	go func() {
		err := a.mqtt.Start(ctx)
		if err != nil {
			log.Printf("Failed to start MQTT: %s", err)
		}
	}()

	<-ctx.Done()
	log.Println("Stopping application...")
	return nil
}
