package pixie

import (
	"github.com/MyrtIO/myrtio-go"
	"github.com/MyrtIO/myrtio-go/serial"
)

type Connection struct {
	Path     string
	BaudRate int
	port     myrtio.Transport
}

func NewConnection(path string, baudRate int) *Connection {
	return &Connection{
		Path:     path,
		BaudRate: baudRate,
	}
}

func (p *Connection) Get() myrtio.Transport {
	if p.port != nil && Ping(p.port) {
		return p.port
	}
	port, err := serial.New(p.Path, p.BaudRate)
	if err != nil {
		return nil
	}
	p.port = port
	return port
}
