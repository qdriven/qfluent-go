package qtime

import "time"

func DiffNano(startTime time.Time) (diff int64) {
	diff = int64(time.Since(startTime))
	return
}
