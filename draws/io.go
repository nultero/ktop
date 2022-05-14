package draws

import (
	"fmt"
	"ktop/state"
	"math"

	"github.com/gdamore/tcell/v2"
)

/*

	TODOO: viewstate as a function of coordinates;
	write tests that confirm that the renders do what
	they are supposed to do

*/

// This view shows CPU | Mem usage in a pretty graph.
func Io(scr tcell.Screen, stt *state.State, q state.Quadrant) {
	w, h := scr.Size()
	tl, br := q.GetBounds(w, h)
	ox := br.X - 2 // offset x
	onQ := stt.OnQuad(q)

	memstr := cprint(
		memTxt, stt.Mem.LastToStr(), onQ, stt.IsFocused(state.MemGraph),
	)

	cpustr := cprint(
		cpuTxt, stt.Cpu.LastToStr(), onQ, stt.IsFocused(state.CpuGraph),
	)

	// TODO if you assign these a map, you can put them
	// into a much simpler, less repetitive loop

	for idx, r := range memstr {
		x := ox + idx - len(memstr)
		if stt.IsFocused(state.MemGraph) {
			if onQ && idx == 0 {
				scr.SetContent(x, br.Y, r, empt, stt.ColorTheme.HighlightStyle)

			} else {
				scr.SetContent(x, br.Y, r, empt, stt.ColorTheme.MainStyle)
			}
		} else {
			scr.SetContent(x, br.Y, r, empt, stt.ColorTheme.InactiveStyle)
		}
	}

	for idx, r := range cpustr {
		x := ox + idx - len(cpustr)
		if stt.IsFocused(state.CpuGraph) {
			if onQ {
				// if onQ && idx == 0 {
				scr.SetContent(x, br.Y-1, r, empt, stt.ColorTheme.HighlightStyle)
			} else {
				scr.SetContent(x, br.Y-1, r, empt, stt.ColorTheme.MainStyle)
			}
		} else {
			scr.SetContent(x, br.Y-1, r, empt, stt.ColorTheme.InactiveStyle)

		}
	}

	// bottom of graph line
	for i := ox; i > tl.X; i-- {
		scr.SetContent(i, br.Y-2, '─', empt, stt.ColorTheme.InactiveStyle)
	}

	/*
		TODOOO -- need focused conditional to draw CPU vs. Mem
		Graph draws:
	*/

	xdiff := (ox - 1) - (tl.X + 1)
	subh := br.Y - tl.Y - 2
	charh := subh * 4 // chars' possible heights within the quadrant

	if len(stt.Cpu.Stamps) == 0 {
		return
	}

	sty := stt.ColorTheme.InactiveStyle
	if onQ {
		sty = stt.ColorTheme.MainStyle
	}

	for i := len(stt.Cpu.Stamps) - 1; i > 0; i-- {

		if xdiff == 0 {
			break
		}

		x := ox - (len(stt.Cpu.Stamps) - i)
		y := br.Y - 3

		dots := int(math.Round(
			float64(charh) * stt.Cpu.Stamps[i] / 100.0,
		))

		for y > tl.Y+1 {
			if dots == 0 {
				scr.SetContent(x, y, space, empt, sty)

			} else if dots >= 4 {
				scr.SetContent(x, y, dotrunes[4], empt, sty)
				dots -= 4

			} else {
				scr.SetContent(x, y, dotrunes[dots], empt, sty)
				dots = 0
			}

			y--
		}

		xdiff--
	}
}

// Conditional arrow cursor added to string if focused.
// TODO: customizable cursor char?
func cprint(txt, percent string, onQ, isFocused bool) string {
	if onQ && isFocused {
		return fmt.Sprintf("%c%v%v", arrow, txt, percent)
	}

	return fmt.Sprintf("   %v%v", txt, percent)
}
