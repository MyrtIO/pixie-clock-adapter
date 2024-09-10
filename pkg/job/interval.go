package job

import (
	"time"
)

// Interval runs a function periodically
type Interval struct {
	Period   time.Duration
	Handle   func()
	RunFirst bool
}

// NewInterval creates a new interval
func NewInterval(handle func(), period time.Duration) *Interval {
	i := &Interval{
		Period:   period,
		Handle:   handle,
		RunFirst: true,
	}
	return i
}

// Start starts interval loop. Loop should exit on close event
func (i *Interval) Start(stop <-chan struct{}) {
	go func() {
		if i.RunFirst {
			i.Handle()
		}
		for {
			select {
			case <-time.After(i.Period):
				i.Handle()
			case <-stop:
				return
			}
		}
	}()
}
