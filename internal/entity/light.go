package entity

// RGBColor represents an RGB color
type RGBColor struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
}

// RGBColorFromSlice converts a slice of values to an RGBColor
func RGBColorFromSlice[T uint8 | int](values []T) RGBColor {
	return RGBColor{
		R: uint8(values[0]),
		G: uint8(values[1]),
		B: uint8(values[2]),
	}
}

// LightPowerState represents the power state of the light
type LightPowerState string

const (
	// LightPowerStateOn represents the light is on
	LightPowerStateOn LightPowerState = "ON"
	// LightPowerStateOff represents the light is off
	LightPowerStateOff LightPowerState = "OFF"
)

// Bool returns true if the light is on
func (l LightPowerState) Bool() bool {
	return l == LightPowerStateOn
}

// LightPowerStateFromBool returns the power state of the light
func LightPowerStateFromBool(enabled bool) LightPowerState {
	if enabled {
		return LightPowerStateOn
	}
	return LightPowerStateOff
}

// LightEffect represents the effect of the pixie clock light
type LightEffect string

const (
	//revive:disable
	LightEffectStatic LightEffect = "static"
	LightEffectSmooth LightEffect = "smooth"
	LightEffectZoom   LightEffect = "zoom"
	//revive:enable
)

var lightEffectCodes = map[LightEffect]uint8{
	LightEffectStatic: 0,
	LightEffectSmooth: 1,
	LightEffectZoom:   2,
}

// Code returns the code of the effect
func (l LightEffect) Code() uint8 {
	return lightEffectCodes[l]
}

// LightEffectFromCode returns the effect from the code
func LightEffectFromCode(code uint8) LightEffect {
	for k, v := range lightEffectCodes {
		if v == code {
			return k
		}
	}
	return LightEffectStatic
}

// LightState represents the state of the pixie clock light
type LightState struct {
	State      LightPowerState `json:"state"`
	Color      *RGBColor       `json:"color"`
	Effect     *LightEffect    `json:"effect"`
	Brightness *uint8          `json:"brightness"`
}
