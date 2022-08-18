package kproc

import (
	"fmt"
	"ktop/state"
	"os"
	"strconv"
	"strings"
)

func pollCPU(stt *state.State) error {
	bytes, err := cpuBytes(stt.Handles.KProcStat)
	if err != nil {
		return err
	}

	nums, err := readBytes(bytes)
	if err != nil {
		return err
	}

	if len(stt.Cpu.Slices) == 2 {
		prev := stt.Cpu.Slices[0]
		diffSum := 0
		for i, n := range nums {
			v := n - prev[i]
			prev[i] = nums[i]
			nums[i] = v
			diffSum += v
		}
		stt.Cpu.DiffSum = diffSum

	} else {
		stt.Cpu.Slices = append(stt.Cpu.Slices, nums)
	}

	curSum := 0
	for _, n := range nums {
		curSum += n
	}

	delta := curSum - stt.Cpu.Sum
	idle := nums[3] - stt.Cpu.LCI

	stt.Cpu.LCI = nums[3]
	stt.Cpu.SumPrev = stt.Cpu.Sum
	stt.Cpu.Sum = curSum
	stt.Cpu.SumNoIdle = float64(curSum - nums[3])

	pcUsed := delta - idle

	percentage := 100 * (float64(pcUsed) / float64(delta))
	stt.Cpu.Stamps = append(stt.Cpu.Stamps, percentage)

	if len(stt.Cpu.Stamps) > stt.MaxStamps {
		stt.Cpu.Stamps = stt.Cpu.Stamps[1:]
	}

	return nil
}

func readBytes(cpuBytes []byte) ([]int, error) {
	numStrs := strings.Split(
		string(cpuBytes),
		" ",
	)

	nums := []int{}
	for _, s := range numStrs[1:] {
		if len(s) == 0 {
			continue
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			return nums, err
		}

		nums = append(nums, n)
	}

	return nums, nil
}

// Equivalent to `head -n 1 /proc/stat`. First line of summary CPU stats.
// The stt.Handles.KProcStat fd being passed in is /proc/stat.
func cpuBytes(f *os.File) ([]byte, error) {
	bytes := []byte{}
	cpuBytes := make([]byte, 64)

statloop: // probably a simpler way to do this
	for {
		n, err := f.Read(cpuBytes)
		if err != nil {
			return bytes, fmt.Errorf(
				"err reading /proc/stat bytes: %w", err,
			)
		} else if n == 0 {
			return bytes, fmt.Errorf(
				"reading first line from /proc/stat prematurely hit EOF",
			)
		}

		for i, b := range cpuBytes {
			if b == '\n' {
				bytes = append(bytes, cpuBytes[:i]...)
				break statloop
			}
		}

		bytes = append(bytes, cpuBytes...)
	}

	return bytes, nil
}
