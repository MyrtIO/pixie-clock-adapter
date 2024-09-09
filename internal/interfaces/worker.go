package interfaces

import (
	"time"
)

// Worker represents a job that will be run periodically in the background.
type Worker interface {
	Name() string
	Interval() time.Duration
	Run()
}
