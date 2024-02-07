package main

import (
	"fmt"
	"log"
	"os"
	"pixie_adapter/internal/app"
	"pixie_adapter/internal/usecase"
	"pixie_adapter/pkg/pixie"
	"strconv"

	"github.com/MyrtIO/myrtio-go/serial"
)

const clockBaudRate = 28800
const httpPort = 17085

func main() {
	paths, err := serial.Discover()
	if err != nil {
		log.Panic(err)
	}
	if len(paths) == 0 {
		fmt.Println("Serial devices is not found")
		os.Exit(1)
	}
	p := pixie.NewConnection(paths[0], clockBaudRate)
	u := usecase.New()
	s := app.New(u, p)
	log.Printf("Starting server on " + strconv.Itoa(httpPort) + " port")
	s.Start(httpPort)
}
