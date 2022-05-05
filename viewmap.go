package main

import "ktop/ktdata"

type viewmap map[string]ktdata.View

func (v viewmap) isFocused(s string, stt *ktdata.State) bool {
	if view, ok := v[s]; ok {
		return view == stt.Focused
	}

	return false
}

var vmp = viewmap{
	cpuTxt: 1,
	memTxt: 2,
}
