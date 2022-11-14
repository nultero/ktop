package state

const stamps_t_null = -0.1

// Kind of a ring buffer so should be fairly static memory-wise.
type stamps_t struct {
	Data []float64
	Idx  int
}

func (s *stamps_t) Add(stamp float64) {
	s.Data[s.Idx] = stamp
	s.Idx++
	if s.Idx == len(s.Data) {
		s.Idx = 0
	}
}

func (s *stamps_t) GetLastN(n int) []float64 {
	vals := make([]float64, n+1)
	for idx := s.Idx; n > 0; n-- {
		if s.Data[s.Idx] == stamps_t_null {
			break
		}
		vals[n] = s.Data[idx]
		idx--
		if idx == -1 {
			idx = len(s.Data) - 1
		}
	}

	return vals
}

// Dumps out negative values for all empties.
// Kind of niche, only useful for first X number
// of draws and stamps.
func newStamps_t(capacity int) stamps_t {
	s := stamps_t{
		Data: make([]float64, capacity),
		Idx:  0,
	}

	for i := range s.Data {
		s.Data[i] = stamps_t_null
	}

	return s
}
