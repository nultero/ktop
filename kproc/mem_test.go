package kproc

import "testing"

func Test_PollMem(t *testing.T) {

	var target float32 = 34.419758

	testBytes := []byte(
		"MemTotal:       16201568 kB\n" +
			"MemAvailable:   10625028 kB",
	)

	mf, err := getMemFlt(testBytes)
	if err != nil {
		t.Errorf("error getting membytes: %v", err)
	}

	if mf != target {
		t.Errorf("wanted '%v' for membytes, got '%v'", target, mf)
	}
}
