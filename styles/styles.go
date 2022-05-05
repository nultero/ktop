package styles

import "github.com/gdamore/tcell/v2"

func AllBlack() tcell.Style {
	return tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorBlack)
}

func Blk() tcell.Style {
	return tcell.StyleDefault.
		Foreground(tcell.ColorBlack.TrueColor()).
		Background(tcell.ColorBlack.TrueColor())
}

func BlueFg() tcell.Style {
	return AllBlack().
		Foreground(tcell.ColorBlue)
}

func CyanFg() tcell.Style {
	return AllBlack().
		Foreground(tcell.ColorLightCyan)
}

func GreenFg() tcell.Style {
	return AllBlack().
		Foreground(tcell.ColorGreen)
}

func Matrix() tcell.Style {
	c := tcell.NewRGBColor(0, 210, 17)
	return tcell.StyleDefault.
		Foreground(c).
		Background(tcell.ColorBlack.TrueColor())
}

func InvalidRed() tcell.Style {
	return AllBlack().
		Foreground(tcell.ColorRed)
}

func PinkFg() tcell.Style {
	return AllBlack().
		Foreground(tcell.ColorHotPink)
}
