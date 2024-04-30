package gomp

// #include "include/textdraw.h"
import "C"
import (
	"errors"
	"unsafe"
)

type TextDrawAlignment int

const (
	TextdrawAlignmentDefault TextDrawAlignment = iota
	TextdrawAlignmentLeft
	TextdrawAlignmentCenter
	TextdrawAlignmentRight
)

type TextdrawStyle int

const (
	TextdrawStyle0 TextdrawStyle = iota
	TextdrawStyle1
	TextdrawStyle2
	TextdrawStyle3
	TextdrawStyle4
	TextdrawStyle5
	TextdrawStyleFontBeckettRegular
	TextdrawStyleFontAharoniBold
	TextdrawStyleFontBankGothic
	TextdrawStylePricedown
	TextdrawStyleSprite
	TextdrawStylePreview
)

type Textdraw struct {
	handle unsafe.Pointer
	player *Player
}

func NewTextdraw(pos Vector2, text string, plr *Player) (*Textdraw, error) {
	cText := newCString(text)
	defer freeCString(cText)

	cTd := C.textDraw_create(C.float(pos.X), C.float(pos.Y), cText)
	if cTd == nil {
		return nil, errors.New("textdraw limit reached")
	}

	return &Textdraw{handle: cTd}, nil
}

func FreeTextdraw(td *Textdraw) {
	if td.player == nil {
		C.textDraw_release(td.handle)
		return
	}

	C.playerTextDraw_release(td.handle, td.player.handle)
}

func (td *Textdraw) ID() int {
	if td.player == nil {
		return int(C.textDraw_getID(td.handle))
	}

	return int(C.playerTextDraw_getID(td.handle))
}

func (td *Textdraw) SetPosition(pos Vector2) {
	C.textDraw_setPosition(td.handle, C.float(pos.X), C.float(pos.Y))
}

func (td *Textdraw) Position() Vector2 {
	pos := C.textDraw_getPosition(td.handle)

	return Vector2{
		X: float32(pos.x),
		Y: float32(pos.y),
	}
}

func (td *Textdraw) SetText(text string) {
	cText := newCString(text)
	defer freeCString(cText)

	C.textDraw_setText(td.handle, cText)
}

func (td *Textdraw) Text() string {
	text := C.textDraw_getText(td.handle)

	return C.GoStringN(text.buf, C.int(text.length))
}

func (td *Textdraw) SetLetterSize(size Vector2) {
	C.textDraw_setLetterSize(td.handle, C.float(size.X), C.float(size.Y))
}

func (td *Textdraw) LetterSize() Vector2 {
	size := C.textDraw_getLetterSize(td.handle)

	return Vector2{
		X: float32(size.x),
		Y: float32(size.y),
	}
}

func (td *Textdraw) SetTextSize(size Vector2) {
	C.textDraw_setTextSize(td.handle, C.float(size.X), C.float(size.Y))
}

func (td *Textdraw) TextSize() Vector2 {
	size := C.textDraw_getTextSize(td.handle)

	return Vector2{
		X: float32(size.x),
		Y: float32(size.y),
	}
}

func (td *Textdraw) SetAlignment(alignment TextDrawAlignment) {
	C.textDraw_setAlignment(td.handle, C.int(alignment))
}

func (td *Textdraw) Alignment() TextDrawAlignment {
	return TextDrawAlignment(C.textDraw_getAlignment(td.handle))
}

func (td *Textdraw) SetColor(color Color) {
	C.textDraw_setColour(td.handle, C.uint(color))
}

func (td *Textdraw) Color() Color {
	return Color(C.textDraw_getLetterColour(td.handle))
}

func (td *Textdraw) EnableBox() {
	C.textDraw_useBox(td.handle, 1)
}

func (td *Textdraw) DisableBox() {
	C.textDraw_useBox(td.handle, 0)
}

func (td *Textdraw) IsBoxEnabled() bool {
	return C.textDraw_hasBox(td.handle) != 0
}

func (td *Textdraw) SetBoxColor(color Color) {
	C.textDraw_setBoxColour(td.handle, C.uint(color))
}

func (td *Textdraw) BoxColor() Color {
	return Color(C.textDraw_getBoxColour(td.handle))
}

func (td *Textdraw) SetShadow(shadow int) {
	C.textDraw_setShadow(td.handle, C.int(shadow))
}

func (td *Textdraw) Shadow() int {
	return int(C.textDraw_getShadow(td.handle))
}

func (td *Textdraw) SetOutline(outline int) {
	C.textDraw_setOutline(td.handle, C.int(outline))
}

func (td *Textdraw) Outline() int {
	return int(C.textDraw_getOutline(td.handle))
}

func (td *Textdraw) SetBackgroundColor(color Color) {
	C.textDraw_setBackgroundColour(td.handle, C.uint(color))
}

func (td *Textdraw) BackgroundColor() Color {
	return Color(C.textDraw_getBackgroundColour(td.handle))
}

func (td *Textdraw) SetStyle(style TextdrawStyle) {
	C.textDraw_setStyle(td.handle, C.int(style))
}

func (td *Textdraw) Style() TextdrawStyle {
	return TextdrawStyle(C.textDraw_getStyle(td.handle))
}

func (td *Textdraw) EnableProportionality() {
	C.textDraw_setProportional(td.handle, 1)
}

func (td *Textdraw) DisableProportionality() {
	C.textDraw_setProportional(td.handle, 0)
}

func (td *Textdraw) IsProportional() bool {
	return C.textDraw_isProportional(td.handle) != 0
}

func (td *Textdraw) EnableSelection() {
	C.textDraw_setSelectable(td.handle, 1)
}

func (td *Textdraw) DisableSelection() {
	C.textDraw_setSelectable(td.handle, 0)
}

func (td *Textdraw) IsSelectable() bool {
	return C.textDraw_isSelectable(td.handle) != 0
}

func (td *Textdraw) SetPreviewModel(model int) {
	C.textDraw_setPreviewModel(td.handle, C.int(model))
}

func (td *Textdraw) PreviewModel() int {
	return int(C.textDraw_getPreviewModel(td.handle))
}

func (td *Textdraw) SetPreviewRotation(rot Vector3) {
	C.textDraw_setPreviewRotation(td.handle, C.float(rot.X), C.float(rot.Y), C.float(rot.Z))
}

func (td *Textdraw) PreviewRotation() Vector3 {
	rot := C.textDraw_getPreviewRotation(td.handle)

	return Vector3{
		X: float32(rot.x),
		Y: float32(rot.y),
		Z: float32(rot.z),
	}
}

func (td *Textdraw) SetPreviewVehicleColor(color VehicleColor) {
	C.textDraw_setPreviewVehicleColour(td.handle, C.int(color.Primary), C.int(color.Secondary))
}

func (td *Textdraw) PreviewVehicleColor() VehicleColor {
	color := C.textDraw_getPreviewVehicleColour(td.handle)

	return VehicleColor{
		Primary:   Color(color.primary),
		Secondary: Color(color.secondary),
	}
}

func (td *Textdraw) SetPreviewZoom(zoom float32) {
	C.textDraw_setPreviewZoom(td.handle, C.float(zoom))
}

func (td *Textdraw) PreviewZoom() float32 {
	return float32(C.textDraw_getPreviewZoom(td.handle))
}

func (td *Textdraw) ShowFor(plr *Player) {
	C.textDraw_showForPlayer(td.handle, plr.handle)
}

func (td *Textdraw) HideFor(plr *Player) {
	C.textDraw_hideForPlayer(td.handle, plr.handle)
}

func (td *Textdraw) IsShownFor(plr *Player) bool {
	return C.textDraw_isShownForPlayer(td.handle, plr.handle) != 0
}

func (td *Textdraw) SetTextFor(plr *Player, text string) {
	cText := newCString(text)
	defer freeCString(cText)

	C.textDraw_setTextForPlayer(td.handle, plr.handle, cText)
}
