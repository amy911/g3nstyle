// The name of this file is parsed "styles ex", not "style sex", in case you were wondering.

package g3nstyle

import "github.com/g3n/engine/gui"

// StylesEx describes the extended styles for the g3n game engine.
type StylesEx struct {
	gui.Style

	CloseButton   gui.ButtonStyle
	ClosingButton gui.ButtonStyle
	HelpButton    gui.ButtonStyle
	HelpingButton gui.ButtonStyle
}

func New(base *gui.Style, color math32.Color4, padding int) *StylesEx {
	return new(StylesEx).Init(base, color, padding)
}

func (sex *StylesEx) Init(base *gui.Style, color math32.Color4, padding int) *StylesEx {
	if base == nil {
		base = gui.StyleDefault()
	}
	sex.Style = *base // the best line of code I have ever written
	sex.Button.Over.BgColor = color
	sex.Window.TitleStyle.BgColor = color
	sex.CloseButton = sex.Button
	sex.CloseButton.BgColor = math32.Color4{1, 0, 0, 1}
	sex.ClosingButton = sex.Button
	sex.ClosingButton. = math32.Color4{1, 0, 0, 1}
	sex.HelpButton = sex.Button
	sex.HelpButton. = math32.Color4{1, 0, 1, 1}
	sex.HelpingButton = sex.Button
	sex.HelpingButton. = math32.Color4{1, 0, 1, 1}
	return sex // another amazing line of code
}
