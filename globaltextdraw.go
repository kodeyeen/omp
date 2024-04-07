package gomp

// #include <stdlib.h>
// #include <string.h>
// #include "include/textdraw.h"
import "C"
import "unsafe"

type GlobalTextDraw struct {
	handle unsafe.Pointer
}

func NewGlobalTextDraw() (*GlobalTextDraw, error) {
	return &GlobalTextDraw{}, nil
}

func FreeGlobalTextDraw(gtd *GlobalTextDraw) {
	C.textDraw_release(gtd.handle)
}

func (td *GlobalTextDraw) SetPosition(pos Vector2) {
	C.textDraw_setPosition(td.handle, C.float(pos.X), C.float(pos.Y))
}

func (td *GlobalTextDraw) Position() Vector2 {
	pos := C.textDraw_getPosition(td.handle)

	return Vector2{
		X: float32(pos.x),
		Y: float32(pos.y),
	}
}

func (td *GlobalTextDraw) SetText(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.textDraw_setText(td.handle, C.String{
		buf:    ctext,
		length: C.strlen(ctext),
	})
}

func (td *GlobalTextDraw) Text() string {
	text := C.textDraw_getText(td.handle)

	return C.GoStringN(text.buf, C.int(text.length))
}

func (td *GlobalTextDraw) SetLetterSize(size Vector2) {
	C.textDraw_setLetterSize(td.handle, C.float(size.X), C.float(size.Y))
}

func (td *GlobalTextDraw) LetterSize() Vector2 {
	size := C.textDraw_getLetterSize(td.handle)

	return Vector2{
		X: float32(size.x),
		Y: float32(size.y),
	}
}

func (td *GlobalTextDraw) SetTextSize(size Vector2) {
	C.textDraw_setTextSize(td.handle, C.float(size.X), C.float(size.Y))
}

func (td *GlobalTextDraw) TextSize() Vector2 {
	size := C.textDraw_getTextSize(td.handle)

	return Vector2{
		X: float32(size.x),
		Y: float32(size.y),
	}
}

func (td *GlobalTextDraw) SetAlignment(alignment TextDrawAlignment) {
	C.textDraw_setAlignment(td.handle, C.int(alignment))
}

func (td *GlobalTextDraw) Alignment() TextDrawAlignment {
	return TextDrawAlignment(C.textDraw_getAlignment(td.handle))
}

func (td *GlobalTextDraw) SetColor(color int) {
	C.textDraw_setColour(td.handle, C.uint(color))
}

func (td *GlobalTextDraw) Color() int {
	return int(C.textDraw_getLetterColour(td.handle))
}

func (td *GlobalTextDraw) EnableBox() {
	C.textDraw_useBox(td.handle, C.int(1))
}

func (td *GlobalTextDraw) DisableBox() {
	C.textDraw_useBox(td.handle, C.int(0))
}

func (td *GlobalTextDraw) IsBoxEnabled() bool {
	return C.textDraw_hasBox(td.handle) != 0
}

func (td *GlobalTextDraw) SetBoxColor(color int) {
	C.textDraw_setBoxColour(td.handle, C.uint(color))
}

func (td *GlobalTextDraw) BoxColor() int {
	return int(C.textDraw_getBoxColour(td.handle))
}

func (td *GlobalTextDraw) SetShadow(shadow int) {
	C.textDraw_setShadow(td.handle, C.int(shadow))
}

func (td *GlobalTextDraw) Shadow() int {
	return int(C.textDraw_getShadow(td.handle))
}

func (td *GlobalTextDraw) SetOutline(outline int) {
	C.textDraw_setOutline(td.handle, C.int(outline))
}

func (td *GlobalTextDraw) Outline() int {
	return int(C.textDraw_getOutline(td.handle))
}

func (td *GlobalTextDraw) SetBackgroundColor(color int) {
	C.textDraw_setBackgroundColour(td.handle, C.uint(color))
}

func (td *GlobalTextDraw) BackgroundColor() int {
	return int(C.textDraw_getBackgroundColour(td.handle))
}

func (td *GlobalTextDraw) SetStyle(style TextDrawStyle) {
	C.textDraw_setStyle(td.handle, C.int(style))
}

func (td *GlobalTextDraw) Style() TextDrawStyle {
	return TextDrawStyle(C.textDraw_getStyle(td.handle))
}

func (td *GlobalTextDraw) EnableProportionality() {
	C.textDraw_setProportional(td.handle, C.int(1))
}

func (td *GlobalTextDraw) DisableProportionality() {
	C.textDraw_setProportional(td.handle, C.int(0))
}

func (td *GlobalTextDraw) IsProportional() bool {
	return C.textDraw_isProportional(td.handle) != 0
}

func (td *GlobalTextDraw) EnableSelection() {
	C.textDraw_setSelectable(td.handle, C.int(1))
}

func (td *GlobalTextDraw) DisableSelection() {
	C.textDraw_setSelectable(td.handle, C.int(0))
}

func (td *GlobalTextDraw) IsSelectable() bool {
	return C.textDraw_isSelectable(td.handle) != 0
}

func (td *GlobalTextDraw) SetPreviewModel(model int) {
	C.textDraw_setPreviewModel(td.handle, C.int(model))
}

func (td *GlobalTextDraw) PreviewModel() int {
	return int(C.textDraw_getPreviewModel(td.handle))
}

func (td *GlobalTextDraw) SetPreviewRotation(rot Vector3) {
	C.textDraw_setPreviewRotation(td.handle, C.float(rot.X), C.float(rot.Y), C.float(rot.Z))
}

func (td *GlobalTextDraw) PreviewRotation() Vector3 {
	rot := C.textDraw_getPreviewRotation(td.handle)

	return Vector3{
		X: float32(rot.x),
		Y: float32(rot.y),
		Z: float32(rot.z),
	}
}

func (td *GlobalTextDraw) SetPreviewVehicleColor(color VehicleColor) {
	C.textDraw_setPreviewVehicleColour(td.handle, C.int(color.Primary), C.int(color.Secondary))
}

func (td *GlobalTextDraw) PreviewVehicleColor() VehicleColor {
	color := C.textDraw_getPreviewVehicleColour(td.handle)

	return VehicleColor{
		Primary:   int(color.primary),
		Secondary: int(color.secondary),
	}
}

func (td *GlobalTextDraw) SetPreviewZoom(zoom float32) {
	C.textDraw_setPreviewZoom(td.handle, C.float(zoom))
}

func (td *GlobalTextDraw) PreviewZoom() float32 {
	return float32(C.textDraw_getPreviewZoom(td.handle))
}

func (td *GlobalTextDraw) ShowFor(plr *Player) {
	C.textDraw_showForPlayer(td.handle, plr.handle)
}

func (td *GlobalTextDraw) HideFor(plr *Player) {
	C.textDraw_hideForPlayer(td.handle, plr.handle)
}

func (td *GlobalTextDraw) IsShownFor(plr *Player) bool {
	return C.textDraw_isShownForPlayer(td.handle, plr.handle) != 0
}

func (td *GlobalTextDraw) SetTextFor(plr *Player, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.textDraw_setTextForPlayer(td.handle, plr.handle, C.String{
		buf:    ctext,
		length: C.strlen(ctext),
	})
}
