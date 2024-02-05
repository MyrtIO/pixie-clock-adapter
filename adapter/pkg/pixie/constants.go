package pixie

const (
	FeatureSystem     byte = 0
	FeatureClock      byte = 1
	FeatureIndicators byte = 2
	FeatureStopwatch  byte = 3
)

const (
	// System actions
	ActionSystemPing    byte = 0
	ActionSystemGetName byte = 1
	// Clock actions
	ActionClockActivate byte = 0
	ActionClockSetTime  byte = 1
	// Indicators actions
	ActionIndicatorsSetColor      byte = 0
	ActionIndicatorsSetBrightness byte = 1
	ActionIndicatorsGetColor      byte = 2
	ActionIndicatorsGetBrightness byte = 3
	ActionIndicatorsSetPower      byte = 4
	ActionIndicatorsGetPower      byte = 5
	// Stopwatch actions
	ActionStopwatchActivate byte = 0
)
