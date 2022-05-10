package state

import "github.com/gdamore/tcell/v2"

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
	return stt.quad == q
}

// Primary state method of changing views / quads.
func (stt *State) MoveQuad(k tcell.Key) {

	switch k {
	case tcell.KeyDown:
		if stt.quad == QuadTopRight {
			stt.quad = QuadBottomRight
		} else if stt.quad == QuadTopLeft {
			stt.quad = QuadBottomLeft
		}

	case tcell.KeyUp:
		if stt.quad == QuadBottomRight {
			stt.quad = QuadTopRight
		} else if stt.quad == QuadBottomLeft {
			stt.quad = QuadTopLeft
		}

	case tcell.KeyLeft:
		if stt.quad == QuadBottomRight {
			stt.quad = QuadBottomLeft
		} else if stt.quad == QuadTopRight {
			stt.quad = QuadTopLeft
		}

	case tcell.KeyRight:
		if stt.quad == QuadBottomLeft {
			stt.quad = QuadBottomRight
		} else if stt.quad == QuadTopLeft {
			stt.quad = QuadTopRight
		}
	}
}

// Returns topleft and bottomright Coords of a given quadrant.
// Upper/rightmost areas are arbitrarily greedier: an odd-number sized
// terminal w/h will give the extra point of size to the upper or
// right quadrants.
func GetQuadrantXY(w, h int, q Quadrant) (Coord, Coord) {

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
