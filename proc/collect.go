package proc

import "ktop/state"

// Gathers all of the data necessary for a new
// table of processes and associated drawing.
func Collect(s *state.State) error {

	err := getCPUstats(s)
	if err != nil {
		return err
	}

	err = getMemStats(s)
	if err != nil {
		return err
	}

	s.Handles.Reset()
	return nil
}
