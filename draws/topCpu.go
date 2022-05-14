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
	// onQ := stt.OnQuad(q)

	buf := 2

	lx := tl.X + buf
	ox := br.X - buf // offset x

	// height difference
	hd := (br.Y - buf) - (tl.Y - buf)

	pcs := stt.Top.Percents()

	if len(pcs) == 0 {
		return
	}

	for i := len(pcs) - 1; hd > 0; {
		y := br.Y - hd
		if procSl, ok := stt.Top[pcs[i]]; ok {
			pname := procSl[0]

			for i := 0; lx+i < ox-pcSp; i++ {
				if i < len(pname) {
					scr.SetContent(lx+i, y, rune(pname[i]), empt, stt.ColorTheme.MainStyle)
				} else {
					scr.SetContent(lx+i, y, space, empt, stt.ColorTheme.MainStyle)
				}
			}

			cpu := fmtPc(100.0 * pcs[i])
			x := br.X - pcSp
			for i := 0; i < 4; i++ {
				scr.SetContent(x+i, y, rune(cpu[i]), empt, stt.ColorTheme.MainStyle)
			}

			if len(procSl) > 1 {
				//elide the entry we just drew
				stt.Top[pcs[i]] = stt.Top[pcs[i]][1:]
				// re-increment i to negate the minus below &
				// return to this nameslice
				i++
			}
		}

		i--
		hd--
	}

	// for idx, r := range memstr {
	// 	x := ox + idx - len(memstr)
	// 	if stt.IsFocused(state.MemGraph) {
	// 		if onQ && idx == 0 {
	// 			scr.SetContent(x, br.Y, r, empt, stt.ColorTheme.HighlightStyle)

	// 		} else {
	// 			scr.SetContent(x, br.Y, r, empt, stt.ColorTheme.MainStyle)
	// 		}
	// 	} else {
	// 		scr.SetContent(x, br.Y, r, empt, stt.ColorTheme.InactiveStyle)
	// 	}
	// }

	// for idx, r := range cpustr {
	// 	x := ox + idx - len(cpustr)
	// 	if stt.IsFocused(state.CpuGraph) {
	// 		if onQ {
	// 			// if onQ && idx == 0 {
	// 			scr.SetContent(x, br.Y-1, r, empt, stt.ColorTheme.HighlightStyle)
	// 		} else {
	// 			scr.SetContent(x, br.Y-1, r, empt, stt.ColorTheme.MainStyle)
	// 		}
	// 	} else {
	// 		scr.SetContent(x, br.Y-1, r, empt, stt.ColorTheme.InactiveStyle)

	// 	}
	// }

	// // bottom of graph line
	// for i := ox; i > tl.X; i-- {
	// 	scr.SetContent(i, br.Y-2, '─', empt, stt.ColorTheme.InactiveStyle)
	// }

	// /*
	// 	TODOOO -- need focused conditional to draw CPU vs. Mem
	// 	Graph draws:
	// */

	// xdiff := (ox - 1) - (tl.X + 1)
	// subh := br.Y - tl.Y - 2
	// charh := subh * 4 // chars' possible heights within the quadrant

	// if len(stt.Cpu.Stamps) == 0 {
	// 	return
	// }

	// for i := len(stt.Cpu.Stamps) - 1; i > 0; i-- {

	// 	if xdiff == 0 {
	// 		break
	// 	}

	// 	x := ox - (len(stt.Cpu.Stamps) - i)
	// 	y := br.Y - 3

	// 	dots := int(math.Round(
	// 		float64(charh) * stt.Cpu.Stamps[i] / 100.0,
	// 	))

	// 	sty := stt.ColorTheme.MainStyle
	// 	if onQ {
	// 		sty = stt.ColorTheme.AccentStyle
	// 	}

	// 	for y > tl.Y+1 {
	// 		if dots == 0 {
	// 			scr.SetContent(x, y, space, empt, sty)

	// 		} else if dots >= 4 {
	// 			scr.SetContent(x, y, dotrunes[4], empt, sty)
	// 			dots -= 4

	// 		} else {
	// 			scr.SetContent(x, y, dotrunes[dots], empt, sty)
	// 			dots = 0
	// 		}

	// 		y--
	// 	}

	// 	xdiff--
	// }
}

func fmtPc(f float64) string {
	if f < multidigit {
		return fmt.Sprintf(" %.3f", f)
	}

	return fmt.Sprintf("%.3f", f)
}
