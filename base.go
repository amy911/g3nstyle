package styles

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
)

var baseBasicStyle = gui.BasicStyle { // BasicStyle (from gui/panel.go)
	PanelStyle = basePanelStyle,
	FgColor = baseColorStyle.Text,
}

var baseBorderColor = baseColorStyle.BgDark

var baseButtonPanelStyle = gui.PanelStyle { // PanelStyle (from gui/panel.go)
	Margin: gui.RectBounds{0, 0, 0, 0},
	Border: gui.RectBounds{1, 1, 1, 1},
	Padding: gui.RectBounds{2, 4, 2, 4},
	BorderColor: baseBorderColor,
	BgColor: baseColorStyle.BgMed,
}

var baseButtonStyle = gui.ButtonStyle { // ButtonStyle (from gui/button.go)
	PanelStyle: baseButtonPanelStyle,
	FgColor: baseColorStyle.Text,
}

var baseButtonStyles = gui.ButtonStyles { // ButtonStyles (from gui/button.go)
	Normal: baseButtonStyle,
	Over: baseButtonStyle,
	Over.PanelStyle.BgColor = baseColorStyle.BgOver,
	Focus: baseButtonStyle,
	Focus.PanelStyle.BorderColor = baseColorStyle.BgOver,
	Pressed: baseButtonStyle,
	Pressed.PanelStyle.BgColor = baseColorStyle.BgOver,
	Pressed.PanelStyle.Border = gui.RectBounds{2, 2, 2, 2},
	Pressed.PanelStyle.Padding = gui.RectBounds{2, 2, 0, 4},
	Disabled: baseButtonStyle,
	Disabled.BorderColor: baseColorStyle.TextDis,
	Disabled.FgColor: baseColorStyle.TextDis,
}

var baseColorStyle = gui.ColorStyle{ // ColorStyle (from gui/style.go)
	BgDark: math32.Color4{0, 0, 0, 0.5},
	BgMed: math32.Color4{0, 0, 0, 0.375},
	BgNormal: math32.Color4{0, 0, 0, 0.25},
	BgOver: math32.Color4{0x21/255.0, 0x96/255.0, 0xF3/255.0, 0.375},
	Highlight: math32.Color4{0x82/255.0, 0xB1/255.0, 0xFF/255.0, 1},
	Select: math32.Color4{0x21/255.0, 0x96/255.0, 0xF3/255.0, 1},
	Text: math32.Color4{0xFA/255.0, 0xFA/255.0, 0xFA/255.0, 1},
	TextDis: math32.Color4{0x9E/255.0, 0x9E/255.0, 0x9E/255.0, 1},
}

var baseFontAttributes = text.FontAttributes{ // FontAttributes (from text/font.go)
	PointSize: 14, // float64
	DPI: 72, // float64
	LineSpacing: 1, // float64
	Hinting: font.HintingNone, // font.Hinting: .HintingNone, .HintingVertical, or .HintingFull
}

var basePanelStyle = gui.PanelStyle { // PanelStyle (from gui/panel.go)
	Margin: gui.RectBounds{0, 0, 0, 0},
	Border: gui.RectBounds{0, 0, 0, 0},
	Padding: gui.RectBounds{0, 0, 0, 0},
	BorderColor: baseBorderColor,
	BgColor: baseColorStyle.BgNormal,
}

var Base = gui.Style{
	Color: baseColorStyle, // ColorStyle (from gui/style.go)
	Font: nil, // *text.Font
	FontIcon: nil, // *text.Font
	Label: { // LabelStyle (from gui/label.go)
		PanelStyle: basePanelStyle,
		PanelStyle.BgColor: math32.Color4{0, 0, 0, 0},
		FontAttributes: baseFontAttributes,
		FgColor: baseColorStyle.Text,
	},
	Button: baseButtonStyles, // ButtonStyles (from gui/button.go)
	CheckRadio: { // CheckRadioStyles
	},
	Edit: { // EditStyles
	},
	ScrollBar: { // ScrollBarStyles
	},
	Slider: { // SliderStyles
	},
	Splitter: { // SplitterStyles
	},
	Window: { // WindowStyles
	},
	ItemScroller: { // ItemScrollerStyles
	},
	Scroller: { // ScrollerStyle
	},
	List: { // ListStyles
	},
	DropDown: { // DropDownStyles
	},
	Folder: { // FolderStyles
	},
	Tree: { // TreeStyles
	},
	ControlFolder: { // ControlFolderStyles
	},
	Menu: { // MenuStyles
	},
	Table: { // TableStyles
	},
	ImageButton: { // ImageButtonStyles
	},
	TabBar: { // TabBarStyles
	}
}
