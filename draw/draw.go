package draw

import (
	"ktop/state"

	"github.com/gdamore/tcell/v2"
)

func Draw(screen tcell.Screen, stt *state.State) {
	w, h := screen.Size()
	for _, q := range stt.Quads {
		bounds := q.GetBounds(w, h)
		component := stt.GetComponentInQuad(q)
		switch component {
		case state.IoCpu:
			ioCpu(screen, stt, bounds)
		default:
		}
	}
}
