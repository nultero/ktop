package okcpu

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Poll(lastIdle, sum *int) (float32, error) {
	bytes, err := cpuBytes()
	if err != nil {
		return 0.0, err
	}

	nums, err := readBytes(bytes)
	if err != nil {
		return 0.0, err
	}

	curSum := 0
	for _, n := range nums {
		curSum += n
	}

	delta := curSum - *sum
	idle := nums[3] - *lastIdle

	*lastIdle = nums[3]
	*sum = curSum

	pcUsed := delta - idle

	percentage := 100 * (float32(pcUsed) / float32(delta))

	// percentage = trunc(percentage, *prec)
	return percentage, nil
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
func cpuBytes() ([]byte, error) {
	bytes := []byte{}
	cpuBytes := make([]byte, 64)

	f, err := os.Open("/proc/stat")
	defer f.Close()
	if err != nil {
		return bytes, fmt.Errorf("err from reading /proc/stat: %w", err)
	}

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

func trunc(num, closestUnit float64) float64 {
	return math.Round(num/closestUnit) * closestUnit
}
