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
}

func NewTextdraw(text string, pos Vector2) (*Textdraw, error) {
	cText := newCString(text)
	defer freeCString(cText)

	cTd := C.textDraw_create(C.float(pos.X), C.float(pos.Y), cText)
	if cTd == nil {
		return nil, errors.New("textdraw limit was reached")
	}

	return &Textdraw{handle: cTd}, nil
}

func FreeTextdraw(td *Textdraw) {
	C.textDraw_release(td.handle)
}

func (td *Textdraw) ID() int {
	return int(C.textDraw_getID(td.handle))
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

type PlayerTextdraw struct {
	handle unsafe.Pointer
}

func NewPlayerTextdraw(plr *Player, text string, pos Vector2) (*PlayerTextdraw, error) {
	cText := newCString(text)
	defer freeCString(cText)

	cTd := C.playerTextDraw_create(plr.handle, C.float(pos.X), C.float(pos.Y), cText)
	if cTd == nil {
		return nil, errors.New("player textdraw limit was reached")
	}

	return &PlayerTextdraw{handle: cTd}, nil
}

func FreePlayerTextdraw(td *PlayerTextdraw, plr *Player) {
	C.playerTextDraw_release(td.handle, plr.handle)
}

func (td *PlayerTextdraw) ID() int {
	return int(C.playerTextDraw_getID(td.handle))
}

func (td *PlayerTextdraw) SetPosition(pos Vector2) {
	C.playerTextDraw_setPosition(td.handle, C.float(pos.X), C.float(pos.Y))
}

func (td *PlayerTextdraw) Position() Vector2 {
	pos := C.playerTextDraw_getPosition(td.handle)

	return Vector2{
		X: float32(pos.x),
		Y: float32(pos.y),
	}
}

func (td *PlayerTextdraw) SetText(text string) {
	cText := newCString(text)
	defer freeCString(cText)

	C.playerTextDraw_setText(td.handle, cText)
}

func (td *PlayerTextdraw) Text() string {
	text := C.playerTextDraw_getText(td.handle)

	return C.GoStringN(text.buf, C.int(text.length))
}

func (td *PlayerTextdraw) SetLetterSize(size Vector2) {
	C.playerTextDraw_setLetterSize(td.handle, C.float(size.X), C.float(size.Y))
}

func (td *PlayerTextdraw) LetterSize() Vector2 {
	size := C.playerTextDraw_getLetterSize(td.handle)

	return Vector2{
		X: float32(size.x),
		Y: float32(size.y),
	}
}

func (td *PlayerTextdraw) SetTextSize(size Vector2) {
	C.playerTextDraw_setTextSize(td.handle, C.float(size.X), C.float(size.Y))
}

func (td *PlayerTextdraw) TextSize() Vector2 {
	size := C.playerTextDraw_getTextSize(td.handle)

	return Vector2{
		X: float32(size.x),
		Y: float32(size.y),
	}
}

func (td *PlayerTextdraw) SetAlignment(alignment TextDrawAlignment) {
	C.playerTextDraw_setAlignment(td.handle, C.int(alignment))
}

func (td *PlayerTextdraw) Alignment() TextDrawAlignment {
	return TextDrawAlignment(C.playerTextDraw_getAlignment(td.handle))
}

func (td *PlayerTextdraw) SetColor(color Color) {
	C.playerTextDraw_setColour(td.handle, C.uint(color))
}

func (td *PlayerTextdraw) Color() Color {
	return Color(C.playerTextDraw_getLetterColour(td.handle))
}

func (td *PlayerTextdraw) EnableBox() {
	C.playerTextDraw_useBox(td.handle, 1)
}

func (td *PlayerTextdraw) DisableBox() {
	C.playerTextDraw_useBox(td.handle, 0)
}

func (td *PlayerTextdraw) IsBoxEnabled() bool {
	return C.playerTextDraw_hasBox(td.handle) != 0
}

func (td *PlayerTextdraw) SetBoxColor(color Color) {
	C.playerTextDraw_setBoxColour(td.handle, C.uint(color))
}

func (td *PlayerTextdraw) BoxColor() Color {
	return Color(C.playerTextDraw_getBoxColour(td.handle))
}

func (td *PlayerTextdraw) SetShadow(shadow int) {
	C.playerTextDraw_setShadow(td.handle, C.int(shadow))
}

func (td *PlayerTextdraw) Shadow() int {
	return int(C.playerTextDraw_getShadow(td.handle))
}

func (td *PlayerTextdraw) SetOutline(outline int) {
	C.playerTextDraw_setOutline(td.handle, C.int(outline))
}

func (td *PlayerTextdraw) Outline() int {
	return int(C.playerTextDraw_getOutline(td.handle))
}

func (td *PlayerTextdraw) SetBackgroundColor(color Color) {
	C.playerTextDraw_setBackgroundColour(td.handle, C.uint(color))
}

func (td *PlayerTextdraw) BackgroundColor() Color {
	return Color(C.playerTextDraw_getBackgroundColour(td.handle))
}

func (td *PlayerTextdraw) SetStyle(style TextdrawStyle) {
	C.playerTextDraw_setStyle(td.handle, C.int(style))
}

func (td *PlayerTextdraw) Style() TextdrawStyle {
	return TextdrawStyle(C.playerTextDraw_getStyle(td.handle))
}

func (td *PlayerTextdraw) EnableProportionality() {
	C.playerTextDraw_setProportional(td.handle, 1)
}

func (td *PlayerTextdraw) DisableProportionality() {
	C.playerTextDraw_setProportional(td.handle, 0)
}

func (td *PlayerTextdraw) IsProportional() bool {
	return C.playerTextDraw_isProportional(td.handle) != 0
}

func (td *PlayerTextdraw) EnableSelection() {
	C.playerTextDraw_setSelectable(td.handle, 1)
}

func (td *PlayerTextdraw) DisableSelection() {
	C.playerTextDraw_setSelectable(td.handle, 0)
}

func (td *PlayerTextdraw) IsSelectable() bool {
	return C.playerTextDraw_isSelectable(td.handle) != 0
}

func (td *PlayerTextdraw) SetPreviewModel(model int) {
	C.playerTextDraw_setPreviewModel(td.handle, C.int(model))
}

func (td *PlayerTextdraw) PreviewModel() int {
	return int(C.playerTextDraw_getPreviewModel(td.handle))
}

func (td *PlayerTextdraw) SetPreviewRotation(rot Vector3) {
	C.playerTextDraw_setPreviewRotation(td.handle, C.float(rot.X), C.float(rot.Y), C.float(rot.Z))
}

func (td *PlayerTextdraw) PreviewRotation() Vector3 {
	rot := C.playerTextDraw_getPreviewRotation(td.handle)

	return Vector3{
		X: float32(rot.x),
		Y: float32(rot.y),
		Z: float32(rot.z),
	}
}

func (td *PlayerTextdraw) SetPreviewVehicleColor(color VehicleColor) {
	C.playerTextDraw_setPreviewVehicleColour(td.handle, C.int(color.Primary), C.int(color.Secondary))
}

func (td *PlayerTextdraw) PreviewVehicleColor() VehicleColor {
	color := C.playerTextDraw_getPreviewVehicleColour(td.handle)

	return VehicleColor{
		Primary:   Color(color.primary),
		Secondary: Color(color.secondary),
	}
}

func (td *PlayerTextdraw) SetPreviewZoom(zoom float32) {
	C.playerTextDraw_setPreviewZoom(td.handle, C.float(zoom))
}

func (td *PlayerTextdraw) PreviewZoom() float32 {
	return float32(C.playerTextDraw_getPreviewZoom(td.handle))
}

func (td *PlayerTextdraw) Show() {
	C.playerTextDraw_show(td.handle)
}

func (td *PlayerTextdraw) Hide() {
	C.playerTextDraw_hide(td.handle)
}

func (td *PlayerTextdraw) IsShown() bool {
	return C.playerTextDraw_isShown(td.handle) != 0
}
