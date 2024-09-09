package repository

import (
	"fmt"
	"pixie_adapter/internal/entity"
	"pixie_adapter/internal/interfaces"
	"pixie_adapter/pkg/pixie"
	"pixie_adapter/pkg/ptr"
)

// LightRepository provides access to the light
type LightRepository struct {
	conn interfaces.TransportProvider
}

var _ interfaces.LightRepository = (*LightRepository)(nil)

func newLightRepository(conn interfaces.TransportProvider) *LightRepository {
	return &LightRepository{
		conn: conn,
	}
}

// GetState returns the current state of the light
func (c *LightRepository) GetState() (entity.LightState, error) {
	var state entity.LightState
	tx, err := c.conn.Get()
	if err != nil {
		return state, err
	}

	brightness, err := pixie.GetBrightness(tx)
	if err != nil {
		return state, fmt.Errorf("error getting brightness: %s", err)
	}
	color, err := pixie.GetColor(tx)
	if err != nil {
		return state, fmt.Errorf("error getting color: %s", err)
	}
	effect, err := pixie.GetEffect(tx)
	if err != nil {
		return state, fmt.Errorf("error getting effect: %s", err)
	}
	isEnabled, err := pixie.GetPower(tx)
	if err != nil {
		return state, fmt.Errorf("error getting power: %s", err)
	}

	return entity.LightState{
		Brightness: ptr.From(brightness),
		Color:      ptr.From(entity.RGBColorFromSlice(color)),
		Effect:     ptr.From(entity.LightEffectFromCode(effect)),
		State:      entity.LightPowerStateFromBool(isEnabled),
	}, nil
}

// SetState sets the state of the light
func (c *LightRepository) SetState(state entity.LightState) (hasChanges bool, err error) {
	tx, err := c.conn.Get()
	if err != nil {
		return
	}

	if state.Effect != nil {
		effect, err := pixie.GetEffect(tx)
		if err != nil {
			return hasChanges, fmt.Errorf("error getting effect: %s", err)
		}
		stateEffect := state.Effect.Code()
		if effect != stateEffect {
			_, err = pixie.SetEffect(tx, stateEffect)
			if err != nil {
				return hasChanges, fmt.Errorf("error setting effect: %s", err)
			}
			hasChanges = true
		}
	}
	if state.Brightness != nil {
		brightness, err := pixie.GetBrightness(tx)
		if err != nil {
			return hasChanges, fmt.Errorf("error getting brightness: %s", err)
		}
		if *state.Brightness != brightness {
			_, err = pixie.SetBrightness(tx, *state.Brightness)
			if err != nil {
				return hasChanges, fmt.Errorf("error setting brightness: %s", err)
			}
			hasChanges = true
		}
	}
	if state.Color != nil {
		color, err := pixie.GetColor(tx)
		if err != nil {
			return hasChanges, fmt.Errorf("error getting color: %s", err)
		}
		if *state.Color != entity.RGBColorFromSlice(color) {
			_, err = pixie.SetColor(tx, state.Color.R, state.Color.G, state.Color.B)
			if err != nil {
				return hasChanges, fmt.Errorf("error setting color: %s", err)
			}
			hasChanges = true
		}
	}
	isEnabled, err := pixie.GetPower(tx)
	if err != nil {
		return hasChanges, fmt.Errorf("error getting power: %s", err)
	}
	if state.State.Bool() != isEnabled {
		_, err = pixie.SetPower(tx, state.State.Bool())
		if err != nil {
			return hasChanges, fmt.Errorf("error setting power: %s", err)
		}
		hasChanges = true
	}
	return hasChanges, nil
}
