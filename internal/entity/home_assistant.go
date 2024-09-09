package entity

type HALightConfig struct {
	Name              string   `json:"name"`
	Icon              string   `json:"icon"`
	StateTopic        string   `json:"state_topic"`
	CommandTopic      string   `json:"command_topic"`
	AvailabilityTopic string   `json:"availability_topic"`
	UniqueID          string   `json:"unique_id"`
	Brightness        bool     `json:"brightness"`
	Effect            bool     `json:"effect"`
	Schema            string   `json:"schema"`
	EffectList        []string `json:"effect_list"`
	Supported         []string `json:"supported_color_modes"`
	Device            struct {
		Name        string   `json:"name"`
		Identifiers []string `json:"identifiers"`
	} `json:"device"`
}
