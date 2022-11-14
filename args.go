package main

import (
	"fmt"
	"ktop/calcs"
	"ktop/proc"
	"ktop/state"
	"os"
	"time"
)

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
		for {
			proc.Collect(stt)
			calcs.Aggregate(stt)

			for _, p := range stt.SortedProcesses.Keys {
				fmt.Println(p*100.0, stt.SortedProcesses.Map[p])
			}

			time.Sleep(stt.Time.PollRate)
		}
	}

	if funcQuits {
		os.Exit(0)
	}
}
