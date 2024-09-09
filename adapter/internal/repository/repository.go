package repository

import (
	"pixie_adapter/internal/interfaces"
	"pixie_adapter/pkg/pixie"
)

type Repository struct {
	light  *LightRepository
	time   *TimeRepository
	system *SystemRepository
}

var _ interfaces.Repositories = (*Repository)(nil)

func New(conn *pixie.Connection) *Repository {
	return &Repository{
		light:  newLightRepository(conn),
		time:   newTimeRepository(conn),
		system: newSystemRepository(conn),
	}
}

func (r *Repository) Light() interfaces.LightRepository {
	return r.light
}

func (r *Repository) Time() interfaces.TimeRepository {
	return r.time
}

func (r *Repository) System() interfaces.SystemRepository {
	return r.system
}
