package timing

import (
	"time"
)

// Interval runs a function periodically
type Interval struct {
	Delay  time.Duration
	Handle func()
}

// NewInterval creates a new interval
func NewInterval(delay time.Duration, handle func()) *Interval {
	i := &Interval{
		Delay:  delay,
		Handle: handle,
	}
	return i
}

// Start starts interval loop. Loop should exit on close event
func (i *Interval) Start(stop <-chan struct{}) {
	go func() {
		i.Handle()
		for {
			select {
			case <-time.After(i.Delay):
				i.Handle()
			case <-stop:
				return
			}
		}
	}()
}
