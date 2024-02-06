"""Support for Pixie Clock light"""
from datetime import timedelta
from logging import getLogger
from typing import (
    Callable,
    Optional,
    Final,
    Any,
)
from homeassistant.helpers.typing import (
    ConfigType,
    DiscoveryInfoType,
    HomeAssistantType,
)
from homeassistant.components.light import (
    ATTR_BRIGHTNESS,
    ATTR_HS_COLOR,
    ATTR_EFFECT,
    COLOR_MODE_HS,
    PLATFORM_SCHEMA,
    LightEntity,
)
import homeassistant.util.color as color_util
import homeassistant.helpers.config_validation as cv
import voluptuous as vol
from .const import CONF_ADDR
from .api import PixieAPI

_LOGGER = getLogger('PixieLight')

SCAN_INTERVAL = timedelta(seconds=10)

effects = [
    "Static",
    "Smooth",
    "Zoom",
]

# pylint: disable=unused-argument,too-many-instance-attributes

PLATFORM_SCHEMA = PLATFORM_SCHEMA.extend({
    vol.Required(CONF_ADDR): cv.string
})

async def async_setup_platform(
    hass: HomeAssistantType,
    config: ConfigType,
    async_add_entities: Callable,
    discovery_info: Optional[DiscoveryInfoType] = None,
) -> None:
    """Set up Pixie Clock light platform"""
    if CONF_ADDR in config:
        async_add_entities([
            PixieClock(hass, config[CONF_ADDR])
        ])

class PixieClock(LightEntity):
    """Pixie Clock light entity"""

    _attr_supported_color_modes: Final[set[str]] = {
        COLOR_MODE_HS,
    }
    _attr_color_mode = COLOR_MODE_HS
    _attr_brightness: int = 255
    _attr_is_on: bool = True
    _attr_available: bool = False
    _attr_effect_list = effects
    _attr_effect = effects[0]
    _attr_hs_color: tuple[float, float] = (0, 0)
    _api: PixieAPI
    _next_state = None

    def __init__(self, hass, addr: str):
        self._hass = hass
        self._api = PixieAPI(hass, addr)

    @property
    def unique_id(self) -> str:
        """Return the unique id of the device."""
        return "indicators"
    
    @property
    def icon(self) -> str:
        return "mdi:clock-digital"

    async def async_update(self) -> None:
        """Check if keyboard is available"""
        if self._next_state is not None:
            state = self._next_state
            self._next_state = None
        else:
            state = await self._api.get_state()
        if state is None:
            self._attr_available = False
        else:
            self._attr_available = True
            self._attr_hs_color = color_util.color_RGB_to_hs(*state["color"])
            self._attr_brightness = state["brightness"]
            self._attr_effect = effects[state["effect"]]
        self.async_write_ha_state()

    async def async_turn_on(self, **kwargs: Any) -> None:
        """Turn the light on."""
        color = self._attr_hs_color
        brightness = self._attr_brightness
        effect = self._attr_effect
        if ATTR_EFFECT in kwargs:
            effect = kwargs[ATTR_EFFECT]
        if ATTR_HS_COLOR in kwargs:
            color = kwargs[ATTR_HS_COLOR]
        if ATTR_BRIGHTNESS in kwargs:
            brightness = kwargs[ATTR_BRIGHTNESS]
        self._next_state = {
            "color": list(color_util.color_hs_to_RGB(*color)),
            "brightness": brightness,
            "effect": effects.index(effect)
        }
        self._attr_is_on = True
        self.async_write_ha_state()
        if await self._api.set_state(
            self._next_state["color"],
            self._next_state["brightness"],
            self._next_state["effect"]
        ):
            self._attr_available = True
            self._attr_is_on = True
        else:
            self._attr_available = False

    async def async_turn_off(self, **kwargs: Any) -> None:
        """Turn the light off."""
        if not await self._api.turn_off():
            self._attr_available = False
            return
        self._attr_is_on = False
        self.async_write_ha_state()
