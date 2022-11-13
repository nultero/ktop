package draw

import "github.com/gdamore/tcell/v2"

const msg = "screen size too small"

func Invalid(scr tcell.Screen, sty tcell.Style) {
	scr.Clear()
	x, _ := scr.Size()
	y := 0
	for i := 0; i < len(msg); i++ {
		if i > x {
			y++
		}
		scr.SetContent(
			i, y, rune(msg[i]), emptyChars, sty,
		)
	}

}
