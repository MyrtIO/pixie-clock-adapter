# Pixie Clock Adapter

This repository contains the server and Home Assistant integration that allows you to add [Pixie Clock](https://github.com/MyrtIO/pixie-clock).

## Usage

First of all, install a server on your computer that is connected to Pixie Clock.

```sh
make build
make install
```

Next, add the integration from the [custom_components](./custom_components/) folder to your Home Assistant configuration. In the configuration yaml add:

```yaml
light:
  - platform: pixie_clock
    address: '192.168.1.5' # Your server local IP address
```
