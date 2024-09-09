package interfaces

import (
	"pixie_adapter/internal/entity"
	"time"
)

// Repositories provides access to the device features
type Repositories interface {
	Time() TimeRepository
	Light() LightRepository
	System() SystemRepository
}

// TimeRepository provides access to the time
type TimeRepository interface {
	Set(time.Time) error
}

// LightRepository provides access to the light
type LightRepository interface {
	SetState(entity.LightState) (bool, error)
	GetState() (entity.LightState, error)
}

// SystemRepository provides access to the system
type SystemRepository interface {
	IsConnected() bool
}
