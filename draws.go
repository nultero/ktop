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

// Standard draw of whichever screen the focus happens to be on.
func stdDraw(scr tcell.Screen, stt *state.State) {

	width, h := scr.Size()
	// width = (width / 2) + 10
	h /= 2
	width -= 3

	// start := width + 10

	for i := width; i > 0; i-- {
		scr.SetContent(i, h-1, '─', empt, stt.ColorTheme.InactiveStyle)
	}

	// TODOOO clean up this blunt impl of left-shifting graph grid area
	// for x := start - 1; x > 2; x-- {
	// for x := 2; x < start-1; x++ {
	for x := 2; x < width; x++ {
		for y := h - 2; y > 1; y-- {
			prv, _, _, _ := scr.GetContent(x+1, y)
			scr.SetContent(x, y, prv, empt, stt.ColorTheme.MainStyle)
		}
	}

	cpuPc := stt.CpuStamps[len(stt.CpuStamps)-1]
	mem := stt.RamStamps[len(stt.RamStamps)-1]

	subh := h - 2
	dotmaxh := (h - 2) * 4                                               // dots possible within the height monospaces available
	dots := int(math.Round(float64(dotmaxh) * (float64(cpuPc) / 100.0))) // TODO make this int conversion a little tighter
	for subh > 1 {
		if dots == 0 {
			scr.SetContent(width, subh, space, empt, stt.ColorTheme.MainStyle)
		} else if dots >= 4 {
			scr.SetContent(width, subh, dotrunes[4], empt, stt.ColorTheme.MainStyle)
			dots -= 4
		} else {
			scr.SetContent(width, subh, dotrunes[dots], empt, stt.ColorTheme.MainStyle)
			dots -= dots
		}

		subh--
	}

	cpuStr := fmt.Sprintf("%.2f", cpuPc)
	memStr := fmt.Sprintf("%.2f", mem)

	strs := []string{cpuStr, memStr}

	for idx, s := range strs {
		w := width - 10
		txt := iotexts[idx]

		in := false

		if idx == 0 {
			scr.SetContent(w-2, h, arrow, empt, stt.ColorTheme.MainStyle)
			in = true
		}

		for _, r := range txt {
			if !in {
				scr.SetContent(w, h+idx, r, empt, stt.ColorTheme.InactiveStyle)
			} else {
				scr.SetContent(w, h+idx, r, empt, stt.ColorTheme.MainStyle)
			}
			w++
		}

		if len(s) < 5 {
			scr.SetContent(w, h+idx, space, empt, stt.ColorTheme.MainStyle)
			w++
		}

		for i := 0; i < len(s); i++ {
			if !in {
				scr.SetContent(w, h+idx, rune(s[i]), empt, stt.ColorTheme.InactiveStyle)
			} else {
				scr.SetContent(w, h+idx, rune(s[i]), empt, stt.ColorTheme.MainStyle)
			}

			w++
		}
	}

	scr.Show()
}

func ioDraw(scr tcell.Screen, stt *state.State, q state.Quadrant) {
	w, h := scr.Size()
	tl, br := state.GetQuadrantXY(w, h, q)
	ox := br.X - 2 // offset x

	// TODOOOO active/inactive styles checks/assigns here

	memstr := fmt.Sprintf("%v%v", memTxt, stt.LastMemPCString())
	for idx, r := range memstr {
		scr.SetContent(ox+idx-len(memstr), br.Y, r, empt, stt.ColorTheme.InactiveStyle)
	}

	// TODO tmp draw
	cpustr := fmt.Sprintf("%c%v%v", arrow, cpuTxt, stt.LastCpuPCString())
	// cpustr := fmt.Sprintf("%c%v%v", arrow, "emotional wellbeing", stt.LastCpuPCString())
	for idx, r := range cpustr {
		if idx == 0 {
			scr.SetContent(ox+idx-len(cpustr), br.Y-1, r, empt, stt.ColorTheme.HighlightStyle)
		} else {
			scr.SetContent(ox+idx-len(cpustr), br.Y-1, r, empt, stt.ColorTheme.MainStyle)
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

		for y > tl.Y+1 {
			if dots == 0 {
				scr.SetContent(x, y, space, empt, stt.ColorTheme.MainStyle)

			} else if dots >= 4 {
				scr.SetContent(x, y, dotrunes[4], empt, stt.ColorTheme.MainStyle)
				dots -= 4

			} else {
				scr.SetContent(x, y, dotrunes[dots], empt, stt.ColorTheme.MainStyle)
				dots = 0
			}

			y--
		}

		xdiff--
	}

	// subh := h - 2
	// dotmaxh := (h - 2) * 4                                               // dots possible within the height monospaces available
	// dots := int(math.Round(float64(dotmaxh) * (float64(cpuPc) / 100.0))) // TODO make this int conversion a little tighter
	// for subh > 1 {
	// 	if dots == 0 {
	// 		scr.SetContent(width, subh, space, empt, stt.ColorTheme.MainStyle)
	// 	} else if dots >= 4 {
	// 		scr.SetContent(width, subh, dotrunes[4], empt, stt.ColorTheme.MainStyle)
	// 		dots -= 4
	// 	} else {
	// 		scr.SetContent(width, subh, dotrunes[dots], empt, stt.ColorTheme.MainStyle)
	// 		dots -= dots
	// 	}

	// 	subh--
	// }

	// cpuStr := fmt.Sprintf("%.2f", cpuPc)
	// memStr := fmt.Sprintf("%.2f", mem)

	// strs := []string{cpuStr, memStr}

	// for idx, s := range strs {
	// 	w := width - 10
	// 	txt := texts[idx]

	// 	in := false

	// 	if idx == 0 {
	// 		scr.SetContent(w-2, h, arrow, empt, stt.ColorTheme.MainStyle)
	// 		in = true
	// 	}

	// 	for _, r := range txt {
	// 		if !in {
	// 			scr.SetContent(w, h+idx, r, empt, stt.ColorTheme.InactiveStyle)
	// 		} else {
	// 			scr.SetContent(w, h+idx, r, empt, stt.ColorTheme.MainStyle)
	// 		}
	// 		w++
	// 	}

	// 	if len(s) < 5 {
	// 		scr.SetContent(w, h+idx, space, empt, stt.ColorTheme.MainStyle)
	// 		w++
	// 	}

	// 	for i := 0; i < len(s); i++ {
	// 		if !in {
	// 			scr.SetContent(w, h+idx, rune(s[i]), empt, stt.ColorTheme.InactiveStyle)
	// 		} else {
	// 			scr.SetContent(w, h+idx, rune(s[i]), empt, stt.ColorTheme.MainStyle)
	// 		}

	// 		w++
	// 	}
	// }
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
