package state

import "time"

func Default() (State, error) {
	s := State{}

	h := handles_t{}
	err := h.init()
	if err != nil {
		return s, err
	} else {
		s.Handles = h
	}

	s.Components = [4]Comp{
		TableOfProcs, // QUpperLeft
		IoCpu,        // QUpperRight
		Net,          // QBtmLeft
		Disk,         // QBtmRight
	}

	// redundant, but makes for a nice consumer api in /draw dir
	s.Quads = [4]Quad{0, 1, 2, 3}

	s.Cpu = cpu_t{
		Prev:           []int{},
		Cur:            []int{},
		LastCPUIdle:    0,
		LastCPUPercent: 0,
		LastSum:        0,
	}

	s.Time = time_t{
		PollRate: 800 * time.Millisecond,
	}

	return s, nil
}
