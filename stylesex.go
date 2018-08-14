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
	bs.Normal.BgColor.A = color4.A
	bs.Over.BgColor = color4
	bs.Focus.BgColor = color4
	bs.Pressed.BgColor = color4
	bs.Pressed.BgColor.A = 1
	bs.Disabled.BgColor.A = color4.A
}

func buttonsExHelper(bs *gui.ButtonStyles, color math32.Color, lit bool) {
	if !lit {
		bs.Over.BgColor.R = color.R
		bs.Over.BgColor.G = color.G
		bs.Over.BgColor.B = color.B
		bs.Focus.BgColor.R = color.R
		bs.Focus.BgColor.G = color.G
		bs.Focus.BgColor.B = color.B
		bs.Pressed.BgColor.R = color.R
		bs.Pressed.BgColor.G = color.G
		bs.Pressed.BgColor.B = color.B
		bs.Pressed.BgColor.A = 1
	} else {
		bs.Normal.BgColor.R = 0.5 * (bs.Normal.BgColor.R + color.R)
		bs.Normal.BgColor.G = 0.5 * (bs.Normal.BgColor.G + color.G)
		bs.Normal.BgColor.B = 0.5 * (bs.Normal.BgColor.B + color.B)
		bs.Normal.BgColor.A = 0.5 * (bs.Normal.BgColor.A + 1)
		bs.Over.BgColor.R = color.R
		bs.Over.BgColor.G = color.G
		bs.Over.BgColor.B = color.B
		bs.Over.BgColor.A = 1
		bs.Focus.BgColor.R = color.R
		bs.Focus.BgColor.G = color.G
		bs.Focus.BgColor.B = color.B
		bs.Focus.BgColor.A = 1
		bs.Pressed.BgColor.R = color.R
		bs.Pressed.BgColor.G = color.G
		bs.Pressed.BgColor.B = color.B
		bs.Pressed.BgColor.A = 1
	}
}
