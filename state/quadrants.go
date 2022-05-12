package state

import (
	"github.com/gdamore/tcell/v2"
)

type Coord struct {
	X, Y int
}

type Quadrant uint8

const (
	QuadTopRight    Quadrant = 0
	QuadTopLeft     Quadrant = 1
	QuadBottomLeft  Quadrant = 2
	QuadBottomRight Quadrant = 3
)

// Returns whether state is focused on quad q.
func (stt *State) OnQuad(q Quadrant) bool {
	return stt.Quad == q
}

// Primary state method of changing views / quads.
func (stt *State) MoveQuad(k tcell.Key) {
	switch k {
	case tcell.KeyDown:
		if stt.Quad == QuadTopRight {
			stt.Quad = QuadBottomRight
		} else if stt.Quad == QuadTopLeft {
			stt.Quad = QuadBottomLeft
		}

	case tcell.KeyUp:
		if stt.Quad == QuadBottomRight {
			stt.Quad = QuadTopRight
		} else if stt.Quad == QuadBottomLeft {
			stt.Quad = QuadTopLeft
		}

	case tcell.KeyLeft:
		if stt.Quad == QuadBottomRight {
			stt.Quad = QuadBottomLeft
		} else if stt.Quad == QuadTopRight {
			stt.Quad = QuadTopLeft
		}

	case tcell.KeyRight:
		if stt.Quad == QuadBottomLeft {
			stt.Quad = QuadBottomRight
		} else if stt.Quad == QuadTopLeft {
			stt.Quad = QuadTopRight
		}
	}
}

// Returns topleft and bottomright Coords of a given Quadrant.
// Upper/rightmost areas are arbitrarily greedier: an odd-number sized
// terminal w/h will give the extra point of size to the upper or
// right Quadrants.
func (q Quadrant) GetBounds(w, h int) (Coord, Coord) {

	h -= 2 // bottom edge buffer
	tl, br := Coord{}, Coord{}

	switch q {
	case QuadTopLeft:
		tl = Coord{0, 0}
		br = Coord{w / 2, h / 2}
		if isOdd(h) {
			br.Y++
		}

	case QuadBottomLeft:
		tl = Coord{0, h / 2}
		br = Coord{w / 2, h}

	case QuadTopRight:
		tl = Coord{w / 2, 0}
		if isOdd(w) {
			tl.X++
		}

		br = Coord{w, h / 2}
		if isOdd(h) {
			br.Y++
		}

	case QuadBottomRight:
		tl = Coord{w / 2, h / 2}
		br = Coord{w, h}
		if isOdd(w) {
			br.X++
		}
	}

	return tl, br
}

func isOdd(n int) bool {
	return n%2 != 0
}
