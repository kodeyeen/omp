package gomp

// #include <stdlib.h>
// #include <string.h>
// #include "include/playertextdraw.h"
import "C"
import "unsafe"

type PlayerTextDraw struct {
	handle unsafe.Pointer
}

func NewPlayerTextDraw(plr Player, pos Vector2, text string) *PlayerTextDraw {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	plrtd := C.playerTextDraw_create(plr.handle, C.float(pos.X), C.float(pos.Y), C.String{
		buf:    ctext,
		length: C.strlen(ctext),
	})

	return &PlayerTextDraw{handle: plrtd}
}

func FreePlayerTextDraw(plrtd *PlayerTextDraw, plr Player) {
	C.playerTextDraw_release(plrtd.handle, plr.handle)
}

func (td *PlayerTextDraw) SetPosition(pos Vector2) {
	C.playerTextDraw_setPosition(td.handle, C.float(pos.X), C.float(pos.Y))
}

func (td *PlayerTextDraw) Position() Vector2 {
	pos := C.playerTextDraw_getPosition(td.handle)

	return Vector2{
		X: float32(pos.x),
		Y: float32(pos.y),
	}
}

func (td *PlayerTextDraw) SetText(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.playerTextDraw_setText(td.handle, C.String{
		buf:    ctext,
		length: C.strlen(ctext),
	})
}

func (td *PlayerTextDraw) Text() string {
	text := C.playerTextDraw_getText(td.handle)

	return C.GoStringN(text.buf, C.int(text.length))
}

func (td *PlayerTextDraw) SetLetterSize(size Vector2) {
	C.playerTextDraw_setLetterSize(td.handle, C.float(size.X), C.float(size.Y))
}

func (td *PlayerTextDraw) LetterSize() Vector2 {
	size := C.playerTextDraw_getLetterSize(td.handle)

	return Vector2{
		X: float32(size.x),
		Y: float32(size.y),
	}
}

func (td *PlayerTextDraw) SetTextSize(size Vector2) {
	C.playerTextDraw_setTextSize(td.handle, C.float(size.X), C.float(size.Y))
}

func (td *PlayerTextDraw) TextSize() Vector2 {
	size := C.playerTextDraw_getTextSize(td.handle)

	return Vector2{
		X: float32(size.x),
		Y: float32(size.y),
	}
}

func (td *PlayerTextDraw) SetAlignment(alignment TextDrawAlignment) {
	C.playerTextDraw_setAlignment(td.handle, C.int(alignment))
}

func (td *PlayerTextDraw) Alignment() TextDrawAlignment {
	return TextDrawAlignment(C.playerTextDraw_getAlignment(td.handle))
}

func (td *PlayerTextDraw) SetColor(color int) {
	C.playerTextDraw_setColour(td.handle, C.uint(color))
}

func (td *PlayerTextDraw) Color() int {
	return int(C.playerTextDraw_getLetterColour(td.handle))
}

func (td *PlayerTextDraw) EnableBox() {
	C.playerTextDraw_useBox(td.handle, C.int(1))
}

func (td *PlayerTextDraw) DisableBox() {
	C.playerTextDraw_useBox(td.handle, C.int(0))
}

func (td *PlayerTextDraw) IsBoxEnabled() bool {
	return C.playerTextDraw_hasBox(td.handle) != 0
}

func (td *PlayerTextDraw) SetBoxColor(color int) {
	C.playerTextDraw_setBoxColour(td.handle, C.uint(color))
}

func (td *PlayerTextDraw) BoxColor() int {
	return int(C.playerTextDraw_getBoxColour(td.handle))
}

func (td *PlayerTextDraw) SetShadow(shadow int) {
	C.playerTextDraw_setShadow(td.handle, C.int(shadow))
}

func (td *PlayerTextDraw) Shadow() int {
	return int(C.playerTextDraw_getShadow(td.handle))
}

func (td *PlayerTextDraw) SetOutline(outline int) {
	C.playerTextDraw_setOutline(td.handle, C.int(outline))
}

func (td *PlayerTextDraw) Outline() int {
	return int(C.playerTextDraw_getOutline(td.handle))
}

func (td *PlayerTextDraw) SetBackgroundColor(color int) {
	C.playerTextDraw_setBackgroundColour(td.handle, C.uint(color))
}

func (td *PlayerTextDraw) BackgroundColor() int {
	return int(C.playerTextDraw_getBackgroundColour(td.handle))
}

func (td *PlayerTextDraw) SetStyle(style int) {
	C.playerTextDraw_setStyle(td.handle, C.int(style))
}

func (td *PlayerTextDraw) Style() int {
	return int(C.playerTextDraw_getStyle(td.handle))
}

func (td *PlayerTextDraw) EnableProportionality() {
	C.playerTextDraw_setProportional(td.handle, C.int(1))
}

func (td *PlayerTextDraw) DisableProportionality() {
	C.playerTextDraw_setProportional(td.handle, C.int(0))
}

func (td *PlayerTextDraw) IsProportional() bool {
	return C.playerTextDraw_isProportional(td.handle) != 0
}

func (td *PlayerTextDraw) EnableSelection() {
	C.playerTextDraw_setSelectable(td.handle, C.int(1))
}

func (td *PlayerTextDraw) DisableSelection() {
	C.playerTextDraw_setSelectable(td.handle, C.int(0))
}

func (td *PlayerTextDraw) IsSelectable() bool {
	return C.playerTextDraw_isSelectable(td.handle) != 0
}

func (td *PlayerTextDraw) SetPreviewModel(model int) {
	C.playerTextDraw_setPreviewModel(td.handle, C.int(model))
}

func (td *PlayerTextDraw) PreviewModel() int {
	return int(C.playerTextDraw_getPreviewModel(td.handle))
}

func (td *PlayerTextDraw) SetPreviewRotation(rot Vector3) {
	C.playerTextDraw_setPreviewRotation(td.handle, C.float(rot.X), C.float(rot.Y), C.float(rot.Z))
}

func (td *PlayerTextDraw) PreviewRotation() Vector3 {
	rot := C.playerTextDraw_getPreviewRotation(td.handle)

	return Vector3{
		X: float32(rot.x),
		Y: float32(rot.y),
		Z: float32(rot.z),
	}
}

func (td *PlayerTextDraw) SetPreviewVehicleColor(color VehicleColor) {
	C.playerTextDraw_setPreviewVehicleColour(td.handle, C.int(color.Primary), C.int(color.Secondary))
}

func (td *PlayerTextDraw) PreviewVehicleColor() VehicleColor {
	color := C.playerTextDraw_getPreviewVehicleColour(td.handle)

	return VehicleColor{
		Primary:   int(color.primary),
		Secondary: int(color.secondary),
	}
}

func (td *PlayerTextDraw) SetPreviewZoom(zoom float32) {
	C.playerTextDraw_setPreviewZoom(td.handle, C.float(zoom))
}

func (td *PlayerTextDraw) PreviewZoom() float32 {
	return float32(C.playerTextDraw_getPreviewZoom(td.handle))
}

func (td *PlayerTextDraw) Show() {
	C.playerTextDraw_show(td.handle)
}

func (td *PlayerTextDraw) Hide() {
	C.playerTextDraw_hide(td.handle)
}

func (td *PlayerTextDraw) IsShown() bool {
	return C.playerTextDraw_isShown(td.handle) != 0
}
