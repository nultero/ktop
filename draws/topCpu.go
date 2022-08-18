package draws

import (
	"fmt"
	"ktop/state"

	"github.com/gdamore/tcell/v2"
)

const (
	// spaces to save to format the percentages
	pcSp       = 7
	multidigit = 10.0
)

// This view lists out the most-consuming CPU processes.
func TopCpu(scr tcell.Screen, stt *state.State, q state.Quadrant) {
	w, h := scr.Size()
	tl, br := q.GetBounds(w, h)

	// Zero everything out
	for x := tl.X; x < br.X; x++ {
		for y := br.Y; y < tl.Y; y++ {
			scr.SetContent(x, y, space, empt, stt.ColorTheme.MainStyle)
		}
	}

	buf := 2

	lx := tl.X + buf
	ox := br.X - buf // offset x

	// height difference
	hd := (br.Y - buf) - (tl.Y - buf)

	pcs := stt.Top.Percents()

	if len(pcs) == 0 {
		return
	}

	idx := len(pcs) - 1
	for hd > 0 {
		y := br.Y - hd
		if procSl, ok := stt.Top[pcs[idx]]; ok {

			pname := procSl[0]

			for i := 0; lx+i < ox-pcSp; i++ {
				if i < len(pname) {
					scr.SetContent(lx+i, y, rune(pname[i]), empt, stt.ColorTheme.MainStyle)
				} else {
					scr.SetContent(lx+i, y, space, empt, stt.ColorTheme.MainStyle)
				}
			}

			cpu := fmtPc(pcs[idx])
			x := br.X - pcSp
			for i := 0; i < 4; i++ {
				scr.SetContent(x+i, y, rune(cpu[i]), empt, stt.ColorTheme.MainStyle)
			}

			if len(procSl) > 1 {
				//elide the entry we just drew
				stt.Top[pcs[idx]] = stt.Top[pcs[idx]][1:]
				// re-increment idx to negate the minus below &
				// return to this nameslice
				idx++
			}
		}

		idx--
		hd--
	}
}

func fmtPc(f float64) string {
	if f < multidigit {
		return fmt.Sprintf(" %.3f", f)
	}

	return fmt.Sprintf("%.3f", f)
}
