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
	SetColor(entity.RGBColor) error
	GetColor() (entity.RGBColor, error)
	SetBrightness(uint8) error
	GetBrightness() (uint8, error)
	SetPower(enabled bool) error
	GetPower() (bool, error)
	SetEffect(entity.LightEffect) error
	GetEffect() (entity.LightEffect, error)
}

type SystemRepository interface {
	IsConnected() bool
}
