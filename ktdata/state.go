package ktdata

import (
	"ktop/styles"
	"time"
)

type State struct {
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
