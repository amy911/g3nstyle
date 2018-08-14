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

func New(base *gui.Style, color *math32.Color, alpha float32, padding int) *StylesEx {
	return new(StylesEx).Init(base, color, alpha, padding)
}

func (sex *StylesEx) Init(base *gui.Style, color *math32.Color, alpha float32, padding int) *StylesEx {
	if base == nil {
		base = gui.StyleDefault()
	}
	color4 := math32.Color4{0x21/255.0, 0x96/255.0, 0xF3/255.0, alpha}
	if color != nil {
		color4.R = color.R
		color4.G = color.G
		color4.B = color.B
	}
	sex.Style = *base // the best line of code I have ever written
	buttonsHelper(&sex.Button, color4)
	sex.Window.Normal.TitleStyle.BgColor = color4
	sex.Window.Over.TitleStyle.BgColor = color4
	sex.Window.Focus.TitleStyle.BgColor = color4
	sex.Window.Disabled.TitleStyle.BgColor.A = color4.A
	sex.CloseButton = sex.Button
	buttonsExHelper(&sex.CloseButton, math32.Color4{1, 0, 0, color4.A}, false)
	sex.ClosingButton = sex.Button
	buttonsExHelper(&sex.ClosingButton, math32.Color4{1, 0, 0, color4.A}, true)
	sex.HelpButton = sex.Button
	buttonsExHelper(&sex.HelpButton, math32.Color4{1, 0, 1, color4.A}, false)
	sex.HelpingButton = sex.Button
	buttonsExHelper(&sex.HelpingButton, math32.Color4{1, 0, 1, color4.A}, true)
	return sex // another amazing line of code
}

func buttonsHelper(bs *gui.ButtonStyles, color4 math32.Color4) {
	bs.Over.BgColor = color4
	bs.Focus.BgColor = color4
	bs.Pressed.BgColor = color4
	bs.Pressed.BgColor.A = 1
}

func buttonsExHelper(bs *gui.ButtonStyles, color4 math32.Color4, full bool) {
	buttonExHelper(&bs.Normal, color4, full)
	buttonExHelper(&bs.Over, color4, full)
	buttonExHelper(&bs.Focus, color4, full)
	buttonExHelper(&bs.Pressed, color4, true)
	bs.Pressed.BgColor.A = 1
}

func buttonExHelper (bs *gui.ButtonStyle, color4 math32.Color4, full bool) {
	bs.BgColor = color4
	if !full {
		bs.BgColor.R *= 0.75
		bs.BgColor.G *= 0.75
		bs.BgColor.B *= 0.75
	}
}
