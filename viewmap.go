package main

import "ktop/state"

type viewmap map[string]state.View

func (v viewmap) isFocused(s string, stt *state.State) bool {
	if view, ok := v[s]; ok {
		return view == stt.FocusedComp
	}

	return false
}

var vmp = viewmap{
	cpuTxt: 1,
	memTxt: 2,
}
