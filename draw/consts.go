package draw

var emptyChars = []rune{}

const (
	space      = ' '
	multidigit = 10.0
	cpu        = "CPU: "
	mem        = "MEM: "
	lineChar   = '─'
)

var pips = []rune{
	' ',
	'⡀',
	'⡄',
	'⡆',
	'⡇',
	// '⣇',
	// '⣧',
	// '⣷',
	// '⣿',
}

var pipslen = float64(len(pips))
