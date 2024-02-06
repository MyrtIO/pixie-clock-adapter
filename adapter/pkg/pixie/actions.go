package pixie

import (
	"time"

	"github.com/MyrtIO/myrtio-go"
)

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

func SetPower(tx myrtio.Transport, enabled bool) (bool, error) {
	var enabledByte byte = 0
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
