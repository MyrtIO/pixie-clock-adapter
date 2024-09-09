package worker

import (
	"log"
	"pixie_adapter/internal/interfaces"
	"time"
)

// TimeSync syncs the time of the pixie clock
type TimeSync struct {
	repos interfaces.Repositories
}

var _ interfaces.Worker = (*TimeSync)(nil)

// NewTimeSync creates a new TimeSync worker
func NewTimeSync(repos interfaces.Repositories) *TimeSync {
	return &TimeSync{
		repos: repos,
	}
}

// Name of the worker
func (t *TimeSync) Name() string {
	return "time_sync"
}

// Interval between job runs
func (t *TimeSync) Interval() time.Duration {
	return 60 * time.Second
}

// Run the worker
func (t *TimeSync) Run() {
	if !t.repos.System().IsConnected() {
		return
	}
	err := t.repos.Time().Set(time.Now())
	if err != nil {
		log.Printf("Error updating time: %s\n", err)
	}
}
