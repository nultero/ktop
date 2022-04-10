package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

const space = ' '

var empt = []rune{}

const multiDigit float32 = 10.0

func drawChars(s tcell.Screen, cpuPc float32) {
	w, h := s.Size()
	w /= 2
	h /= 2

	cpuStr := fmt.Sprintf("%.2f", cpuPc)
	if cpuPc < multiDigit {
		s.SetContent(w, h, space, empt, tcell.StyleDefault)
		w--
	}

	for i := len(cpuStr) - 1; i >= 0; i-- {
		s.SetContent(w, h, rune(cpuStr[i]), empt, tcell.StyleDefault)
		w--
	}

	s.Show()
}
