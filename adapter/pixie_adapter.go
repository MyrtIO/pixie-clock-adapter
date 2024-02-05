package main

import (
	"pixie_adapter/internal/pixie"
	"pixie_adapter/internal/server"
)

const clockPath = "/dev/cu.wchusbserial14320"
const clockBaudRate = 9600
const httpPort = 17085

func main() {
	connection := pixie.NewConnection(clockPath, clockBaudRate)
	server.Run(httpPort, connection)
}
