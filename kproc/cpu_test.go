package kproc

import (
	"testing"
)

func Test_cpuBytes(t *testing.T) {
	bytes := []byte("cpu  676058 748 194770 10723540 5012 0 7822 0 0 0")

	nums, err := readBytes(bytes)
	if err != nil {
		t.Errorf("errored: %v", err)
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	w := 11607950
	if sum != w {
		t.Errorf("wanted %d, got %d", w, sum)
	}
}
