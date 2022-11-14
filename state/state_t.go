package state

import (
	"ktop/styles"
)

type State struct {
	Bounds [2]int

	Handles handles_t // Slice of *os.File; Keeps the various /proc/ files' handles open here.

	Components [4]Comp // The ordered slice of components on the screen
	Quads      [4]Quad // Corresponds to the indices of the Components struct
	Cpu        cpu_t   // Keeps two unaltered CPU stamp slices, and the last CPU sums and idles.
	Mem        stamps_t

	StampLimit int

	Time time_t

	Theme styles.Theme
}
