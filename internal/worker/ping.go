package worker

import (
	"pixie_adapter/internal/interfaces"
	"time"
)

// Ping worker checks if the pixie clock is connected
type Ping struct {
	repos interfaces.Repositories
}

var _ interfaces.Worker = (*Ping)(nil)

// NewPing creates a new Ping worker
func NewPing(repos interfaces.Repositories) *Ping {
	return &Ping{
		repos: repos,
	}
}

// Name of the worker
func (p *Ping) Name() string {
	return "ping"
}

// Interval between job runs
func (p *Ping) Interval() time.Duration {
	return 5 * time.Second
}

// Run the worker
func (p *Ping) Run() {
	p.repos.System().IsConnected()
}
