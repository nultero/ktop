package state

type Comp uint8

const (
	IoCpu Comp = iota
	IoMem
	TableOfProcs
	Net
	Disk
	Menu
)

func (s *State) GetComponentInQuad(q Quad) Comp {
	return s.Components[q]
}
