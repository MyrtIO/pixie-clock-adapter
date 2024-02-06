package usecase

import (
	"pixie_adapter/internal/dto"
	"pixie_adapter/pkg/pixie"

	"github.com/MyrtIO/myrtio-go"
)

type Lights struct{}

func NewLights() *Lights {
	return &Lights{}
}

func (l *Lights) SetState(tx myrtio.Transport, request *dto.LightsStateRequest) error {
	_, err := pixie.SetPower(tx, true)
	if err != nil {
		return err
	}
	_, err = pixie.SetColor(tx, request.Color[0], request.Color[1], request.Color[2])
	if err != nil {
		return err
	}
	_, err = pixie.SetBrightness(tx, request.Brightness)
	if err != nil {
		return err
	}
	_, err = pixie.SetEffect(tx, request.Effect)
	if err != nil {
		return err
	}
	return nil
}

func (l *Lights) GetState(tx myrtio.Transport) (*dto.LightsStateResponse, error) {
	color, err := pixie.GetColor(tx)
	if err != nil {
		return nil, err
	}
	brightness, err := pixie.GetBrightness(tx)
	if err != nil {
		return nil, err
	}
	effect, err := pixie.GetEffect(tx)
	if err != nil {
		return nil, err
	}
	enabled, err := pixie.GetPower(tx)
	if err != nil {
		return nil, err
	}
	return &dto.LightsStateResponse{
		Color:      bytesToInts(color),
		Brightness: brightness,
		Enabled:    enabled,
		Effect:     effect,
	}, nil
}

func (l *Lights) TurnOff(tx myrtio.Transport) error {
	_, err := pixie.SetPower(tx, false)
	if err != nil {
		return err
	}
	return nil
}
