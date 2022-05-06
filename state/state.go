package state

import (
	"fmt"
	"ktop/styles"
	"time"
)

type View uint8

type State struct {
	FocusedComp View // The active component within the focused quad
	FocusedQuad Quadrant

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
		FocusedComp: 1, // corresponds to CPU in ../viewmap.go
		FocusedQuad: QuadTopRight,
		PollRate:    500 * time.Millisecond,
		LCI:         0,
		CpuSum:      0,
		CpuStamps:   []float64{},
		RamStamps:   []float64{},
		MaxStamps:   120,

		ColorTheme: styles.CrystalTheme(),

		NeedsRedraw: false,
	}
}

const multiDigit float64 = 10.0

func (s *State) LastCpuPC() float64 {
	return s.CpuStamps[len(s.CpuStamps)-1]
}

func (s *State) LastCpuPCString() string {
	c := s.CpuStamps[len(s.CpuStamps)-1]
	if c < multiDigit {
		return fmt.Sprintf(" %.2f", c)
	}

	return fmt.Sprintf("%.2f", c)
}

func (s *State) LastMemPC() float64 {
	return s.RamStamps[len(s.RamStamps)-1]
}

func (s *State) LastMemPCString() string {
	m := s.RamStamps[len(s.RamStamps)-1]
	if m < multiDigit {
		return fmt.Sprintf(" %.2f", m)
	}

	return fmt.Sprintf("%.2f", m)
}
