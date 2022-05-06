package styles

import "github.com/gdamore/tcell/v2"

/*
	Themes package.
*/

type Theme struct {
	MainStyle      tcell.Style
	AccentStyle    tcell.Style
	HighlightStyle tcell.Style
	InactiveStyle  tcell.Style
}

func CrystalTheme() Theme {
	return Theme{
		MainStyle:      CyanFg(),
		AccentStyle:    CyanFg(),
		HighlightStyle: PinkFg(),
		InactiveStyle:  CyanFg().Foreground(tcell.ColorGray),
	}
}

func CyberPunkTheme() Theme {
	return Theme{
		MainStyle:      PinkVibrantFg().Background(tcell.ColorBlack.TrueColor()),
		AccentStyle:    CyanFg().Background(tcell.ColorBlack.TrueColor()),
		HighlightStyle: YellowFg().Background(tcell.ColorBlack.TrueColor()),
		InactiveStyle:  PinkDarkFg().Background(tcell.ColorBlack.TrueColor()),
	}
}
