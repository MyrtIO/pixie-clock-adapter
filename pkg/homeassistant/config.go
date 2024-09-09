package homeassistant

// ColorMode represents the supported color modes
type ColorMode string

const (
	// ColorModeRGB represents the RGB color mode
	ColorModeRGB ColorMode = "rgb"
)

// SchemaType represents the type of the schema
type SchemaType string

const (
	// SchemaTypeJSON represents the JSON schema
	SchemaTypeJSON SchemaType = "json"
)

// DeviceConfig represents the device config
type DeviceConfig struct {
	Identifiers []string `json:"identifiers"`
	Name        string   `json:"name"`
}

// LightConfig represents the light config
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
