package state

import (
	"errors"
	"fmt"
)

type proc_t struct {
	name  string
	utime [2]int64 // user mode jiffies
	stime [2]int64 // kernel mode jiffies
}

// Adds a new process to the proctable.
func (pm PIDMap) NewProc(name string, pid uint64, utime, stime int64) {
	pm[pid] = proc_t{
		name:  name,
		utime: [2]int64{-1, utime},
		stime: [2]int64{-1, stime},
	}
}

func (pm PIDMap) UpdateProc(pid uint64, utime, stime int64) error {
	if proc, ok := pm[pid]; ok {
		proc.utime[0] = proc.utime[1]
		proc.utime[1] = utime

		proc.stime[0] = proc.stime[1]
		proc.stime[1] = stime

	} else {
		return errors.New(
			fmt.Sprintf(
				"pid '%v' was supposed to be alive and in memory, but was not found",
				pid,
			),
		)
	}

	return nil
}

func (pt proc_t) Utime() int {
	return int(pt.utime[1] - pt.utime[0])
}

func (pt proc_t) Stime() int {
	return int(pt.stime[1] - pt.stime[0])
}

func (pt proc_t) Name() string {
	return pt.name
}
