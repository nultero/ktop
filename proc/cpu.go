package proc

import (
	"ktop/state"
	"strconv"
)

// Roughly equivalent to
//
//	head -n 1 /proc/stat
func getCPUstats(s *state.State) error {

	fd := s.Handles.Cpu

	buf := make([]byte, 150)
	_, err := fd.Read(buf)
	if err != nil {
		return err
	}

	numBytes := [][]byte{}
	tmp := []byte{}
	for _, b := range buf {

		if b == '\n' {
			break
		}

		if b == 32 { // space
			if len(tmp) > 0 {
				numBytes = append(numBytes, tmp)
				tmp = []byte{}
			}
		} else if b < 60 { // must be numeric
			tmp = append(tmp, b)
		}
	}

	nums := make([]int, len(numBytes))
	for i, bytes := range numBytes {
		s := string(bytes)
		n, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		nums[i] = n
	}

	s.Cpu.Add(nums)

	return nil
}
