package calcs

import (
	"ktop/state"
	"sort"
)

func sortProcs(s *state.State) {
	keys := []float64{}
	m := map[float64]string{}
	for _, p := range s.ProcessMap {
		pc := p.CpuPc()
		if pc > 0.01 {
			keys = append(keys, pc)
			for {
				_, ok := m[pc]
				if ok { // collisions
					pc += 0.001
				} else {
					m[pc] = p.Name()
					break
				}
			}
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	s.SortedProcesses.Keys = keys
	s.SortedProcesses.Map = m
}
