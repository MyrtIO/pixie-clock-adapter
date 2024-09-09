package mqtt

import (
	"encoding/json"
	"log"
	"pixie_adapter/internal/entity"
	"pixie_adapter/internal/interfaces"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const topicLightUpdate = "myrt/pixie/light/set"
const topicLightState = "myrt/pixie/light"
const topicLightAvailability = "myrt/pixie/light/available"
const topicLightConfig = "homeassistant/light/pixie_clock_light/config"

type Handler struct {
	client mqtt.Client
	repos  interfaces.Repositories
}

func newHandler(repos interfaces.Repositories) Handler {
	return Handler{
		repos: repos,
	}
}

func (h *Handler) Router(c mqtt.Client) *Router {
	h.client = c
	router := newRouter(h.client)
	router.OnTopicUpdate(topicLightUpdate, h.HandleUpdateLightState)
	router.Report(h.HandleReportLightState, 10*time.Second)
	router.Report(h.HandleReportConfig, 10*time.Second)
	router.Report(h.HandleReportAvailability, 10*time.Second)

	return router
}

func (h *Handler) HandleReportConfig(client mqtt.Client) {
	config := entity.HALightConfig{
		Name:              "PixieClock",
		UniqueID:          "pixie_clock_light",
		Icon:              "mdi:clock-digital",
		Brightness:        true,
		Effect:            true,
		Schema:            "json",
		StateTopic:        topicLightState,
		CommandTopic:      topicLightUpdate,
		AvailabilityTopic: topicLightAvailability,
		EffectList:        []string{"static", "smooth", "zoom"},
		Supported:         []string{"rgb"},
		Device: struct {
			Name        string   `json:"name"`
			Identifiers []string `json:"identifiers"`
		}{
			Name: "PixieClock",
			Identifiers: []string{
				"pixie_clock",
			},
		},
	}

	msg, _ := json.Marshal(config)
	token := h.client.Publish(topicLightConfig, 0, false, msg)
	token.Wait()
}

func (h *Handler) HandleReportAvailability(client mqtt.Client) {
	_, err := h.repos.Light().GetPower()
	var token mqtt.Token
	if err != nil {
		token = h.client.Publish(topicLightAvailability, 0, false, "offline")
	} else {
		token = h.client.Publish(topicLightAvailability, 0, false, "online")
	}

	token.Wait()
}

func (h *Handler) HandleUpdateLightState(client mqtt.Client, msg mqtt.Message) {
	var state entity.LightState
	err := json.Unmarshal(msg.Payload(), &state)
	if err != nil {
		log.Printf("Error parsing message: %s\n", err)
	}

	log.Printf("Received state update: %+v\n", state)

	hasChanges := false
	effect, _ := h.repos.Light().GetEffect()
	if effect != entity.LightEffectEmpty && effect != state.Effect {
		err = h.repos.Light().SetEffect(state.Effect)
		if err != nil {
			log.Printf("Error setting effect: %s\n", err)
		}
		hasChanges = true
	}
	brightness, _ := h.repos.Light().GetBrightness()
	if state.Brightness != 0 && brightness != state.Brightness {
		err = h.repos.Light().SetBrightness(state.Brightness)
		if err != nil {
			log.Printf("Error setting brightness: %s\n", err)
		}
		hasChanges = true
	}
	color, _ := h.repos.Light().GetColor()
	if state.Color != entity.ColorBlack && color != state.Color {
		err = h.repos.Light().SetColor(state.Color)
		if err != nil {
			log.Printf("Error setting color: %s\n", err)
		}
		hasChanges = true
	}
	isEnabled, _ := h.repos.Light().GetPower()
	if state.State != entity.LightPowerStateEmpty &&
		(isEnabled != (state.State == entity.LightPowerStateOn)) {
		err = h.repos.Light().SetPower(state.State == entity.LightPowerStateOn)
		if err != nil {
			log.Printf("Error setting power: %s\n", err)
		}
		hasChanges = true
	}

	if hasChanges {
		h.HandleReportLightState(client)
	}
}

func (h *Handler) HandleReportLightState(client mqtt.Client) {
	var state entity.LightState

	brightness, err := h.repos.Light().GetBrightness()
	if err != nil {
		log.Printf("Error getting brightness: %s\n", err)
		return
	}
	color, err := h.repos.Light().GetColor()
	if err != nil {
		log.Printf("Error getting color: %s\n", err)
		return
	}
	effect, err := h.repos.Light().GetEffect()
	if err != nil {
		log.Printf("Error getting effect: %s\n", err)
		return
	}
	isEnabled, err := h.repos.Light().GetPower()
	if err != nil {
		log.Printf("Error getting power: %s\n", err)
		return
	}
	if isEnabled {
		state.State = entity.LightPowerStateOn
	} else {
		state.State = entity.LightPowerStateOff
	}
	state.Brightness = brightness
	state.Color = color
	state.Effect = effect
	state.ColorMode = entity.LightColorModeRGB

	var bytes []byte
	bytes, err = json.Marshal(state)
	if err != nil {
		log.Printf("Error marshalling state: %s\n", err)
		return
	}
	client.Publish(topicLightState, 0, false, bytes)
}
