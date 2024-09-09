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
	Identifiers []string `json:"identifiers,omitempty"`
	Name        string   `json:"name"`
	Software    string   `json:"sw_version,omitempty"`
	Hardware    string   `json:"hw_version,omitempty"`
	ViaDevice   string   `json:"via_device,omitempty"`
}

// LightConfig represents the light config
type LightConfig struct {
	Name                string       `json:"name"`
	Icon                string       `json:"icon,omitempty"`
	StateTopic          string       `json:"state_topic,omitempty"`
	CommandTopic        string       `json:"command_topic,omitempty"`
	AvailabilityTopic   string       `json:"availability_topic,omitempty"`
	UniqueID            string       `json:"unique_id"`
	Brightness          bool         `json:"brightness,omitempty"`
	Effect              bool         `json:"effect,omitempty"`
	Schema              SchemaType   `json:"schema,omitempty"`
	EffectList          []string     `json:"effect_list,omitempty"`
	SupportedColorModes []ColorMode  `json:"supported_color_modes,omitempty"`
	Device              DeviceConfig `json:"device,omitempty"`
}
