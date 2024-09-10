package job_test

import (
	"pixie_adapter/pkg/job"
	"testing"
	"time"
)

func TestInterval(t *testing.T) {
	t.Parallel()
	counter := 0
	callback := func() {
		counter++
	}
	stop := make(chan struct{})

	job := job.NewInterval(callback, 10*time.Millisecond)
	job.Start(stop)
	time.Sleep(22 * time.Millisecond)

	stop <- struct{}{}
	if counter != 3 {
		t.Errorf("Expected 3, got %d", counter)
	}
}

func TestIntervalWithoutPreRun(t *testing.T) {
	t.Parallel()
	counter := 0
	callback := func() {
		counter++
	}
	stop := make(chan struct{})

	job := job.NewInterval(callback, 10*time.Millisecond)
	job.RunFirst = false
	job.Start(stop)
	time.Sleep(22 * time.Millisecond)

	stop <- struct{}{}
	if counter != 2 {
		t.Errorf("Expected 2, got %d", counter)
	}
}
