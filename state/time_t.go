package state

import "time"

type time_t struct {
	pstart time.Time
}

func defaultTime_t() time_t {
	return time_t{
		pstart: time.Now(),
	}
}
