package main

import (
	"log"
	"pixie_adapter/internal/app"
	"pixie_adapter/internal/usecase"
	"pixie_adapter/pkg/pixie"
	"strconv"
)

const clockPath = "/dev/cu.wchusbserial14320"
const clockBaudRate = 9600
const httpPort = 17085

func main() {
	u := usecase.New()
	p := pixie.NewConnection(clockPath, clockBaudRate)
	s := app.New(u, p)
	log.Printf("Starting server on " + strconv.Itoa(httpPort) + " port")
	s.Start(httpPort)
}
