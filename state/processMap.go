package state

import (
	"errors"
	"fmt"
)

type procMap map[uint64]proc_t

type sortedProcMap struct {
	Map  map[float64]string
	Keys []float64
}

type proc_t struct {
	cpuPc float64
	name  string
	utime [2]int64 // user mode jiffies
	stime [2]int64 // kernel mode jiffies
}

func (pt proc_t) Prev() int64 {
	return pt.utime[0] + pt.stime[0]
}

func (pt proc_t) Cur() int64 {
	return pt.utime[1] + pt.stime[1]
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

func (pt proc_t) CpuPc() float64 {
	return pt.cpuPc
}

// Adds a new process to the proctable.
func (pm procMap) NewProc(
	name string, pid uint64,
	utime, stime int64,
	cpuSum, cpuLast float64,
) {

	slicePc := float64(utime+stime) / cpuSum
	cpuPc := cpuLast * slicePc

	pm[pid] = proc_t{
		cpuPc: cpuPc,
		name:  name,
		utime: [2]int64{-1, utime},
		stime: [2]int64{-1, stime},
	}
}

func (pm procMap) UpdateProc(
	pid uint64,
	utime, stime int64,
	cpuSum, cpuLast float64,
) error {

	var proc proc_t
	if procObj, ok := pm[pid]; ok {
		proc = procObj
	} else {
		return errors.New(
			fmt.Sprintf(
				"pid '%v' was supposed to be alive and in memory, but was not found",
				pid,
			),
		)
	}

	proc.utime[0] = proc.utime[1]
	proc.utime[1] = utime

	proc.stime[0] = proc.stime[1]
	proc.stime[1] = stime

	diffU := utime - proc.utime[0]
	diffS := stime - proc.stime[0]

	slicePc := float64(diffU+diffS) / cpuSum
	proc.cpuPc = slicePc * cpuLast

	pm[pid] = proc
	return nil
}
