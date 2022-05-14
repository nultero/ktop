package main

import (
	"fmt"
	"ktop/kproc"
	"ktop/state"
	"ktop/styles"
	"os"
	"time"
)

const dvnull = "/dev/null"

// Checks for certain flags and otherwise prints things
// like help if those flags are present, but otherwise
// does nothing.
func parseArgs(args []string, stt *state.State) {

	//
	// TODOO obviously, set up a minimal help and a porcelain flag
	//

	funcQuits := false
	if len(args) >= 1 { // some crappy testing scaffold
		funcQuits = true
	} else {
		return
	}

	if args[0] == "t" {
		stt := state.DefaultState()
		for {
			kproc.Top(&stt)
			// pcs, pnames, err := kproc.Top(&stt)
			// if err != nil {
			// 	panic(err)
			// }
			ls := stt.Top.Percents()
			for _, f := range ls {
				fmt.Printf(
					"%.2f\t%v\t%v\n",
					f/stt.Cpu.Last(),
					stt.Top[f],
					stt.Cpu.Last(),
				)
			}

			time.Sleep(time.Second)
		}
	} else if args[0] == "--cyberpunk" {
		stt.ColorTheme = styles.CyberPunkTheme()
		funcQuits = false
	}

	b := []byte{}

	for _, arg := range args {
		b = append(b, []byte(arg)...)
	}

	os.WriteFile(dvnull, b, 0755)

	if funcQuits {
		os.Exit(0)
	}
}
