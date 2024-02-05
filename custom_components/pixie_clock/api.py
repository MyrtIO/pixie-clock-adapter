"""Pixie Clock adapter API"""
import aiohttp
from async_timeout import timeout
from logging import getLogger
from homeassistant.helpers.typing import HomeAssistantType
from homeassistant.helpers.aiohttp_client import async_get_clientsession

from .const import PORT

_LOGGER = getLogger('PixieAPI')

class PixieAPI:
    _addr: str
    _hass: HomeAssistantType

    def __init__(self, hass: HomeAssistantType, addr: str):
        self._addr = f"http://{addr}:{PORT}"
        self._hass = hass

    def _get_session(self) -> aiohttp.ClientSession:
        return async_get_clientsession(self._hass)
    
    async def get_state(self):
        try:
            session = self._get_session()
            async with timeout(2):
                resp = await session.get(f"{self._addr}/")
                return await resp.json()
        except Exception: # pylint: disable=broad-except
            return None
    
    async def set_state(self, color: list[int], brightness: int):
        try:
            session = self._get_session()
            resp = await session.put(f"{self._addr}/", json={
                "color": color,
                "brightness": brightness,
            })
            _LOGGER.warning(resp.status)
            return resp.status == 200
        except Exception: # pylint: disable=broad-except
            return False
    
    async def turn_off(self):
        try:
            session = self._get_session()
            resp = await session.post(f"{self._addr}/disable")
            return resp.status == 200
        except Exception: # pylint: disable=broad-except
            return False
