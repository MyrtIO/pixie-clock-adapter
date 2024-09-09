package homeassistant

type ColorMode string

const (
	ColorModeRGB ColorMode = "rgb"
)

type SchemaType string

const (
	SchemaTypeJSON SchemaType = "json"
)

type DeviceConfig struct {
	Identifiers []string `json:"identifiers"`
	Name        string   `json:"name"`
}

type LightConfig struct {
	Name                string       `json:"name"`
	Icon                string       `json:"icon"`
	StateTopic          string       `json:"state_topic"`
	CommandTopic        string       `json:"command_topic"`
	AvailabilityTopic   string       `json:"availability_topic"`
	UniqueID            string       `json:"unique_id"`
	Brightness          bool         `json:"brightness"`
	Effect              bool         `json:"effect"`
	Schema              SchemaType   `json:"schema"`
	EffectList          []string     `json:"effect_list"`
	SupportedColorModes []ColorMode  `json:"supported_color_modes"`
	Device              DeviceConfig `json:"device"`
}
