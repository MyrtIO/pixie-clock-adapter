package worker

import (
	"log"
	"pixie_adapter/internal/interfaces"
	"pixie_adapter/pkg/timing"
)

type Runner struct {
	workers []interfaces.Worker
}

func NewRunner(workers ...interfaces.Worker) *Runner {
	return &Runner{
		workers: workers,
	}
}

func (r *Runner) Run(stop <-chan struct{}) {
	for _, w := range r.workers {
		interval := timing.NewInterval(w.Interval(), w.Run)
		log.Printf("Starting %s worker\n", w.Name())
		interval.Start(stop)
	}
}
