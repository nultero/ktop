package kproc

import (
	"testing"
)

// live example of a proc/pid/stat file taken from my zsh shell process
// the lone '4 1' of the first line are the relevant bits
// more: https://man7.org/linux/man-pages/man5/proc.5.html
var zshProc = []string{
	"13812 (zsh) S 13811 13812 13812 34818 13812 4194304 2358 10572 0 1 4 1 10 5 20",
	"0 1 0 1492334 13103104 1503 18446744073709551615 94170633760768 94170634535281",
	"140735486853872 0 0 0 0 3686404 134295555 1 0 0 17 4 0 0 0 0 0 94170634652336 9",
	"4170634681580 94170641149952 140735486860914 140735486860919 140735486860919 14",
	"0735486865387 0"}

var bytes = []byte(zshProc[0] + zshProc[1]) // about the same size as my byte buffer

func Test_parseStat(t *testing.T) {

	const (
		z = "zsh"

		t_utime int64 = 4
		t_stime int64 = 1
	)

	zsh, utime, stime := parseStat(bytes)

	if z != zsh {
		t.Errorf("wanted: %v, got: %v", z, zsh)
	}

	if t_utime != utime {
		t.Errorf("wanted: %v, got: %v", t_utime, utime)
	}

	if t_stime != stime {
		t.Errorf("wanted: %v, got: %v", t_stime, stime)
	}

}
