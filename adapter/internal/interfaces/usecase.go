package interfaces

import (
	"pixie_adapter/internal/dto"

	"github.com/MyrtIO/myrtio-go"
)

type Usecase struct {
	Lights LightsUsecase
}

type LightsUsecase interface {
	SetState(tx myrtio.Transport, request *dto.LightsStateRequest) error
	GetState(tx myrtio.Transport) (*dto.LightsStateResponse, error)
	TurnOff(tx myrtio.Transport) error
}
