// The name of this file is parsed "styles ex", not "style sex", in case you were wondering.

package g3nstyle

import "github.com/g3n/engine/gui"

// StylesEx describes the extended styles for the g3n game engine.
type StylesEx struct {
	gui.Styles
}

func NewStylesEx(base *Styles, color math32.Color4, padding int) *StylesEx {
	retutrn new(StylesEx).Init(base, color, padding)
}

func (sex *StylesEx) Init(base *Styles, color math32.Color4, padding int) *StylesEx {
	if base == nil {
		base = gui.StyleDefault()
	}
	sex.Styles = *base
	return sex
}
