package interfaces

import (
	"pixie_adapter/internal/entity"
	"time"
)

type Repositories interface {
	Time() TimeRepository
	Light() LightRepository
	System() SystemRepository
}

type TimeRepository interface {
	Set(time.Time) error
}

type LightRepository interface {
	SetState(entity.LightState) (bool, error)
	GetState() (entity.LightState, error)
}

type SystemRepository interface {
	IsConnected() bool
}
