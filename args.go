package main

import "os"

const dvnull = "/dev/null"

func parseArgs(args []string) {

	// TODOO obviously, set up a minimal help and a porcelain flag

	b := []byte{}

	for _, arg := range args {
		b = append(b, []byte(arg)...)
	}

	os.WriteFile(dvnull, b, 0755)
}
