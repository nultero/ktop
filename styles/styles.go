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

func CrystalFocus() tcell.Style {
	c := tcell.NewRGBColor(69, 237, 226)
	return AllBlack().
		Foreground(c)
}

func CyanFg() tcell.Style {
	return AllBlack().
		Foreground(tcell.ColorLightCyan)
}

func GreenFg() tcell.Style {
	return AllBlack().
		Foreground(tcell.ColorGreen)
}

func MagentaFg() tcell.Style {
	return AllBlack().
		Foreground(tcell.ColorDarkMagenta)
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

func PinkVibrantFg() tcell.Style {
	return AllBlack().
		Foreground(tcell.ColorDeepPink)
}

func PinkDarkFg() tcell.Style {
	c := tcell.NewRGBColor(120, 0, 67)
	return AllBlack().
		Foreground(c)
}

func LightYellowFg() tcell.Style {
	return AllBlack().
		Foreground(tcell.ColorLightYellow)
}

func YellowFg() tcell.Style {
	return AllBlack().
		Foreground(tcell.ColorYellow)
}
