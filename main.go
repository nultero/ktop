package main

import (
	"ktop/draw"
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

func init() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
}

func main() {

	stt, err := state.Default()
	if err != nil {
		panic(err)
	}

	// err = kproc.Top(&stt)
	// if err != nil {
	// 	panic(err)
	// }
	stt.Theme = styles.CrystalTheme()

	// maybe leave args here to overwrite default state
	parseArgs(os.Args[1:], &stt)

	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err = screen.Init(); err != nil {
		panic(err)
	}

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

				case tcell.KeyDown, tcell.KeyUp, tcell.KeyRight, tcell.KeyLeft:
					// stt.MoveQuad(ev.Key())
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

mainloop:
	for {
		select {
		case <-quit:
			break mainloop

		case <-rerender:

		case <-time.After(stt.Time.PollRate):
			draw.Cpu(screen, stt)
		}

		// if isDrawable(screen.Size()) {
		// 	draws.TopCpu(screen, &stt, state.QuadTopLeft)
		// 	draws.Io(screen, &stt, state.QuadTopRight)

		// 	// TODO - bottom left is netwk?
		// 	draws.Io(screen, &stt, state.QuadBottomRight)
		// 	draws.Io(screen, &stt, state.QuadBottomLeft)
		// } else {
		// 	draws.Invalid(screen, stt.ColorTheme.MainStyle, minX, minY)
		// }

		screen.Show() // only calling this once ✓
	}

	screen.Fini()
}

func isDrawable(x, y int) bool {
	if x < minX || y < minY {
		return false
	}

	return true
}
