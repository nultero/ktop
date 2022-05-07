package state

type View uint8

const (
	CpuGraph View = 0
	MemGraph View = 1
)

func (stt *State) IsFocused(v View) bool {
	return v == stt.Cursor
}
