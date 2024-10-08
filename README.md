# Pixie Clock Adapter [![Quality assurance](https://github.com/MyrtIO/pixie-clock-adapter/actions/workflows/quality-assurance.yaml/badge.svg)](https://github.com/MyrtIO/pixie-clock-adapter/actions/workflows/quality-assurance.yaml)

This repository contains a service that allows you to integrate [Pixie Clock](https://github.com/MyrtIO/pixie-clock) into Home Assistant via MQTT.

## Usage

First of all, install an adapter on your computer that is connected to Pixie Clock.

```sh
make build
make install
```

Next, create configuration file at `$HOME/.config/pixie-adapter/config.yaml`.

Configuration file should look like this:

```yaml
mqtt:
  host: "your.mqtt.server"
  port: 1883
  client_id: "PixieClock"
  username: ""
  password: ""
serial:
  port: "/dev/cu.wchusbserial114320" # omit to auto-detect
  baud_rate: 28800
```

After that, you can start the background service:

```sh
pixie-adapter start
```

## Topics

Service provides the following topics:

- `homeassistant/light/pixie_clock_light/config` — for Home Assistant discovery.
- `myrt/pixie/light/set` — for setting light state.
- `myrt/pixie/light` — for getting light state.
- `myrt/pixie/light/available` — for checking light availability.
