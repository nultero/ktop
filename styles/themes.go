package styles

import "github.com/gdamore/tcell/v2"

/*
	Themes package.
*/

type Theme struct {
	MainStyle      tcell.Style
	SecondaryStyle tcell.Style
	AccentStyle    tcell.Style
	HighlightStyle tcell.Style
	InactiveStyle  tcell.Style
}

func CrystalTheme() Theme {
	return Theme{
		MainStyle:      CyanFg(),
		AccentStyle:    CrystalFocus(),
		HighlightStyle: CrystalFocus(),
		InactiveStyle:  CyanFg().Foreground(tcell.ColorGray),
	}
}

func CyberPunkTheme() Theme {
	return Theme{
		MainStyle:      PinkVibrantFg().Background(tcell.ColorBlack.TrueColor()),
		AccentStyle:    PinkDarkFg().Background(tcell.ColorBlack.TrueColor()),
		HighlightStyle: YellowFg().Background(tcell.ColorBlack.TrueColor()),
		InactiveStyle:  PinkDarkFg().Background(tcell.ColorBlack.TrueColor()),
	}
}
