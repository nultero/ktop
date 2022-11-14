package state

type cpu_t struct {
	Prev           []int // Slice of the previous stat's jiffies.
	Cur            []int // Slice of the current stat's jiffies.
	DiffSum        int
	LastCPUIdle    int
	LastCPUPercent float64
	LastSum        int
	LastSumNoIdle  int
	Stamps         stamps_t
}

// Pushes current to previous and sets current to newest
// file data from /proc/stat.
func (c *cpu_t) Add(nums []int) {
	c.Prev = c.Cur
	c.Cur = nums
}
