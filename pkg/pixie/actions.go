package pixie

import (
	"time"

	"github.com/MyrtIO/myrtio-go"
)

// SetTime sets the time of the pixie clock
func SetTime(tx myrtio.Transport, t time.Time) (bool, error) {
	hour := byte(t.Hour())
	minute := byte(t.Minute())
	second := byte(t.Second())
	resp, err := tx.RunAction(&myrtio.Message{
		Feature: FeatureClock,
		Action:  ActionClockSetTime,
		Payload: []byte{hour, minute, second},
	})
	if err != nil {
		return false, err
	}
	return resp.Success(), nil
}

// SetColor sets the color of the pixie clock
func SetColor(tx myrtio.Transport, r, g, b byte) (bool, error) {
	resp, err := tx.RunAction(&myrtio.Message{
		Feature: FeatureIndicators,
		Action:  ActionIndicatorsSetColor,
		Payload: []byte{r, g, b},
	})
	if err != nil {
		return false, err
	}
	return resp.Success(), nil
}

// SetBrightness sets the brightness of the pixie clock
func SetBrightness(tx myrtio.Transport, brightness byte) (bool, error) {
	resp, err := tx.RunAction(&myrtio.Message{
		Feature: FeatureIndicators,
		Action:  ActionIndicatorsSetBrightness,
		Payload: []byte{brightness},
	})
	if err != nil {
		return false, err
	}
	return resp.Success(), nil
}

// SetEffect sets the effect of the pixie clock
func SetEffect(tx myrtio.Transport, effectCode byte) (bool, error) {
	resp, err := tx.RunAction(&myrtio.Message{
		Feature: FeatureIndicators,
		Action:  ActionIndicatorsSetEffect,
		Payload: []byte{effectCode},
	})
	if err != nil {
		return false, err
	}
	return resp.Success(), nil
}

// GetEffect gets the effect of the pixie clock
func GetEffect(tx myrtio.Transport) (byte, error) {
	resp, err := tx.RunAction(&myrtio.Message{
		Feature: FeatureIndicators,
		Action:  ActionIndicatorsGetEffect,
	})
	if err != nil {
		return 0, err
	}
	return resp.Payload[1], nil
}

// Ping checks if the pixie clock is connected
func Ping(tx myrtio.Transport) bool {
	resp, err := tx.RunAction(&myrtio.Message{
		Feature: FeatureSystem,
		Action:  ActionSystemPing,
	})
	if err != nil {
		return false
	}
	return resp.Success()
}

// GetColor gets the color of the pixie clock
func GetColor(tx myrtio.Transport) ([]byte, error) {
	resp, err := tx.RunAction(&myrtio.Message{
		Feature: FeatureIndicators,
		Action:  ActionIndicatorsGetColor,
	})
	if err != nil {
		return nil, err
	}
	return resp.SkipStatus(), nil
}

// GetBrightness gets the brightness of the pixie clock
func GetBrightness(tx myrtio.Transport) (byte, error) {
	resp, err := tx.RunAction(&myrtio.Message{
		Feature: FeatureIndicators,
		Action:  ActionIndicatorsGetBrightness,
	})
	if err != nil {
		return 0, err
	}
	return resp.Payload[1], nil
}

// SetPower sets the power of the pixie clock
func SetPower(tx myrtio.Transport, enabled bool) (bool, error) {
	var enabledByte byte
	if enabled {
		enabledByte = 1
	}
	resp, err := tx.RunAction(&myrtio.Message{
		Feature: FeatureIndicators,
		Action:  ActionIndicatorsSetPower,
		Payload: []byte{enabledByte},
	})
	if err != nil {
		return false, err
	}
	return resp.Success(), nil
}

// GetPower gets the power of the pixie clock
func GetPower(tx myrtio.Transport) (bool, error) {
	resp, err := tx.RunAction(&myrtio.Message{
		Feature: FeatureIndicators,
		Action:  ActionIndicatorsGetPower,
	})
	if err != nil {
		return false, err
	}
	return resp.Payload[1] == 1, nil
}
