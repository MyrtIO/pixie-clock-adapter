package repository

import (
	"pixie_adapter/internal/entity"
	"pixie_adapter/internal/interfaces"
	"pixie_adapter/pkg/pixie"
)

type LightRepository struct {
	conn *pixie.Connection
}

var _ interfaces.LightRepository = (*LightRepository)(nil)

func newLightRepository(conn *pixie.Connection) *LightRepository {
	return &LightRepository{
		conn: conn,
	}
}

func (c *LightRepository) SetColor(color entity.RGBColor) error {
	tx, err := c.conn.Get()
	if err != nil {
		return err
	}
	_, err = pixie.SetColor(tx, color.R, color.G, color.B)
	return err
}

func (c *LightRepository) GetColor() (color entity.RGBColor, err error) {
	tx, err := c.conn.Get()
	if err != nil {
		return
	}
	values, err := pixie.GetColor(tx)
	if err != nil {
		return
	}

	return entity.RGBColor{
		R: values[0],
		G: values[1],
		B: values[2],
	}, nil
}

func (c *LightRepository) SetBrightness(brightness uint8) error {
	tx, err := c.conn.Get()
	if err != nil {
		return err
	}
	_, err = pixie.SetBrightness(tx, brightness)
	return err
}

func (c *LightRepository) GetBrightness() (uint8, error) {
	tx, err := c.conn.Get()
	if err != nil {
		return 0, err
	}
	return pixie.GetBrightness(tx)
}

func (c *LightRepository) SetPower(enabled bool) error {
	tx, err := c.conn.Get()
	if err != nil {
		return err
	}
	_, err = pixie.SetPower(tx, enabled)
	return err
}

func (c *LightRepository) GetPower() (bool, error) {
	tx, err := c.conn.Get()
	if err != nil {
		return false, err
	}
	return pixie.GetPower(tx)
}

func (c *LightRepository) SetEffect(effect entity.LightEffect) error {
	tx, err := c.conn.Get()
	if err != nil {
		return err
	}
	_, err = pixie.SetEffect(tx, effect.Code())
	return err
}

func (c *LightRepository) GetEffect() (entity.LightEffect, error) {
	tx, err := c.conn.Get()
	if err != nil {
		return entity.LightEffectStatic, err
	}
	effect, err := pixie.GetEffect(tx)
	if err != nil {
		return entity.LightEffectStatic, err
	}
	return entity.LightEffectFromCode(effect), nil
}
