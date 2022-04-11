package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

const space = ' '
const cpuTxt = "CPU: "

var empt = []rune{}

const multiDigit float32 = 10.0

func drawChars(s tcell.Screen, cpuPc, mem float32, sty tcell.Style) {
	w, h := s.Size()
	w /= 2
	h /= 2

	cpuStr := fmt.Sprintf("%.2f", cpuPc)

	for _, r := range cpuTxt {
		s.SetContent(w, h, r, empt, sty)
		w++
	}

	if len(cpuStr) < 5 {
		s.SetContent(w, h, space, empt, tcell.StyleDefault)
		w++
	}

	for i := 0; i < len(cpuStr); i++ {
		s.SetContent(w, h, rune(cpuStr[i]), empt, sty)
		w++
	}

	s.Show()
}
