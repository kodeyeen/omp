package gomp

type TextDraw interface {
	SetPosition(pos Vector3)
	Position() Vector3
	SetText(text string)
	Text() string
	SetLetterSize()
	LetterSize() Vector2
	SetTextSize(size Vector2)
	TextSize() Vector2
	SetAlignment(alignment int)
	Alignment() int
	SetColor(color int)
	Color() int
	EnableBox()
	DisableBox()
	IsBoxEnabled() bool
	SetBoxColor(color int)
	BoxColor() int
	SetShadow(shadow int)
	Shadow() int
	SetOutline(size int)
	Outline() int
	SetBackgroundColor(color int)
	BackgroundColor() int
	SetStyle(style int)
	Style() int
	EnableProportionality()
	DisableProportionality()
	IsProportional() bool
	EnableSelection()
	DisableSelection()
	IsSelectable() bool
	SetPreviewModel(model int)
	PreviewModel() int
	SetPreviewRotation(rot Vector3)
	PreviewRotation() Vector3
	SetPreviewVehicleColor(color VehicleColor)
	PreviewVehicleColor() VehicleColor
	SetPreviewZoom()
	PreviewZoom() float32
	Show()
	Hide()
	IsShown() bool
}
