# Pixie Clock Adapter

This repository contains a service that allows you to integrate [Pixie Clock](https://github.com/MyrtIO/pixie-clock) into Home Assistant via MQTT.

## Usage

First of all, install a server on your computer that is connected to Pixie Clock.

```sh
make build
make install
```

Next, create configuration file at `$HOME/.config/pixie-adapter/config.yaml`.

Configuration file should look like this:

```yaml
mqtt:
  host: "your.mqtt.server"
  port: "1883"
  client_id: "PixieClock"
  username: ""
  password: ""
serial:
  port: "" # omit to auto-detect
  baud_rate: 28800
```

After that, you can start the background service:

```sh
pixie-adapter start
```
