package pixie

import (
	"errors"
	"sync"

	"github.com/MyrtIO/myrtio-go"
	"github.com/MyrtIO/myrtio-go/serial"
)

var ErrNotConnected = errors.New("not connected")

type Connection struct {
	Path     string
	BaudRate int
	port     myrtio.Transport
	mu       sync.Mutex
}

func NewConnection(path string, baudRate int) *Connection {
	return &Connection{
		Path:     path,
		BaudRate: baudRate,
	}
}

func (p *Connection) Get() (myrtio.Transport, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.port != nil && Ping(p.port) {
		return p.port, nil
	}
	port, err := serial.New(p.Path, p.BaudRate)
	if err != nil {
		return nil, ErrNotConnected
	}
	p.port = port
	return port, nil
}
