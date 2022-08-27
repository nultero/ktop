package state

import "os"

const (
	cpu = "/proc/stat"
)

type handles_t struct {
	Cpu *os.File
}

// Opens a lot of the necessary /proc/ files.
func (h *handles_t) init() error {
	f, err := os.Open(cpu)
	if err != nil {
		return err
	}
	h.Cpu = f

	return nil
}

// Closes **ALL** the open file handles.
func (h *handles_t) Close() {
	h.Cpu.Close()
}

// Aggregates all of the file's readers to their file starts.
// Convenience method for the Collect() call in the proc dir.
func (h *handles_t) Reset() {
	h.Cpu.Seek(0, 0)
}
