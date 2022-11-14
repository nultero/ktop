package draw

import (
	"fmt"
	"ktop/state"
	"math"

	"github.com/gdamore/tcell/v2"
)

func ioCpu(scr tcell.Screen, stt *state.State, bounds [4]int) {

	lw, rw, th, bh := bounds[0], bounds[1], bounds[2], bounds[3]

	// draw the border line
	lineHeight := bh - 3
	for x := lw; x < rw; x++ {
		scr.SetContent(
			x, lineHeight, lineChar, emptyChars,
			stt.Theme.MainStyle,
		)
	}

	// draw CPU: line
	lineHeight++
	pc := formatPercent(stt.Cpu.LastCPUPercent)
	str := fmt.Sprintf("%s%s", cpu, pc)
	offset := len(str) - 1
	for x := rw - 1; offset >= 0; x-- {
		scr.SetContent(
			x, lineHeight, rune(str[offset]), emptyChars,
			stt.Theme.HighlightStyle,
		)
		offset--
	}

	// draw MEM: line
	lineHeight++
	pc = formatPercent(stt.Mem.GetLast())
	str = fmt.Sprintf("%s%s", mem, pc)
	offset = len(str) - 1
	for x := rw - 1; offset >= 0; x-- {
		scr.SetContent(
			x, lineHeight, rune(str[offset]), emptyChars,
			stt.Theme.MainStyle,
		)
		offset--
	}

	// draw the CPU pips
	pipFloor := bh - 4
	tpph := pipFloor - th // total possible pip height: floor - top height (zero is top of screen)
	tppw := rw - lw - 2   // total possible pip width
	cpuStamps := stt.Cpu.Stamps.GetLastN(tppw)
	x := rw - 1
	for idx := len(cpuStamps) - 1; idx >= 0; idx-- {

		if cpuStamps[idx] == -0.1 {
			break
		}

		pipCount := getPipsThatWillFit(cpuStamps[idx], tpph)
		pipHeight := pipFloor

	piploop:
		for {
			if pipCount >= 4 {
				scr.SetContent(
					x, pipHeight, pips[4], emptyChars,
					stt.Theme.HighlightStyle,
				)
				pipCount -= 4
			} else {
				scr.SetContent(
					x, pipHeight, pips[pipCount], emptyChars,
					stt.Theme.HighlightStyle,
				)

				// clears out any old pips above
				for pipHeight > th {
					scr.SetContent(
						x, pipHeight, space, emptyChars,
						stt.Theme.MainStyle,
					)
					pipHeight--
				}

				break piploop
			}
			pipHeight--
		}

		x--
	}
}

// Tpph passed in will be *4'd to "fill in"
// the  number of pips possible.
func getPipsThatWillFit(pc float64, tpph int) int {
	tp := float64(tpph * 4)
	pc /= 100.0
	if pc < 0.01 {
		return 0
	}
	return int(math.Round(pc * tp))
}

func formatPercent(f float64) string {
	if f < multidigit {
		return fmt.Sprintf(" %.0f%%", f)
	}

	return fmt.Sprintf("%.0f%%", f)
}
