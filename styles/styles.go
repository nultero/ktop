package styles

import "github.com/gdamore/tcell/v2"

func AllBlack() tcell.Style {
	return tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorBlack)
}

func BlueFg() tcell.Style {
	return tcell.StyleDefault.
		Foreground(tcell.ColorBlue).
		Background(tcell.ColorBlack)
}

func CyanFg() tcell.Style {
	return tcell.StyleDefault.
		Foreground(tcell.ColorDarkCyan).
		Background(tcell.ColorBlack)
}
