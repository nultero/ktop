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
	quad   Quadrant // The focused quadrant within the frame.

	PollRate time.Duration

	Cpu       cpu_t
	Mem       mem_t
	MaxStamps int // how many cpu / mem stamps to keep
	PidMap    PIDMap

	ColorTheme styles.Theme

	NeedsRedraw bool

	Log log
}

func DefaultState() State {
	return State{
		Cursor: 0, // corresponds to CPU in ../viewmap.go
		quad:   QuadTopRight,

		PollRate: 200 * time.Millisecond,

		Cpu:       defaultCpu_t(),
		Mem:       defaultMem_t(),
		MaxStamps: 120,
		PidMap:    PIDMap{},

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
