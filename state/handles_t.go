package state

import (
	"fmt"
	"os"
)

type FileHandles struct {
	KProcStat *os.File
	KMeminfo  *os.File
}

func (h *FileHandles) Reset() {
	h.KProcStat.Seek(0, 0)
	h.KMeminfo.Seek(0, 0)
}

func InitHandles() (FileHandles, error) {
	h := FileHandles{}

	f, err := os.Open("/proc/stat")
	if err != nil {
		return h, fmt.Errorf("err from reading /proc/stat: %w", err)
	}

	h.KProcStat = f

	f, err = os.Open("/proc/meminfo")
	if err != nil {
		return h, fmt.Errorf("err from reading /proc/meminfo: %w", err)
	}

	h.KMeminfo = f

	return h, nil
}

func (h FileHandles) CloseAll() {
	h.KProcStat.Close()
	h.KMeminfo.Close()
}
