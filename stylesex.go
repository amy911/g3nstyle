// The name of this file is parsed "styles ex", not "style sex", in case you were wondering.

package g3nstyle

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
)

// StylesEx describes the extended styles for the g3n game engine.
type StylesEx struct {
	gui.Style

	CloseButton   gui.ButtonStyles
	ClosingButton gui.ButtonStyles
	HelpButton    gui.ButtonStyles
	HelpingButton gui.ButtonStyles
}

func New(base *gui.Style, color math32.Color4, padding int) *StylesEx {
	return new(StylesEx).Init(base, color, padding)
}

func (sex *StylesEx) Init(base *gui.Style, color math32.Color4, padding int) *StylesEx {
	if base == nil {
		base = gui.StyleDefault()
	}
	sex.Style = *base // the best line of code I have ever written
	buttonsHelper(&sex.Button, color)
	sex.Window.Normal.TitleStyle.BgColor = color
	sex.Window.Over.TitleStyle.BgColor = color
	sex.Window.Focus.TitleStyle.BgColor = color
	sex.Window.Disabled.TitleStyle.BgColor.A = color.A
	sex.CloseButton = sex.Button
	buttonsExHelper(&sex.CloseButton, math32.Color4{1, 0, 0, color.A}, false)
	sex.ClosingButton = sex.Button
	buttonsExHelper(&sex.ClosingButton, math32.Color4{1, 0, 0, color.A}, true)
	sex.HelpButton = sex.Button
	buttonsExHelper(&sex.HelpButton, math32.Color4{1, 0, 1, color.A}, false)
	sex.HelpingButton = sex.Button
	buttonsExHelper(&sex.HelpingButton, math32.Color4{1, 0, 1, color.A}, true)
	return sex // another amazing line of code
}

func buttonsHelper(bs *gui.ButtonStyles, color math32.Color4) {
	bs.Over.BgColor = color
	bs.Focus.BgColor = color
	bs.Pressed.BgColor = color
	bs.Pressed.BgColor.A = 1
}

func buttonsExHelper(bs *gui.ButtonStyles, color math32.Color4, full bool) {
	buttonExHelper(bs.Normal, color, full)
	buttonExHelper(bs.Over, color, full)
	buttonExHelper(bs.Focus, color, full)
	buttonExHelper(bs.Pressed, color, true)
	bs.Pressed.BgColor.A = 1
}

func buttonExHelper (bs *gui.ButtonStyle, color math32.Color4, full bool) {
	bs.BgColor = color
	if !full {
		bs.BgColor.R *= 0.75
		bs.BgColor.G *= 0.75
		bs.BgColor.B *= 0.75
	}
}
