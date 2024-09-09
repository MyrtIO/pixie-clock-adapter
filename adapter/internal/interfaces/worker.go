package interfaces

import (
	"time"
)

type Worker interface {
	Name() string
	Interval() time.Duration
	Run()
}
