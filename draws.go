package main

import (
	"fmt"
	"ktop/styles"
	"math"

	"github.com/gdamore/tcell/v2"
)

const (
	space  = ' '
	arrow  = '➤'
	cpuTxt = "CPU: "
	memTxt = "RAM: "
)

var texts = []string{cpuTxt, memTxt}

var empt = []rune{}

const multiDigit float32 = 10.0

// Standard draw of whichever screen the focus happens to be on.
func stdDraw(scr tcell.Screen, cpuPc, mem float32, sty tcell.Style, stamps []float32) {

	// TODOOOO need a minimum area to draw graphs for, need checks upfront

	width, h := scr.Size()
	width = (width / 2) + 50
	h /= 2

	start := width + 5

	for i := start; i > 0; i-- {
		scr.SetContent(i, h-1, '─', empt, sty)
	}

	// TODOOO clean up this blunt impl of left-shifting graph grid area
	// for x := start - 1; x > 2; x-- {
	// for x := 2; x < start-1; x++ {
	for x := 2; x < start; x++ {
		for y := h - 2; y > 1; y-- {
			prv, _, _, _ := scr.GetContent(x+1, y)
			scr.SetContent(x, y, prv, empt, sty)
		}
	}

	subh := h - 2
	dotmaxh := (h - 2) * 4 // dots possible within the height monospaces available
	dots := int(math.Round(float64(dotmaxh) * (float64(cpuPc) / 100.0)))
	for subh > 1 {
		if dots == 0 {
			scr.SetContent(start, subh, space, empt, sty)
		} else if dots >= 4 {
			scr.SetContent(start, subh, dotrunes[4], empt, sty)
			dots -= 4
		} else {
			scr.SetContent(start, subh, dotrunes[dots], empt, sty)
			dots -= dots
		}

		subh--
	}

	cpuStr := fmt.Sprintf("%.2f", cpuPc)
	memStr := fmt.Sprintf("%.2f", mem)

	strs := []string{cpuStr, memStr}

	for idx, s := range strs {
		w := width
		txt := texts[idx]

		if idx == 0 {
			scr.SetContent(w-2, h, arrow, empt, sty)
		}

		for _, r := range txt {
			scr.SetContent(w, h+idx, r, empt, sty)
			w++
		}

		if len(s) < 5 {
			scr.SetContent(w, h+idx, space, empt, styles.AllBlack())
			w++
		}

		for i := 0; i < len(s); i++ {
			scr.SetContent(w, h+idx, rune(s[i]), empt, sty)
			w++
		}
	}

	scr.Show()
}

func redraw(scr tcell.Screen, sty tcell.Style) {

	// TODO rename this func and docstring it better so I don't come back confused

	w, h := scr.Size()

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			scr.SetContent(i, j, space, empt, sty)
		}
	}

	scr.Show()
}

/*

	Invalid screen size stuff below

*/

var invals = []string{
	" screen size invalid;",
	" needs at least 80 spaces width",
	" and 24 height"}

func invalidSzDraw(scr tcell.Screen, sty tcell.Style) {
	w, h := scr.Size()
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			scr.SetContent(i, j, space, empt, sty)
		}
	}

	for i, s := range invals {
		for idx, r := range s {
			scr.SetContent(idx, i, r, empt, styles.InvalidRed())
		}
	}

	scr.Show()
}
