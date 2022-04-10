package main

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

var empt = []rune{}

var glyphs = []rune{'@', '#', '&', '*', '=', '%', 'Z', 'A'}

var glen = len(glyphs)

var (
	lastX = 0
	lastY = 0
)

func drawChars(s tcell.Screen) {

	w, h := s.Size()

	i := rand.Intn(glen)
	c := glyphs[i]

	s.SetContent(w/2, h/2, c, empt, tcell.StyleDefault)

	s.Show()
}
