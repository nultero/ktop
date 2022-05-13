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

	// jiffydiff
	jifdif := float64(stt.Cpu.Sum - stt.Cpu.SumPrev)

	// total jiffies
	ttl := 0.0

	// TODOOOOO this computation isn't quite there

	for _, val := range stt.PidMap {

		// TODOO needs * Cpu cores

		// ps := 100.0 * float64(val.Cur()-val.Prev()) / cpu
		ps := float64(val.Cur()-val.Prev()) / jifdif
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
