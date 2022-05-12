package draws

import (
	"ktop/state"

	"github.com/gdamore/tcell/v2"
)

// Flushes the main state's style so that on resize, things
// don't bork entirely.
func Refresh(scr tcell.Screen, stt *state.State) {
	w, h := scr.Size()
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			scr.SetContent(i, j, space, empt, stt.ColorTheme.MainStyle)
		}
	}
}
