package repository

import (
	"pixie_adapter/internal/interfaces"
)

// Repository provides access to the device features
type Repository struct {
	light  *LightRepository
	time   *TimeRepository
	system *SystemRepository
}

var _ interfaces.Repositories = (*Repository)(nil)

// New creates a new repository
func New(conn interfaces.TransportProvider) *Repository {
	return &Repository{
		light:  newLightRepository(conn),
		time:   newTimeRepository(conn),
		system: newSystemRepository(conn),
	}
}

// Light provides access to the light
func (r *Repository) Light() interfaces.LightRepository {
	return r.light
}

// Time provides access to the time
func (r *Repository) Time() interfaces.TimeRepository {
	return r.time
}

// System provides access to the system
func (r *Repository) System() interfaces.SystemRepository {
	return r.system
}
