package kproc

import (
	"fmt"
	"ktop/state"
)

// The heaviest routine in ktop.
// Calls pollCpu, pollMem, AND reads the procfs
// for process table info.
func Top(stt *state.State) error {

	err := pollCPU(stt)
	if err != nil {
		return fmt.Errorf(
			"err polling CPU stats: %w", err,
		)
	}

	err = pollMem(stt)
	if err != nil {
		return fmt.Errorf(
			"err polling memory stats: %w", err,
		)
	}

	err = readProcfs(stt)
	if err != nil {
		return fmt.Errorf(
			"error reading process filesystem: %w", err,
		)
	}

	stt.RefreshProcTab()
	stt.Handles.Reset()

	return nil
}
