package state

type mem_t struct {
	Stamps []float64
}

func defaultMem_t() mem_t {
	return mem_t{
		Stamps: []float64{},
	}
}
