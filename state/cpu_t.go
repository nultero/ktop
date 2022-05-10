package state

type cpu_t struct {
	LCI    int // last CPU idle %, used for PollCPU
	Stamps []float64
	Sum    int
}

func defaultCpu_t() cpu_t {
	return cpu_t{
		LCI:    0,
		Stamps: []float64{},
		Sum:    0,
	}
}
