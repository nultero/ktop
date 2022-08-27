package state

import (
	"ktop/styles"
)

type State struct {
	Handles handles_t // Slice of *os.File; Keeps the various /proc/ files' handles open here.

	Cpu  cpu_t // Keeps two unaltered CPU stamp slices, and the last CPU sums and idles.
	Time time_t

	Theme styles.Theme
}
