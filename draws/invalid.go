package draws

import (
	"fmt"
	"ktop/styles"

	"github.com/gdamore/tcell/v2"
)

func invMsg(minX, minY int) []string {
	return []string{
		" screen size invalid;",
		fmt.Sprintf(" needs at least %v  spaces width", minX),
		fmt.Sprintf(" and %v height", minY),
	}
}

// For when the given terminal bounds do not satisfy minX & minY.
func Invalid(scr tcell.Screen, sty tcell.Style, minX, minY int) {
	w, h := scr.Size()
	red := styles.InvalidRed()

	invals := invMsg(minX, minY)

	if h < 3 { // can't even display the invalid stuff;
		str := ""
		for _, s := range invals {
			str += s
		}

		for i, r := range str {
			scr.SetContent(i, 0, r, empt, red)
		}

		return
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			scr.SetContent(i, j, space, empt, sty)
		}
	}

	for i, s := range invals {
		for idx, r := range s {
			scr.SetContent(idx, i, r, empt, red)
		}
	}

	scr.Show()
}
