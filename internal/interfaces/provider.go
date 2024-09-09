package interfaces

import "github.com/MyrtIO/myrtio-go"

// TransportProvider provides MyrtIO transport.
type TransportProvider interface {
	Get() (myrtio.Transport, error)
}
