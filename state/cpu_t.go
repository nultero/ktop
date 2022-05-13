package state

import (
	"fmt"
	"os/exec"
	"strconv"
)

type cpu_t struct {
	LCI     int     // last CPU idle %, used for PollCPU
	HZ      float64 // clock tic hertz, originally a double in procps, used here for conversions to jiffies
	Stamps  []float64
	Sum     int // CPU jiffies currently in play
	SumPrev int // CPU jiffies previous
}

func defaultCpu_t() cpu_t {
	return cpu_t{
		LCI:     0,
		HZ:      getHertz(),
		Stamps:  []float64{},
		Sum:     0,
		SumPrev: 0,
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

func getHertz() float64 {

	/*
	 this honestly seemed like
	 the simplest way to get this info;
	 there is exposed HZ figures in
	 /proc/cpuinfo but they can vary,
	 apparently, and so here is the
	 dumber variant
	*/

	c := exec.Command("getconf", "CLK_TCK")
	b, err := c.Output()
	if err != nil {
		e := fmt.Errorf("err interpreting `getconf` output for clock ticks: %w", err)
		panic(e)
	}

	bytes := []byte{}
	for _, ch := range b {
		if ch != '\n' {
			bytes = append(bytes, ch)
		}
	}

	s := string(bytes)
	hz, err := strconv.ParseFloat(s, 64)
	if err != nil {
		e := fmt.Errorf("err interpreting `getconf` string for clock ticks: %w", err)
		panic(e)
	}

	return hz
}
