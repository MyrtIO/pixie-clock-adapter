package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// AppName represents app name.
const AppName = "pixie-adapter"

// PackageName represents app package name.
const PackageName = "co.myrt.pixie_adapter"

// Version represents current app version.
var Version = "snapshot"

// Config represents the application configuration.
type Config struct {
	MQTT   `yaml:"mqtt"`
	Serial `yaml:"serial"`
}

// MQTT represents the MQTT configuration.
type MQTT struct {
	Host     string `env:"MQTT_HOST" yaml:"host"`
	Port     int    `env:"MQTT_PORT" yaml:"port"`
	Username string `env:"MQTT_USERNAME" yaml:"username"`
	Password string `env:"MQTT_PASSWORD" yaml:"password"`
	ClientID string `env:"MQTT_CLIENT_ID" yaml:"client_id"`
}

// Serial represents the serial configuration.
type Serial struct {
	Port     string `env:"SERIAL_PORT" yaml:"port"`
	BaudRate int    `env:"SERIAL_BAUD_RATE" yaml:"baud_rate"`
}

// New reads configuration from file and environment variables.
func New(path string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("failed to load env from file: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// MQTTServerURL returns the MQTT server URL.
func (c *MQTT) MQTTServerURL() string {
	return "tcp://" + c.Host + ":" + fmt.Sprint(c.Port)
}
