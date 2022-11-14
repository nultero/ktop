package state

type Quad uint8

const (
	QTopLeft  Quad = 0
	QTopRight Quad = 1
	QBtmLeft  Quad = 2
	QBtmRight Quad = 3
)

func (q Quad) GetBounds(w, h int) [4]int {
	lw, rw := 0, 0
	th, bh := 0, 0

	// if q is on the left
	if q%2 == 0 {
		rw = w / 2
	} else {
		lw = w/2 + 1
		rw = w
	}

	// if q is on top: same logic for height
	if q < 2 {
		bh = h / 2
	} else {
		th = h/2 + 1
		bh = h
	}

	return [4]int{lw, rw, th, bh}
}
