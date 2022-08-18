package state

import (
	"fmt"
)

type cpu_t struct {
	LCI       int // last CPU idle %, used for PollCPU
	Slices    [][]int
	DiffSum   int
	Stamps    []float64
	Sum       int     // CPU jiffies currently in play
	SumNoIdle float64 // CPU jiffies currently in play
	SumPrev   int     // CPU jiffies previous
}

func defaultCpu_t() cpu_t {
	return cpu_t{
		LCI:       0,
		DiffSum:   0,
		Slices:    [][]int{},
		Stamps:    []float64{},
		Sum:       0,
		SumNoIdle: 0.0,
		SumPrev:   0,
	}
}

const multiDigit float64 = 10.0

// Last CPU percent.
func (c cpu_t) Last() float64 {
	return c.Stamps[len(c.Stamps)-1]
}

// Last CPU percent Sprintf'd to a string.
func (c cpu_t) LastToStr() string {
	pc := c.Stamps[len(c.Stamps)-1]
	if pc < multiDigit {
		return fmt.Sprintf(" %.2f", pc)
	}

	return fmt.Sprintf("%.2f", pc)
}
