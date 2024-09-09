package repository

import (
	"pixie_adapter/internal/interfaces"
	"pixie_adapter/pkg/pixie"
	"time"
)

// TimeRepository provides access to the time
type TimeRepository struct {
	conn *pixie.Connection
}

var _ interfaces.TimeRepository = (*TimeRepository)(nil)

func newTimeRepository(conn *pixie.Connection) *TimeRepository {
	return &TimeRepository{
		conn: conn,
	}
}

// Set sets the time of the pixie clock
func (t *TimeRepository) Set(nextTime time.Time) error {
	tx, err := t.conn.Get()
	if err != nil {
		return err
	}
	_, err = pixie.SetTime(tx, nextTime)
	return err
}
