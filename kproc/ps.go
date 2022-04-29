package kproc

import (
	"fmt"
	"io/fs"
	"os"
)

func Top() error {
	dir, err := os.ReadDir("/proc")
	if err != nil {
		return err
	}

	procs := []fs.DirEntry{}

	for _, f := range dir {
		if '0' <= f.Name()[0] && f.Name()[0] <= '9' {
			procs = append(procs, f)
		}
	}

	for _, f := range procs {
		fmt.Print(f.Name(), " ")
	}

	fmt.Print("\n")
	return nil
}
