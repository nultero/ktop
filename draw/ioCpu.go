package draw

import (
	"fmt"
	"ktop/state"

	"github.com/gdamore/tcell/v2"
)

func ioCpu(scr tcell.Screen, stt *state.State, bounds [4]int) {

	lw, rw, _, bh := bounds[0], bounds[1], bounds[2], bounds[3]

	// draw the border line
	lineHeight := bh - 3
	for x := lw; x < rw; x++ {
		scr.SetContent(
			x, lineHeight, lineChar, emptyChars,
			stt.Theme.MainStyle,
		)
	}

	// draw CPU: line
	lineHeight++
	pc := formatPercent(stt.Cpu.LastCPUPercent)
	str := fmt.Sprintf("%s%s", cpu, pc)
	offset := len(str) - 1
	for x := rw - 1; offset >= 0; x-- {
		scr.SetContent(
			x, lineHeight, rune(str[offset]), emptyChars,
			stt.Theme.MainStyle,
		)
		offset--
	}
}

func formatPercent(f float64) string {
	if f < multidigit {
		return fmt.Sprintf(" %.0f%%", f)
	}

	return fmt.Sprintf("%.0f%%", f)
}
