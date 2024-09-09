package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"pixie_adapter/internal/config"
	"pixie_adapter/internal/controller/mqtt"
	"pixie_adapter/internal/repository"
	"pixie_adapter/internal/worker"
	"pixie_adapter/pkg/pixie"

	"github.com/MyrtIO/myrtio-go/serial"
)

type Application struct {
	mqtt   *mqtt.Controller
	runner *worker.Runner
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

	portPath := config.PortPath
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
	baudRate := config.BaudRate
	if baudRate == 0 {
		baudRate = 28800
	}

	conn := pixie.NewConnection(portPath, baudRate)
	repos := repository.New(conn)
	mqtt := mqtt.New(config, repos)

	if repos.System().IsConnected() {
		log.Println("Clock is connected")
	} else {
		log.Println("Clock is not connected")
	}

	runner := worker.NewRunner(
		worker.NewTimeSync(repos),
		worker.NewPing(repos),
	)

	return &Application{
		mqtt:   mqtt,
		runner: runner,
	}
}

// Start is a method of the Application struct that starts the application.
func (app *Application) Start() error {
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)

	defer cancel()

	app.runner.Run(ctx.Done())
	err := app.mqtt.Start(ctx)
	if err != nil {
		return err
	}

	return nil
}
