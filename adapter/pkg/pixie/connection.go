package pixie

import (
	"sync"
	"time"

	"github.com/MyrtIO/myrtio-go"
	"github.com/MyrtIO/myrtio-go/serial"
)

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

func (p *Connection) Get() myrtio.Transport {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.port != nil && Ping(p.port) {
		return p.port
	}
	port, err := serial.New(p.Path, p.BaudRate)
	if err != nil {
		return nil
	}
	p.port = port
	// Update current time on connect
	SetTime(port, time.Now())
	return port
}
