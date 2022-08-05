package state

type proc_t struct {
	cpuPc float64
	name  string
	utime [2]int64 // user mode jiffies
	stime [2]int64 // kernel mode jiffies
}

func (pt proc_t) Prev() int64 {
	return pt.utime[0] + pt.stime[0]
}

func (pt proc_t) Cur() int64 {
	return pt.utime[1] + pt.stime[1]
}

func (pt proc_t) Utime() int {
	return int(pt.utime[1] - pt.utime[0])
}

func (pt proc_t) Stime() int {
	return int(pt.stime[1] - pt.stime[0])
}

func (pt proc_t) Name() string {
	return pt.name
}
