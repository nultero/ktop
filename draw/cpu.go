package draw

import (
	"fmt"
	"ktop/state"

	"github.com/gdamore/tcell/v2"
)

func Cpu(scr tcell.Screen, stt state.State) {
	pc := stt.Cpu.LastCPUPercent
	str := fmtPc(pc)

	x := len(str) + 1

	for i := 0; i < x+1; i++ {
		scr.SetContent(
			i, 0, space, emptyChars,
			stt.Theme.MainStyle,
		)
	}

	for i, r := range str {
		scr.SetContent(
			i+1, 0, r, emptyChars,
			stt.Theme.MainStyle,
		)
	}
}

func fmtPc(f float64) string {
	if f < multidigit {
		return fmt.Sprintf(" %.0f%%", f)
	}

	return fmt.Sprintf("%.0f%%", f)
}
