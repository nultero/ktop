package main

import (
	"ktop/kproc"
	"ktop/state"
	"os"
	"time"
)

const dvnull = "/dev/null"

// Checks for certain flags and otherwise prints things
// like help if those flags are present, but otherwise
// does nothing.
func parseArgs(args []string) {

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
		for i := 0; i < 3; i++ {
			err := kproc.Top(&stt)
			if err != nil {
				panic(err)
			}

			time.Sleep(time.Second)
		}
		stt.Log.Dump()
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
