package worker

import (
	"log"
	"pixie_adapter/internal/interfaces"
	"time"
)

type TimeSync struct {
	repos interfaces.Repositories
}

var _ interfaces.Worker = (*TimeSync)(nil)

func NewTimeSync(repos interfaces.Repositories) *TimeSync {
	return &TimeSync{
		repos: repos,
	}
}

func (t *TimeSync) Name() string {
	return "time_sync"
}

func (t *TimeSync) Interval() time.Duration {
	return 60 * time.Second
}

func (t *TimeSync) Run() {
	if !t.repos.System().IsConnected() {
		return
	}
	err := t.repos.Time().Set(time.Now())
	if err != nil {
		log.Printf("Error updating time: %s\n", err)
	}
}
