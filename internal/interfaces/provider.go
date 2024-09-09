package interfaces

import "github.com/MyrtIO/myrtio-go"

type TransportProvider interface {
	Get() myrtio.Transport
}
