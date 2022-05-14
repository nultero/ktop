package state

import "time"

type time_t struct {
	pstart   time.Time
	PollRate time.Duration
}

func defaultTime_t() time_t {
	return time_t{
		pstart:   time.Now(),
		PollRate: 400 * time.Millisecond,
	}
}
