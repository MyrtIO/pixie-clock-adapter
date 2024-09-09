package timesync

import (
	"fmt"
	"pixie_adapter/pkg/pixie"
	"pixie_adapter/pkg/timing"
	"time"
)

const syncInterval = 120 * time.Second

// New creates new time syncer
func New(conn *pixie.Connection) *timing.Interval {
	sync := func() {
		tx, _ := conn.Get()
		if tx == nil {
			return
		}
		_, err := pixie.SetTime(tx, time.Now())
		if err != nil {
			fmt.Printf("Error syncing time: %v\n", err)
		}
	}
	interval := timing.NewInterval(syncInterval, sync)
	return interval
}
