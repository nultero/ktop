package state

import (
	"fmt"
	"ktop/styles"
	"time"
)

type log []string

type PIDMap map[uint64]proc_t

type State struct {
	Cursor View     // The active component within the focused quad
	Quad   Quadrant // The focused quadrant within the frame.

	PollRate time.Duration

	Cpu       cpu_t
	Mem       mem_t
	MaxStamps int // how many cpu / mem stamps to keep
	PidMap    PIDMap
	Top       ProcTab // the table of processes
	Time      time_t  // used in cpu calcs
	Total     float64 // total jiffies in the proctab

	ColorTheme styles.Theme

	NeedsRedraw bool

	Log log
}

func DefaultState() State {
	return State{
		Cursor: 0, // corresponds to CPU in ../viewmap.go
		Quad:   QuadTopRight,

		Cpu:       defaultCpu_t(),
		Mem:       defaultMem_t(),
		MaxStamps: 120,
		PidMap:    PIDMap{},
		Top:       ProcTab{},
		Time:      defaultTime_t(),
		Total:     0.0,

		// this is just the default;
		// intended to be overwritten by a config
		ColorTheme: styles.CrystalTheme(),

		NeedsRedraw: false,

		Log: log{},
	}
}

// Calls fmt.Println on every line in the log.
//
// Useful for outputting captured, nonfatal errors
// such as those that might come from reading
// from deeper chunks of procfs.
func (l log) Dump() {
	for _, ln := range l {
		fmt.Println(ln)
	}
}
