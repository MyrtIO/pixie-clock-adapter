package mqtt

import (
	"encoding/json"
	"log"
	"os"
	"pixie_adapter/internal/config"
	"pixie_adapter/internal/entity"
	"pixie_adapter/internal/interfaces"
	"pixie_adapter/pkg/homeassistant"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	topicLightStateUpdate  = "myrt/pixie/light/set"
	topicLightState        = "myrt/pixie/light"
	topicLightAvailability = "myrt/pixie/light/available"
	topicLightConfig       = "homeassistant/light/pixie_clock_light/config"

	stateReportInterval        = 50 * time.Second
	availabilityReportInterval = 60 * time.Second
	availabilityCheckInterval  = 5 * time.Second
	configReportInterval       = 120 * time.Second
)

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}
	return hostname
}

var entityConfig = homeassistant.LightConfig{
	Name:                "light",
	UniqueID:            "pixie_clock_light",
	Icon:                "mdi:clock-digital",
	Brightness:          true,
	Effect:              true,
	Schema:              homeassistant.SchemaTypeJSON,
	StateTopic:          topicLightState,
	CommandTopic:        topicLightStateUpdate,
	AvailabilityTopic:   topicLightAvailability,
	SupportedColorModes: []homeassistant.ColorMode{homeassistant.ColorModeRGB},
	EffectList:          []string{"static", "smooth", "zoom"},
	Device: homeassistant.DeviceConfig{
		Identifiers: []string{"pixie_clock_light"},
		Name:        "PixieClock",
		Software:    config.Version,
		Hardware:    "v1",
		ViaDevice:   getHostname(),
	},
}

// Handler handles MQTT messages
type Handler struct {
	client                 mqtt.Client
	repos                  interfaces.Repositories
	nextAvailabilityReport time.Time
	isConnected            bool
}

func newHandler(repos interfaces.Repositories) Handler {
	return Handler{
		repos: repos,
	}
}

// Router creates a new handled router
func (h *Handler) Router(c mqtt.Client) *Router {
	h.client = c
	router := newRouter(h.client)
	router.OnTopicUpdate(topicLightStateUpdate, h.HandleUpdateLightState)
	router.Report(h.HandleReportLightState, stateReportInterval)
	router.Report(h.HandleReportConfig, configReportInterval)
	router.Report(h.HandleReportAvailability, availabilityCheckInterval)
	return router
}

// HandleReportConfig reports the config
func (h *Handler) HandleReportConfig(client mqtt.Client) {
	msg, _ := json.Marshal(entityConfig)
	h.safePublish(client, topicLightConfig, msg)
}

// HandleReportAvailability reports the availability
func (h *Handler) HandleReportAvailability(client mqtt.Client) {
	var message string
	isConnected := h.repos.System().IsConnected()
	if isConnected == h.isConnected && time.Now().Before(h.nextAvailabilityReport) {
		return
	}

	h.isConnected = isConnected
	if isConnected {
		message = "online"
	} else {
		message = "offline"
	}
	h.safePublish(client, topicLightAvailability, message)
	h.nextAvailabilityReport = time.Now().Add(availabilityReportInterval)
}

// HandleUpdateLightState handles a light state update request
func (h *Handler) HandleUpdateLightState(client mqtt.Client, msg mqtt.Message) {
	var state entity.LightState
	err := json.Unmarshal(msg.Payload(), &state)
	if err != nil {
		log.Printf("Error parsing message: %s\n", err)
	}

	hasChanges, err := h.repos.Light().SetState(state)
	if err != nil {
		log.Printf("Error setting state: %s\n", err)
		return
	}
	if hasChanges {
		h.HandleReportLightState(client)
	}
}

// HandleReportLightState reports the light state
func (h *Handler) HandleReportLightState(client mqtt.Client) {
	if !h.repos.System().IsConnected() {
		return
	}
	state, err := h.repos.Light().GetState()
	if err != nil {
		log.Printf("Error getting state: %s\n", err)
		return
	}

	var data []byte
	data, err = json.Marshal(state)
	if err != nil {
		log.Printf("Error marshalling state: %s\n", err)
		return
	}
	h.safePublish(client, topicLightState, data)
}

func (h *Handler) safePublish(client mqtt.Client, topic string, payload interface{}) {
	token := client.Publish(topic, 0, false, payload)
	token.Wait()
	if token.Error() != nil {
		log.Printf("Error publishing: %s\n", token.Error())
	}
}
