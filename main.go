package main

import (
	"ktop/kproc"
	"ktop/styles"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

var (
	paintRate   = 500 * time.Millisecond
	idle        = 0
	cpuSum      = 0
	percentages = []float32{}
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
	cpuStamps := []float32{}

renderloop:
	for {
		select {
		case <-quit:
			break renderloop

		case <-time.After(paintRate):
		}

		cpuPc, err := kproc.PollCPU(&idle, &cpuSum)
		if err != nil {
			panic(err)
		}

		mem, err := kproc.PollMem()
		if err != nil {
			panic(err)
		}

		if needsRedraw {
			redraw(screen, styles.AllBlack())
			needsRedraw = false
		}

		cpuStamps = append(cpuStamps, cpuPc)
		if len(cpuStamps) > 30 { // TODO arbitrary stamp lens for now
			cpuStamps = cpuStamps[1:]
		}

		if isDrawable(screen.Size()) {
			stdDraw(screen, cpuPc, mem, sty, cpuStamps)
		} else {
			invalidSzDraw(screen, sty)
		}
	}

	screen.Fini()
}

func isDrawable(x, y int) bool {
	if x < 80 || y < 24 {
		return false
	}

	return true
}
