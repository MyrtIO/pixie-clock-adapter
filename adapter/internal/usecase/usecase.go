package usecase

import "pixie_adapter/internal/interfaces"

type Usecase struct {
	Lights interfaces.LightsUsecase
}

func New() interfaces.Usecase {
	usecase := interfaces.Usecase{
		Lights: NewLights(),
	}

	return usecase
}
