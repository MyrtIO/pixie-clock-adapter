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

func GetBrightness(tx myrtio.Transport) (uint8, error) {
	resp, err := tx.RunAction(&myrtio.Message{
		Feature: FeatureIndicators,
		Action:  ActionIndicatorsGetColor,
	})
	if err != nil {
		return 0, err
	}
	return resp.Payload[1], nil
}
