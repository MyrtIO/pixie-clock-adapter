package worker

import (
	"log"
	"pixie_adapter/internal/interfaces"
	"pixie_adapter/pkg/timing"
)

// Runner runs the workers
type Runner struct {
	workers []interfaces.Worker
}

// NewRunner creates a new runner
func NewRunner(workers ...interfaces.Worker) *Runner {
	return &Runner{
		workers: workers,
	}
}

// Run workers
func (r *Runner) Run(stop <-chan struct{}) {
	for _, w := range r.workers {
		interval := timing.NewInterval(w.Interval(), w.Run)
		log.Printf("Starting %s worker\n", w.Name())
		interval.Start(stop)
	}
}
