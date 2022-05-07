package state

import (
	"fmt"
	"ktop/styles"
	"time"
)

type State struct {
	Cursor View     // The active component within the focused quad
	quad   Quadrant // The focused quadrant within the frame.

	PollRate  time.Duration
	LCI       int // last CPU idle %, used for PollCPU
	CpuSum    int
	CpuStamps []float64
	RamStamps []float64
	MaxStamps int // how many cpu / mem stamps to keep

	ColorTheme styles.Theme

	NeedsRedraw bool
}

func DefaultState() State {
	return State{
		Cursor: 0, // corresponds to CPU in ../viewmap.go
		quad:   QuadTopRight,

		PollRate:  200 * time.Millisecond,
		LCI:       0,
		CpuSum:    0,
		CpuStamps: []float64{},
		RamStamps: []float64{},
		MaxStamps: 120,

		ColorTheme: styles.CrystalTheme(),

		NeedsRedraw: false,
	}
}

/*
	Various output methods below
*/

const multiDigit float64 = 10.0

// Last CPU percent.
func (s *State) LCpuPC() float64 {
	return s.CpuStamps[len(s.CpuStamps)-1]
}

// Last CPU percent Sprintf'd to a string.
func (s *State) LCpuPCStr() string {
	c := s.CpuStamps[len(s.CpuStamps)-1]
	if c < multiDigit {
		return fmt.Sprintf(" %.2f", c)
	}

	return fmt.Sprintf("%.2f", c)
}

// Last memory stamp percent.
func (s *State) LMemPC() float64 {
	return s.RamStamps[len(s.RamStamps)-1]
}

// Last memory stamp percent Sprintf'd to a string.
func (s *State) LMemPCStr() string {
	m := s.RamStamps[len(s.RamStamps)-1]
	if m < multiDigit {
		return fmt.Sprintf(" %.2f", m)
	}

	return fmt.Sprintf("%.2f", m)
}
