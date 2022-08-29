package draw

import (
	"fmt"
	"ktop/state"

	"github.com/gdamore/tcell/v2"
)

func Cpu(scr tcell.Screen, stt state.State) {
	pc := stt.Cpu.LastCPUPercent
	str := fmt.Sprintf("%.2f%%", pc)

	for i := range str {
		scr.SetContent(
			i, 0, space, emptyChars,
			stt.Theme.MainStyle,
		)
	}

	for i, r := range str {
		scr.SetContent(
			i, 0, r, emptyChars,
			stt.Theme.MainStyle,
		)
	}
}
