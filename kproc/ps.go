package kproc

import (
	"errors"
	"fmt"
	"ktop/state"
	"os"
	"strconv"
)

func readProcfs(stt *state.State) error {
	dir, err := os.ReadDir("/proc")
	if err != nil {
		return fmt.Errorf("procfs unavailable: %w", err)
	}

	pidSet := map[uint64]struct{}{}

	for _, f := range dir {
		if '0' <= f.Name()[0] && f.Name()[0] <= '9' {

			// these errors aren't all necessarily fatal
			// just appending them to the state.Log is fine

			pid, err := strconv.ParseUint(f.Name(), 0, 32)
			if err != nil {
				e := fmt.Errorf(
					"pid '%v' failed to parse: %w", f.Name(), err,
				).Error()

				stt.Log = append(stt.Log, e)
				if len(stt.Log) > stt.MaxStamps {
					stt.Log = stt.Log[1:]
				}
			}

			bytes, err := readPidBytes(f.Name())
			if err != nil {
				e := fmt.Errorf(
					"error reading pid '%v' bytes: %w", f.Name(), err,
				).Error()

				stt.Log = append(stt.Log, e)
				if len(stt.Log) > stt.MaxStamps {
					stt.Log = stt.Log[1:]
				}
			}

			name, utime, stime := parseStat(bytes)

			pidSet[pid] = struct{}{}

			if _, ok := stt.PidMap[pid]; ok {
				err := stt.PidMap.UpdateProc(pid, utime, stime, stt.Cpu.SumNoIdle, stt.Cpu.Last())
				if err != nil {
					e := fmt.Errorf(
						"error updating pid '%v' bytes: %w", f.Name(), err,
					).Error()

					stt.Log = append(stt.Log, e)
					if len(stt.Log) > stt.MaxStamps {
						stt.Log = stt.Log[1:]
					}
				}
			} else if !ok {
				stt.PidMap.NewProc(name, pid, utime, stime, stt.Cpu.SumNoIdle, stt.Cpu.Last())
			}
		}
	}

	/*
		if there's a pid in the state.PidMap
		that isn't in the procfs,
		it's old and needs to go
	*/
	for pid := range stt.PidMap {
		if _, ok := pidSet[pid]; !ok {
			delete(stt.PidMap, pid)
		}
	}

	return nil
}

func readPidBytes(pid string) ([]byte, error) {
	pbytes := make([]byte, 150)
	f, err := os.Open(fmt.Sprintf("/proc/%v/stat", pid))
	if err != nil {
		return pbytes, err
	}
	defer f.Close()

	n, err := f.Read(pbytes)
	if err != nil {
		return pbytes, fmt.Errorf(
			"err reading /proc/%v/stat bytes: %w", pid, err,
		)
	} else if n == 0 {
		return pbytes, errors.New(fmt.Sprintf(
			"reading line(s) from /proc/%v/stat prematurely hit EOF", pid,
		))
	}

	return pbytes, nil
}

func parseStat(bytes []byte) (string, int64, int64) {

	var (
		name      = []byte{}
		nidx byte = 0 // name's start index

		utime, stime int64

		idxs = []byte{} // indexes of spaces seen

		// in process bool, as some
		// procs can have spaces in
		// their name
		inp bool = false
	)

	for i := 0; i < len(bytes); i++ {
		b := bytes[i]

		if b == ' ' && !inp {
			idxs = append(idxs, byte(i))

			// this is the space after the 15th
			// item, which is the stime jiffies
			if len(idxs) == 15 {

				sX := idxs[len(idxs)-1] // stime end
				uX := idxs[len(idxs)-2] // utime end, stime start
				uS := idxs[len(idxs)-3] //utime start

				uStr := string(bytes[uS+1 : uX])
				utime, _ = strconv.ParseInt(uStr, 0, 64)

				sStr := string(bytes[uX+1 : sX])
				stime, _ = strconv.ParseInt(sStr, 0, 64)
				break
			}

		} else if b == '(' {
			inp = true
			nidx = byte(i + 1)

		} else if b == ')' {
			name = bytes[nidx:i]
			inp = false
		}
	}

	return string(name), utime, stime
}
