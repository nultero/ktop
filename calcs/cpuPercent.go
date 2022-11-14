package calcs

import "ktop/state"

func cpuPercent(stt *state.State) {
	sum := 0
	for _, n := range stt.Cpu.Cur {
		sum += n
	}

	delta := sum - stt.Cpu.LastSum
	stt.Cpu.DiffSum = delta

	idle := stt.Cpu.Cur[3] - stt.Cpu.LastCPUIdle
	stt.Cpu.LastCPUIdle = stt.Cpu.Cur[3]

	stt.Cpu.LastSum = sum
	stt.Cpu.LastSumNoIdle = sum - idle

	used := delta - idle

	lcpu := 100.0 * (float64(used) / float64(delta))
	stt.Cpu.LastCPUPercent = lcpu
	stt.Cpu.Stamps.Add(lcpu)
}
