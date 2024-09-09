package worker

import (
	"pixie_adapter/internal/interfaces"
	"time"
)

type Ping struct {
	repos interfaces.Repositories
}

var _ interfaces.Worker = (*Ping)(nil)

func NewPing(repos interfaces.Repositories) *Ping {
	return &Ping{
		repos: repos,
	}
}

func (p *Ping) Name() string {
	return "ping"
}

func (p *Ping) Interval() time.Duration {
	return 5 * time.Second
}

func (p *Ping) Run() {
	p.repos.System().IsConnected()
}
