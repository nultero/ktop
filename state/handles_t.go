package state

import (
	"fmt"
	"os"
)

type FileHandles struct {
	KProcStat *os.File
	KMeminfo  *os.File
}

func (stt *State) InitHandles() error {
	h := FileHandles{}

	f, err := os.Open("/proc/stat")
	if err != nil {
		return fmt.Errorf("err from reading /proc/stat: %w", err)
	}

	h.KProcStat = f

	f, err = os.Open("/proc/meminfo")
	if err != nil {
		return fmt.Errorf("err from reading /proc/meminfo: %w", err)
	}

	h.KMeminfo = f

	stt.Handles = h
	return nil
}

func (h FileHandles) CloseAll() {
	h.KProcStat.Close()
	h.KMeminfo.Close()
}
