package entity

type RGBColor struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
}

var (
	ColorBlack RGBColor = RGBColor{0, 0, 0}
)

func RGBColorFromSlice(values []uint8) RGBColor {
	return RGBColor{
		R: values[0],
		G: values[1],
		B: values[2],
	}
}

type LightPowerState string

const (
	LightPowerStateOn  LightPowerState = "ON"
	LightPowerStateOff LightPowerState = "OFF"
)

func (l LightPowerState) Bool() bool {
	return l == LightPowerStateOn
}

func LightPowerStateFromBool(enabled bool) LightPowerState {
	if enabled {
		return LightPowerStateOn
	}
	return LightPowerStateOff
}

type LightColorMode string

const (
	LightColorModeRGB   LightColorMode = "rgb"
	LightColorModeEmpty LightColorMode = ""
)

type LightEffect string

const (
	LightEffectStatic LightEffect = "static"
	LightEffectSmooth LightEffect = "smooth"
	LightEffectZoom   LightEffect = "zoom"
	LightEffectEmpty  LightEffect = ""
)

var lightEffectCodes = map[LightEffect]uint8{
	LightEffectStatic: 0,
	LightEffectSmooth: 1,
	LightEffectZoom:   2,
}

func (l LightEffect) Code() uint8 {
	if l == LightEffectEmpty {
		return 0
	}
	return lightEffectCodes[l]
}

func LightEffectFromCode(code uint8) LightEffect {
	for k, v := range lightEffectCodes {
		if v == code {
			return k
		}
	}
	return LightEffectStatic
}

type LightState struct {
	State      LightPowerState `json:"state"`
	Color      *RGBColor       `json:"color"`
	Effect     *LightEffect    `json:"effect"`
	ColorMode  LightColorMode  `json:"color_mode"`
	Brightness *uint8          `json:"brightness"`
}
