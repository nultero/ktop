package state

import "fmt"

type mem_t struct {
	Stamps []float64
}

func defaultMem_t() mem_t {
	return mem_t{
		Stamps: []float64{},
	}
}

// Last memory stamp percent.
func (m mem_t) Last() float64 {
	return m.Stamps[len(m.Stamps)-1]
}

// Last memory stamp percent Sprintf'd to a string.
func (m mem_t) LastToStr() string {
	ms := m.Stamps[len(m.Stamps)-1]
	if ms < multiDigit { // digit is in ./cpu_t.go
		return fmt.Sprintf(" %.2f", ms)
	}

	return fmt.Sprintf("%.2f", ms)
}
