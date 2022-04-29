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
	return tcell.StyleDefault.
		Foreground(tcell.ColorBlue).
		Background(tcell.ColorBlack)
}

func CyanFg() tcell.Style {
	return AllBlack().
		Foreground(tcell.ColorLightCyan)
}

func GreenFg() tcell.Style {
	return tcell.StyleDefault.
		Foreground(tcell.ColorGreen).
		Background(tcell.ColorBlack)
}
