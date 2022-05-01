package main

import (
	"ktop/kproc"
	"ktop/styles"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

/*
	TODO config somewhere
*/

var (
	paintRate   = 500 * time.Millisecond
	ramPollRate = 30 * time.Second
	lci         = 0 // last CPU idle %, used for PollCPU
	cpuSum      = 0
	cpuStamps   = []float32{}
	ramStamps   = []float32{}

	needsRedraw = false
)

func init() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
}

func main() {

	parseArgs(os.Args[1:])

	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err = screen.Init(); err != nil {
		panic(err)
	}

	screen.HideCursor()
	screen.SetStyle(styles.Blk())
	screen.Clear()

	quit := make(chan struct{})
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC, tcell.KeyCtrlQ:
					close(quit)
					return
				case tcell.KeyCtrlL:
					screen.Sync()
				}
			case *tcell.EventResize:
				needsRedraw = true
				screen.Sync()
			}
		}
	}()

	sty := styles.CyanFg()
	inactiveSty := sty.Foreground(tcell.ColorGray)
	// sty := styles.Matrix()

renderloop:
	for {
		select {
		case <-quit:
			break renderloop

		case <-time.After(paintRate):
		}

		cpuPc, err := kproc.PollCPU(&lci, &cpuSum)
		if err != nil {
			panic(err)
		}

		mem, err := kproc.PollMem()
		if err != nil {
			panic(err)
		}

		if needsRedraw { // TODO fix drawing twice, redraw + stdDraw is bad
			redraw(screen, sty)
			needsRedraw = false
		}

		cpuStamps = append(cpuStamps, cpuPc)
		if len(cpuStamps) > 30 { // TODO arbitrary stamp lens for now; clean up later
			cpuStamps = cpuStamps[1:]
		}

		if isDrawable(screen.Size()) {
			stdDraw(screen, cpuPc, mem, sty, inactiveSty, cpuStamps)
		} else {
			invalidSzDraw(screen, sty)
		}
	}

	screen.Fini()
}

func isDrawable(x, y int) bool {
	if x < 30 || y < 16 {
		return false
	}

	return true
}
