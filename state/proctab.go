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

	cpu := float64(stt.Cpu.Sum - stt.Cpu.SumPrev)
	if cpu == 0 {
		// try to prevent divide-by-zero
		// probably rare, but eh
		return
	}

	// TODOOOOO this computation isn't quite there

	for _, val := range stt.PidMap {

		// TODOO needs * Cpu cores

		ps := 100 * float64(val.Cur()-val.Prev()) / cpu

		if procs, ok := stt.Top[ps]; ok {
			procs = append(procs, val.Name())

		} else {
			stt.Top[ps] = []string{val.Name()}
		}
	}
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

// for f := range pmap {
// 	pcs = append(pcs, f)
// }
// sort.Float64s(pcs)

// for _, f := range pcs {
// 	nStrs := pmap[f]
// 	for _, s := range nStrs {
// 		procNames = append(procNames, s)
// 	}
// }
