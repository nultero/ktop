package calcs

import "ktop/state"

func Aggregate(stt *state.State) {
	cpuPercent(stt)
	sortProcs(stt)
}
