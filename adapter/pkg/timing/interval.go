package timing

import (
	"time"
)

type Interval struct {
	Delay  time.Duration
	Handle func()
}

// New creates new interval. Receives delay and handle
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
