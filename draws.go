package main

import (
	"fmt"
	"ktop/state"
	"ktop/styles"
	"math"

	"github.com/gdamore/tcell/v2"
)

/*

	TODOO: viewstate as a function of coordinates;
	write tests that confirm that the renders do what
	they are supposed to do

*/

const (
	space  = ' '
	arrow  = '➤'
	cpuTxt = "CPU: "
	memTxt = "RAM: "
)

var iotexts = []string{cpuTxt, memTxt}

var empt = []rune{}

func ioDraw(scr tcell.Screen, stt *state.State, q state.Quadrant) {
	w, h := scr.Size()
	tl, br := state.GetQuadrantXY(w, h, q)
	ox := br.X - 2 // offset x
	isQ := stt.IsQuad(q)

	memstr := cprint(
		memTxt, stt.LMemPCStr(), isQ, stt.IsFocused(state.MemGraph),
	)

	cpustr := cprint(
		cpuTxt, stt.LCpuPCStr(), isQ, stt.IsFocused(state.CpuGraph),
	)

	// TODO if you assign these a map, you can put them
	// into a much simpler, less repetitive loop

	for idx, r := range memstr {
		x := ox + idx - len(memstr)
		if stt.IsFocused(state.MemGraph) {
			if isQ && idx == 0 {
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
			if isQ {
				// if isQ && idx == 0 {
				scr.SetContent(x, br.Y-1, r, empt, stt.ColorTheme.HighlightStyle)
			} else {
				scr.SetContent(x, br.Y-1, r, empt, stt.ColorTheme.MainStyle)
			}
		} else {
			scr.SetContent(x, br.Y-1, r, empt, stt.ColorTheme.InactiveStyle)

		}
	}

	// drawing the flash line; this is always drawn in the InactiveStyle
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

	if len(stt.CpuStamps) == 0 {
		return
	}

	for i := len(stt.CpuStamps) - 1; i > 0; i-- {

		if xdiff == 0 {
			break
		}

		x := ox - (len(stt.CpuStamps) - i)
		y := br.Y - 3

		dots := int(math.Round(
			float64(charh) * stt.CpuStamps[i] / 100.0,
		))

		sty := stt.ColorTheme.MainStyle
		if isQ {
			sty = stt.ColorTheme.AccentStyle
		}

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
func cprint(txt, percent string, isQ, isFocused bool) string {
	if isQ && isFocused {
		return fmt.Sprintf("%c%v%v", arrow, txt, percent)
	}

	return fmt.Sprintf("  %v%v", txt, percent)
}

// Flushes the main state's style so that on resize, things
// don't bork entirely.
func redraw(scr tcell.Screen, stt *state.State) {
	w, h := scr.Size()
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			scr.SetContent(i, j, space, empt, stt.ColorTheme.MainStyle)
		}
	}
}

/*

	Invalid screen size stuff below

*/

var invals = []string{
	" screen size invalid;",
	" needs at least 30 spaces width", // TODO clean up the invalid size jank/mismatch possibilities
	" and 18 height"}

func invalidSzDraw(scr tcell.Screen, sty tcell.Style) {
	w, h := scr.Size()
	red := styles.InvalidRed()

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
