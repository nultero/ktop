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

/*
	Various output methods below
*/

const multiDigit float64 = 10.0

// Last CPU percent.
func (s *State) LCpuPC() float64 {
	return s.Cpu.Stamps[len(s.Cpu.Stamps)-1]
}

// Last CPU percent Sprintf'd to a string.
func (s *State) LCpuPCStr() string {
	c := s.Cpu.Stamps[len(s.Cpu.Stamps)-1]
	if c < multiDigit {
		return fmt.Sprintf(" %.2f", c)
	}

	return fmt.Sprintf("%.2f", c)
}

// Last memory stamp percent.
func (s *State) LMemPC() float64 {
	return s.Mem.Stamps[len(s.Mem.Stamps)-1]
}

// Last memory stamp percent Sprintf'd to a string.
func (s *State) LMemPCStr() string {
	m := s.Mem.Stamps[len(s.Mem.Stamps)-1]
	if m < multiDigit {
		return fmt.Sprintf(" %.2f", m)
	}

	return fmt.Sprintf("%.2f", m)
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
