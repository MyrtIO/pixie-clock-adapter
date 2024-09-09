package pixie

import (
	"errors"
	"sync"

	"github.com/MyrtIO/myrtio-go"
	"github.com/MyrtIO/myrtio-go/serial"
)

// ErrNotConnected is returned when connection is not established
var ErrNotConnected = errors.New("not connected")

// Connection is a connection to the pixie clock
type Connection struct {
	Path     string
	BaudRate int
	port     myrtio.Transport
	mu       sync.Mutex
}

// NewConnection creates a new Connection
func NewConnection(path string, baudRate int) *Connection {
	return &Connection{
		Path:     path,
		BaudRate: baudRate,
	}
}

// Get returns a connection to the pixie clock
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
