package mqtt

import (
	"encoding/json"
	"log"
	"pixie_adapter/internal/entity"
	"pixie_adapter/internal/interfaces"
	"pixie_adapter/pkg/homeassistant"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const topicLightUpdate = "myrt/pixie/light/set"
const topicLightState = "myrt/pixie/light"
const topicLightAvailability = "myrt/pixie/light/available"
const topicLightConfig = "homeassistant/light/pixie_clock_light/config"

var entityConfig = homeassistant.LightConfig{
	Name:              "light",
	UniqueID:          "pixie_clock_light",
	Icon:              "mdi:clock-digital",
	Brightness:        true,
	Effect:            true,
	Schema:            homeassistant.SchemaTypeJSON,
	StateTopic:        topicLightState,
	CommandTopic:      topicLightUpdate,
	AvailabilityTopic: topicLightAvailability,
	EffectList:        []string{"static", "smooth", "zoom"},
	Device: homeassistant.DeviceConfig{
		Identifiers: []string{"pixie_clock_light"},
		Name:        "PixieClock",
	},
}

// Handler handles MQTT messages
type Handler struct {
	client mqtt.Client
	repos  interfaces.Repositories
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
	router.OnTopicUpdate(topicLightUpdate, h.HandleUpdateLightState)
	router.Report(h.HandleReportLightState, 10*time.Second)
	router.Report(h.HandleReportConfig, 10*time.Second)
	router.Report(h.HandleReportAvailability, 10*time.Second)

	return router
}

// HandleReportConfig reports the config
func (h *Handler) HandleReportConfig(client mqtt.Client) {
	msg, _ := json.Marshal(entityConfig)
	token := client.Publish(topicLightConfig, 0, false, msg)
	token.Wait()
	if token.Error() != nil {
		log.Printf("Error publishing config: %s\n", token.Error())
	}
}

// HandleReportAvailability reports the availability
func (h *Handler) HandleReportAvailability(client mqtt.Client) {
	var token mqtt.Token
	var message string
	if h.repos.System().IsConnected() {
		message = "online"
	} else {
		message = "offline"
	}
	token = client.Publish(topicLightAvailability, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		log.Printf("Error publishing availability: %s\n", token.Error())
	}
}

// HandleUpdateLightState handles a light state update request
func (h *Handler) HandleUpdateLightState(client mqtt.Client, msg mqtt.Message) {
	var state entity.LightState
	err := json.Unmarshal(msg.Payload(), &state)
	if err != nil {
		log.Printf("Error parsing message: %s\n", err)
	}

	log.Printf("Received state update: %+v\n", state)
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
	state, err := h.repos.Light().GetState()
	if err != nil {
		log.Printf("Error getting state: %s\n", err)
		return
	}

	var bytes []byte
	bytes, err = json.Marshal(state)
	if err != nil {
		log.Printf("Error marshalling state: %s\n", err)
		return
	}
	token := client.Publish(topicLightState, 0, false, bytes)
	token.Wait()
	if token.Error() != nil {
		log.Printf("Error publishing state: %s\n", token.Error())
	}
}
