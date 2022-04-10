package main

import (
	"fmt"
	"oktop/okcpu"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

var (
	paintRate   = 500 * time.Millisecond
	idle        = 0
	cpuSum      = 0
	percentages = []float32{}
	percentsBuf = 8
)

func init() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == "t" {
			for {
				cpuPc, err := okcpu.Poll(&idle, &cpuSum)
				if err != nil {
					panic(err)
				}

				fmt.Printf("%.2f\n", cpuPc)
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

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorBlack))
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
				s.Sync()
			}
		}
	}()

	for {
		select {
		case <-quit:
			s.Fini()
			os.Exit(0)
		case <-time.After(paintRate):
		}

		cpuPc, err := okcpu.Poll(&idle, &cpuSum)
		if err != nil {
			panic(err)
		}

		drawChars(s, cpuPc)
	}
}
