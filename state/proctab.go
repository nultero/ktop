package state

import "sort"

// Keyed by percent CPU usage, values are
// a slice of process names at that level of
// of strain (i.e., many sleeping procs will be 0%)
type ProcTab map[float64][]string

// Refreshes the internal state's process table
// (needs a recently completed series of reads
// from the procfs -- all 3 from the precursor
// Top calls).
func (stt *State) RefreshProcTab() {
	stt.Top.flush()
	ttl := 0.0

	for _, val := range stt.PidMap {
		ps := val.cpuPc
		ttl += ps

		if procs, ok := stt.Top[ps]; ok {
			procs = append(procs, val.Name())

		} else {
			stt.Top[ps] = []string{val.Name()}
		}
	}

	stt.Total = ttl
}

// Removes old process data that was ephemeral anyway.
func (pt ProcTab) flush() {
	for k := range pt {
		delete(pt, k)
	}
}

// Returns a sorted slice of the CPU usages.
func (pt ProcTab) Percents() []float64 {
	pcs := []float64{}
	for pc := range pt {
		pcs = append(pcs, pc)
	}

	sort.Float64s(pcs)
	return pcs
}
