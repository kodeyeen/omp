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

	plrtd := C.playerTextDraw_create(plr.Handle(), C.float(pos.X), C.float(pos.Y), C.String{
		buf:    ctext,
		length: C.strlen(ctext),
	})

	return &PlayerTextDraw{handle: plrtd}
}

func FreePlayerTextDraw(plrtd *PlayerTextDraw, plr Player) {
	C.playerTextDraw_release(plrtd.Handle(), plr.Handle())
}

func (td *PlayerTextDraw) Handle() unsafe.Pointer {
	return td.handle
}

func (td *PlayerTextDraw) SetPosition(pos Vector2) {
	C.playerTextDraw_setPosition(td.Handle(), C.float(pos.X), C.float(pos.Y))
}

func (td *PlayerTextDraw) Position() Vector2 {
	pos := C.playerTextDraw_getPosition(td.Handle())

	return Vector2{
		X: float32(pos.x),
		Y: float32(pos.y),
	}
}

func (td *PlayerTextDraw) SetText(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.playerTextDraw_setText(td.Handle(), C.String{
		buf:    ctext,
		length: C.strlen(ctext),
	})
}

func (td *PlayerTextDraw) Text() string {
	text := C.playerTextDraw_getText(td.Handle())

	return C.GoStringN(text.buf, C.int(text.length))
}

func (td *PlayerTextDraw) SetLetterSize(size Vector2) {
	C.playerTextDraw_setLetterSize(td.Handle(), C.float(size.X), C.float(size.Y))
}

func (td *PlayerTextDraw) LetterSize() Vector2 {
	size := C.playerTextDraw_getLetterSize(td.Handle())

	return Vector2{
		X: float32(size.x),
		Y: float32(size.y),
	}
}

func (td *PlayerTextDraw) SetTextSize(size Vector2) {
	C.playerTextDraw_setTextSize(td.Handle(), C.float(size.X), C.float(size.Y))
}

func (td *PlayerTextDraw) TextSize() Vector2 {
	size := C.playerTextDraw_getTextSize(td.Handle())

	return Vector2{
		X: float32(size.x),
		Y: float32(size.y),
	}
}

func (td *PlayerTextDraw) SetAlignment(alignment int) {
	C.playerTextDraw_setAlignment(td.Handle(), C.int(alignment))
}

func (td *PlayerTextDraw) Alignment() int {
	return int(C.playerTextDraw_getAlignment(td.Handle()))
}

func (td *PlayerTextDraw) SetColor(color int) {
	C.playerTextDraw_setColour(td.Handle(), C.uint(color))
}

func (td *PlayerTextDraw) Color() int {
	return int(C.playerTextDraw_getLetterColour(td.Handle()))
}

func (td *PlayerTextDraw) EnableBox() {
	C.playerTextDraw_useBox(td.Handle(), C.int(1))
}

func (td *PlayerTextDraw) DisableBox() {
	C.playerTextDraw_useBox(td.Handle(), C.int(0))
}

func (td *PlayerTextDraw) IsBoxEnabled() bool {
	return C.playerTextDraw_hasBox(td.Handle()) != 0
}

func (td *PlayerTextDraw) SetBoxColor(color int) {
	C.playerTextDraw_setBoxColour(td.Handle(), C.uint(color))
}

func (td *PlayerTextDraw) BoxColor() int {
	return int(C.playerTextDraw_getBoxColour(td.Handle()))
}

func (td *PlayerTextDraw) SetShadow(shadow int) {
	C.playerTextDraw_setShadow(td.Handle(), C.int(shadow))
}

func (td *PlayerTextDraw) Shadow() int {
	return int(C.playerTextDraw_getShadow(td.Handle()))
}

func (td *PlayerTextDraw) SetOutline(outline int) {
	C.playerTextDraw_setOutline(td.Handle(), C.int(outline))
}

func (td *PlayerTextDraw) Outline() int {
	return int(C.playerTextDraw_getOutline(td.Handle()))
}

func (td *PlayerTextDraw) SetBackgroundColor(color int) {
	C.playerTextDraw_setBackgroundColour(td.Handle(), C.uint(color))
}

func (td *PlayerTextDraw) BackgroundColor() int {
	return int(C.playerTextDraw_getBackgroundColour(td.Handle()))
}

func (td *PlayerTextDraw) SetStyle(style int) {
	C.playerTextDraw_setStyle(td.Handle(), C.int(style))
}

func (td *PlayerTextDraw) Style() int {
	return int(C.playerTextDraw_getStyle(td.Handle()))
}

func (td *PlayerTextDraw) EnableProportionality() {
	C.playerTextDraw_setProportional(td.Handle(), C.int(1))
}

func (td *PlayerTextDraw) DisableProportionality() {
	C.playerTextDraw_setProportional(td.Handle(), C.int(0))
}

func (td *PlayerTextDraw) IsProportional() bool {
	return C.playerTextDraw_isProportional(td.Handle()) != 0
}

func (td *PlayerTextDraw) EnableSelection() {
	C.playerTextDraw_setSelectable(td.Handle(), C.int(1))
}

func (td *PlayerTextDraw) DisableSelection() {
	C.playerTextDraw_setSelectable(td.Handle(), C.int(0))
}

func (td *PlayerTextDraw) IsSelectable() bool {
	return C.playerTextDraw_isSelectable(td.Handle()) != 0
}

func (td *PlayerTextDraw) SetPreviewModel(model int) {
	C.playerTextDraw_setPreviewModel(td.Handle(), C.int(model))
}

func (td *PlayerTextDraw) PreviewModel() int {
	return int(C.playerTextDraw_getPreviewModel(td.Handle()))
}

func (td *PlayerTextDraw) SetPreviewRotation(rot Vector3) {
	C.playerTextDraw_setPreviewRotation(td.Handle(), C.float(rot.X), C.float(rot.Y), C.float(rot.Z))
}

func (td *PlayerTextDraw) PreviewRotation() Vector2 {
	rot := C.playerTextDraw_getPreviewRotation(td.Handle())

	return Vector2{
		X: float32(rot.x),
		Y: float32(rot.y),
	}
}

func (td *PlayerTextDraw) SetPreviewVehicleColor(color VehicleColor) {
	C.playerTextDraw_setPreviewVehicleColour(td.Handle(), C.int(color.Primary), C.int(color.Secondary))
}

func (td *PlayerTextDraw) PreviewVehicleColor() VehicleColor {
	color := C.playerTextDraw_getPreviewVehicleColour(td.Handle())

	return VehicleColor{
		Primary:   int(color.primary),
		Secondary: int(color.secondary),
	}
}

func (td *PlayerTextDraw) SetPreviewZoom(zoom float32) {
	C.playerTextDraw_setPreviewZoom(td.Handle(), C.float(zoom))
}

func (td *PlayerTextDraw) PreviewVehicleZoom() float32 {
	return float32(C.playerTextDraw_getPreviewZoom(td.Handle()))
}

func (td *PlayerTextDraw) Show() {
	C.playerTextDraw_show(td.Handle())
}

func (td *PlayerTextDraw) Hide() {
	C.playerTextDraw_hide(td.Handle())
}

func (td *PlayerTextDraw) IsShown() bool {
	return C.playerTextDraw_isShown(td.Handle()) != 0
}
