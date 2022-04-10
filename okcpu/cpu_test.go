package okcpu

import (
	"testing"
)

func Test_sumBytes(t *testing.T) {
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

func Test_trunc(t *testing.T) {
	nums := []float64{
		9.25,
		7.920792,
		6.5491185, // fails here, might just use strfmt instead
	}

	truncs := []float64{
		0.1,
		0.1,
		0.05,
	}

	vals := []float64{
		9.3,
		7.9,
		6.55,
	}

	for i, n := range nums {
		testn := trunc(n, truncs[i])
		if testn != vals[i] {
			t.Errorf("wanted %v, got %v", vals[i], testn)
		}
	}
}
