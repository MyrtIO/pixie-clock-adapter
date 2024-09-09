package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	MQTT   `yaml:"mqtt"`
	Serial `yaml:"serial"`
}

type MQTT struct {
	Host     string `env:"MQTT_HOST" yaml:"host"`
	Port     string `env:"MQTT_PORT" yaml:"port"`
	Username string `env:"MQTT_USERNAME" yaml:"username"`
	Password string `env:"MQTT_PASSWORD" yaml:"password"`
	ClientID string `env:"MQTT_CLIENT_ID" yaml:"client_id"`
}

type Serial struct {
	PortPath string `env:"SERIAL_PORT" yaml:"port_path"`
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

func (c *MQTT) MQTTServerURL() string {
	return "tcp://" + c.Host + ":" + fmt.Sprint(c.Port)
}
