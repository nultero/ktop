package main

import (
	"fmt"
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

// TODO porcelain flag for logging metrics

func main() {

	args := os.Args[1:]

	if len(args) > 0 {
		if args[0] == "t" {
			for {
				cpuPc, err := kproc.PollCPU(&idle, &cpuSum)
				if err != nil {
					panic(err)
				}

				mem, err := kproc.PollMem()
				if err != nil {
					panic(err)
				}

				err = kproc.Top()
				if err != nil {
					panic(err)
				}

				fmt.Printf("cpu: %.2f\nmeminfo: %.2f\n\n", cpuPc, mem)
				time.Sleep(paintRate)
			}

		}
		return
	}

	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err = s.Init(); err != nil {
		panic(err)
	}

	s.HideCursor()
	s.SetStyle(styles.Blk())
	s.Clear()

	quit := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC, tcell.KeyCtrlQ:
					close(quit)
					return
				case tcell.KeyCtrlL:
					s.Sync()
				}
			case *tcell.EventResize:
				needsRedraw = true
				s.Sync()
			}
		}
	}()

	sty := styles.CyanFg()

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
			redraw(s, styles.AllBlack())
			needsRedraw = false
		}

		drawChars(s, cpuPc, mem, sty)
	}

	s.Fini()
}
