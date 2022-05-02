package styles

import "github.com/gdamore/tcell/v2"

/*
	Themes package.
*/

type Theme struct {
	MainStyle     tcell.Style
	AccentStyle   tcell.Style
	InactiveStyle tcell.Style
}

func CrystalTheme() Theme {
	return Theme{
		MainStyle:     CyanFg(),
		AccentStyle:   CyanFg(),
		InactiveStyle: CyanFg().Foreground(tcell.ColorGray),
	}
}
