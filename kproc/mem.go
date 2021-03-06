package kproc

import (
	"fmt"
	"ktop/state"
	"os"
	"strconv"
)

const spc = ' '

func pollMem(stt *state.State) error {
	bytes, err := memBytes()
	if err != nil {
		return err
	}

	mem, err := getMem(bytes)
	if err != nil {
		return err
	}

	stt.Mem.Stamps = append(stt.Mem.Stamps, mem)

	if len(stt.Mem.Stamps) > stt.MaxStamps {
		stt.Mem.Stamps = stt.Mem.Stamps[1:]
	}

	return nil
}

func getMem(bytes []byte) (float64, error) {
	kilobytes := []int{}
	membytes := []byte{}

	for _, b := range bytes {
		if b != spc {
			membytes = append(membytes, b)
		}

		if b == spc && len(membytes) != 0 {
			s := string(membytes)
			kb, err := strconv.Atoi(s)
			if err == nil {
				kilobytes = append(kilobytes, kb)
			}
			membytes = []byte{}
		}
	}

	if len(kilobytes) != 2 {
		return 0.0, fmt.Errorf(
			"error in grabbing memory kilobytes from /proc/meminfo; kb slice: %v",
			kilobytes,
		)
	}

	/*
		Kb 0 should be MemTotal
		Kb 1 should be MemAvail
			100 * (1 - (k1 / k0)) = percent RAM used
			( over range [0,1] )
		See test case for example
	*/

	fl := 100.0 * (1.0 - (float64(kilobytes[1]) / float64(kilobytes[0])))
	return fl, nil
}

func memBytes() ([]byte, error) {
	bytes := []byte{}
	f, err := os.Open("/proc/meminfo")
	defer f.Close()
	if err != nil {
		return bytes, err
	}

	membytes := make([]byte, 200)
	newlines := 0

	for newlines < 2 {
		n, err := f.Read(membytes)
		if err != nil {
			return bytes, fmt.Errorf(
				"err reading /proc/meminfo bytes: %w", err,
			)
		} else if n == 0 {
			return bytes, fmt.Errorf(
				"reading line(s) from /proc/meminfo prematurely hit EOF",
			)
		}

		lnl := 0 // last newline idx
		for i, b := range membytes {
			if b == '\n' {
				newlines++
				if newlines == 1 || newlines == 3 {
					bytes = append(bytes, membytes[lnl:i]...)
				}
				lnl = i
			}
		}
	}

	return bytes, nil
}
