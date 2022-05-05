package ktdata

import (
	"fmt"
	"ktop/styles"
	"time"
)

type View uint8

type State struct {
	Focused   View
	PollRate  time.Duration
	LCI       int // last CPU idle %, used for PollCPU
	CpuSum    int
	CpuStamps []float32
	RamStamps []float32
	MaxStamps int // how many cpu / mem stamps to keep

	ColorTheme styles.Theme

	NeedsRedraw bool
}

func DefaultState() State {
	return State{
		Focused:   1, // corresponds to CPU in ../viewmap.go
		PollRate:  500 * time.Millisecond,
		LCI:       0,
		CpuSum:    0,
		CpuStamps: []float32{},
		RamStamps: []float32{},
		MaxStamps: 30,

		ColorTheme: styles.CrystalTheme(),

		NeedsRedraw: false,
	}
}

const multiDigit float32 = 10.0

func (s *State) LastCpuPC() float32 {
	return s.CpuStamps[len(s.CpuStamps)-1]
}

func (s *State) LastCpuPCString() string {
	c := s.CpuStamps[len(s.CpuStamps)-1]
	if c < multiDigit {
		return fmt.Sprintf(" %.2f", c)
	}

	return fmt.Sprintf("%.2f", c)
}

func (s *State) LastMemPC() float32 {
	return s.RamStamps[len(s.RamStamps)-1]
}

func (s *State) LastMemPCString() string {
	m := s.RamStamps[len(s.RamStamps)-1]
	if m < multiDigit {
		return fmt.Sprintf(" %.2f", m)
	}

	return fmt.Sprintf("%.2f", m)
}
