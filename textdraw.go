package gomp

type TextDrawAlignment int

const (
	TextDrawAlignmentDefault TextDrawAlignment = iota
	TextDrawAlignmentLeft
	TextDrawAlignmentCenter
	TextDrawAlignmentRight
)

type TextDrawStyle int

const (
	TextDrawStyle0 TextDrawStyle = iota
	TextDrawStyle1
	TextDrawStyle2
	TextDrawStyle3
	TextDrawStyle4
	TextDrawStyle5
	TextDrawStyleFontBeckettRegular
	TextDrawStyleFontAharoniBold
	TextDrawStyleFontBankGothic
	TextDrawStylePricedown
	TextDrawStyleSprite
	TextDrawStylePreview
)
