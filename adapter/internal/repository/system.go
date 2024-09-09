package repository

import (
	"pixie_adapter/internal/interfaces"
	"pixie_adapter/pkg/pixie"
)

type SystemRepository struct {
	conn *pixie.Connection
}

var _ interfaces.SystemRepository = (*SystemRepository)(nil)

func newSystemRepository(conn *pixie.Connection) *SystemRepository {
	return &SystemRepository{
		conn: conn,
	}
}

func (s *SystemRepository) IsConnected() bool {
	tx, err := s.conn.Get()
	if err != nil || tx == nil {
		return false
	}
	return pixie.Ping(tx)
}
