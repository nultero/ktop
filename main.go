package main

import (
	"ktop/calcs"
	"ktop/draw"
	"ktop/proc"
	"ktop/state"
	"ktop/styles"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

/*
	~~TODO~~ config somewhere
	Instead of config, maybe have a large
	args list?
	That way they can just be aliased, no file
	wonkery needed
*/

const minX, minY = 30, 16

func main() {

	stt, err := state.Default()
	if err != nil {
		panic(err)
	}
	stt.Theme = styles.CrystalTheme()

	parseArgs(os.Args[1:], &stt)

	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err = screen.Init(); err != nil {
		panic(err)
	}

	// TODO ? mesh func for screen + state

	screen.HideCursor()
	screen.SetStyle(stt.Theme.MainStyle)
	screen.Clear()

	quit := make(chan struct{})
	rerender := make(chan struct{})
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {

				// TODO : also Vim bindings here
				case tcell.KeyDown, tcell.KeyUp, tcell.KeyRight, tcell.KeyLeft:
					// TODO stt.Move method here
					rerender <- struct{}{}

				case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC, tcell.KeyCtrlQ:
					quit <- struct{}{}
					close(quit)
					return

				case tcell.KeyCtrlL:
					screen.Sync()
				}

			case *tcell.EventResize:
				screen.Sync()
				rerender <- struct{}{}
			}
		}
	}()

	// background collector loop
	go func() {
		for {
			proc.Collect(&stt)
			calcs.Aggregate(&stt)
			time.Sleep(stt.Time.PollRate)
		}
	}()

loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(stt.Time.PollRate):
		case <-rerender:
		}

		if isDrawable(screen.Size()) {
			// draw.Cpu(screen, stt)
			draw.Draw(screen, &stt)
		} else {
			draw.Invalid(screen, stt.Theme.MainStyle)
		}

		screen.Show() // only calling this once ✓
	}

	screen.Fini()
}

func isDrawable(x, y int) bool {
	return x > minX || y > minY
}
