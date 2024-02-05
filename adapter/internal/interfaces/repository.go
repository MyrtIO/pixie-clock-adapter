package interfaces

import (
	"time"

	"github.com/MyrtIO/myrtio-go"
)

type Repository struct {
	Pixie PixieRepository
}

type PixieRepository interface {
	SetTime(tx myrtio.Transport, t time.Time) (bool, error)
	SetColor(tx myrtio.Transport, r, g, b byte) (bool, error)
	GetColor(tx myrtio.Transport) ([]byte, error)
	SetBrightness(tx myrtio.Transport, brightness byte) (bool, error)
	GetBrightness(tx myrtio.Transport) (uint8, error)
	SetPower(tx myrtio.Transport, enabled bool) (bool, error)
	GetPower(tx myrtio.Transport) (bool, error)
}
