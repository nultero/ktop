package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

const (
	space  = ' '
	cpuTxt = "CPU: "
	memTxt = "RAM: "
)

var texts = []string{cpuTxt, memTxt}

var empt = []rune{}

const multiDigit float32 = 10.0

func drawChars(scr tcell.Screen, cpuPc, mem float32, sty tcell.Style) {
	width, h := scr.Size()
	width /= 2
	h /= 2

	cpuStr := fmt.Sprintf("%.2f", cpuPc)
	memStr := fmt.Sprintf("%.2f", mem)

	strs := []string{cpuStr, memStr}

	for idx, s := range strs {
		w := width
		txt := texts[idx]
		for _, r := range txt {
			scr.SetContent(w, h+idx, r, empt, sty)
			w++
		}

		if len(s) < 5 {
			scr.SetContent(w, h+idx, space, empt, tcell.StyleDefault)
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
	w, h := scr.Size()

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			scr.SetContent(i, j, space, empt, sty)
		}
	}

	scr.Show()
}
