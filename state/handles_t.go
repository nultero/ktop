package state

import "os"

const (
	cpu = "/proc/stat"
	mem = "/proc/meminfo"
)

type handles_t struct {
	Cpu *os.File
	Mem *os.File
}

// Opens a lot of the necessary /proc/ files.
func (h *handles_t) init() error {
	f, err := os.Open(cpu)
	if err != nil {
		return err
	}
	h.Cpu = f

	f, err = os.Open(mem)
	if err != nil {
		return err
	}
	h.Mem = f

	return nil
}

// Closes **ALL** the open file handles.
func (h *handles_t) Close() {
	h.Cpu.Close()
	h.Mem.Close()
}

// Aggregates all of the file's readers to their file starts.
// Convenience method for the Collect() call in the proc dir.
func (h *handles_t) Reset() {
	h.Cpu.Seek(0, 0)
	h.Mem.Seek(0, 0)
}
