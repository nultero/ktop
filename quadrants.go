package main

type quadrant int8

type coord struct {
	x, y int
}

const (
	quadTopRight    quadrant = 0
	quadTopLeft     quadrant = 1
	quadBottomLeft  quadrant = 2
	quadBottomRight quadrant = 3
)

// Returns topleft and bottomright coords of a given quadrant.
// Upper/rightmost areas are arbitrarily greedier: an odd-number sized
// terminal w/h will give the extra point of size to the upper or
// right quadrants.
func getQuadrantXY(w, h int, q quadrant) (coord, coord) {

	tl, br := coord{}, coord{}

	switch q {
	case quadTopLeft:
		tl = coord{0, 0}
		br = coord{w / 2, h / 2}
		if isOdd(h) {
			br.y--
		}

	case quadBottomLeft:
		tl = coord{0, h / 2}
		br = coord{w / 2, h}

	case quadTopRight:
		tl = coord{w / 2, 0}
		if isOdd(w) {
			tl.x++
		}

		br = coord{w, h / 2}
		if isOdd(h) {
			br.y--
		}

	case quadBottomRight:
		tl = coord{w / 2, h / 2}
		br = coord{w, h}
	}

	return tl, br
}
