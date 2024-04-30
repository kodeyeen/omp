package gomp

// #include "include/textlabel.h"
// #include "include/player.h"
// #include "include/vehicle.h"
import "C"
import (
	"errors"
	"unsafe"
)

type TextLabelAttachmentTarget interface {
	Vehicle | Player
}

type TextLabelAttachmentData[T TextLabelAttachmentTarget] struct {
	Target T
}

type TextLabel struct {
	handle unsafe.Pointer
}

func NewTextLabel(text string, clr Color, pos Vector3, drawDist float32, vw int, los bool) (*TextLabel, error) {
	cText := newCString(text)
	defer freeCString(cText)

	cTl := C.textLabel_create(cText, C.uint(clr), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(drawDist), C.int(vw), newCUchar(los))
	if cTl == nil {
		return nil, errors.New("text label limit reached")
	}

	return &TextLabel{handle: cTl}, nil
}

func FreeTextLabel(tl *TextLabel) {
	C.textLabel_release(tl.handle)
}

func TextLabelAttachedData[T TextLabelAttachmentTarget](tl *TextLabel) (TextLabelAttachmentData[T], error) {
	data := C.textLabel_getAttachmentData(tl.handle)

	t := any(new(T))

	var result TextLabelAttachmentData[T]
	var target any

	switch t.(type) {
	case *Vehicle:
		veh := C.vehicle_getByID(data.vehicleID)
		if veh == nil {
			return result, errors.New("text label is not attached to a vehicle")
		}

		target = Vehicle{handle: veh}
	case *Player:
		plr := C.player_getByID(data.playerID)
		if plr == nil {
			return result, errors.New("text label is not attached to a player")
		}

		target = Player{handle: plr}
	}

	result = TextLabelAttachmentData[T]{
		Target: target.(T),
	}

	return result, nil
}

func (tl *TextLabel) SetText(text string) {
	cText := newCString(text)
	defer freeCString(cText)

	C.textLabel_setText(tl.handle, cText)
}

func (tl *TextLabel) Text() string {
	cTextStr := C.textLabel_getText(tl.handle)

	return C.GoStringN(cTextStr.buf, C.int(cTextStr.length))
}

func (tl *TextLabel) SetColor(clr Color) {
	C.textLabel_setColour(tl.handle, C.uint(clr))
}

func (tl *TextLabel) Color() Color {
	return Color(C.textLabel_getColour(tl.handle))
}

func (tl *TextLabel) SetDrawDistance(drawDist float32) {
	C.textLabel_setDrawDistance(tl.handle, C.float(drawDist))
}

func (tl *TextLabel) DrawDistance() float32 {
	return float32(C.textLabel_getDrawDistance(tl.handle))
}

func (tl *TextLabel) AttachToPlayer(plr *Player, offset Vector3) {
	C.textLabel_attachToPlayer(tl.handle, plr.handle, C.float(offset.X), C.float(offset.Y), C.float(offset.Z))
}

func (tl *TextLabel) AttachToVehicle(veh *Vehicle, offset Vector3) {
	C.textLabel_attachToVehicle(tl.handle, veh.handle, C.float(offset.X), C.float(offset.Y), C.float(offset.Z))
}

func (tl *TextLabel) DetachFromPlayer(pos Vector3) {
	C.textLabel_detachFromPlayer(tl.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (tl *TextLabel) DetachFromVehicle(pos Vector3) {
	C.textLabel_detachFromVehicle(tl.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (tl *TextLabel) EnableLOSTest() {
	C.textLabel_setTestLOS(tl.handle, 1)
}

func (tl *TextLabel) DisableLOSTest() {
	C.textLabel_setTestLOS(tl.handle, 0)
}

func (tl *TextLabel) IsLOSTestEnabled() bool {
	return C.textLabel_getTestLOS(tl.handle) != 0
}

func (tl *TextLabel) IsStreamedInFor(plr *Player) bool {
	return C.textLabel_isStreamedInForPlayer(tl.handle, plr.handle) != 0
}

func (tl *TextLabel) SetPosition(pos Vector3) {
	C.textLabel_setPosition(tl.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (tl *TextLabel) Position() Vector3 {
	pos := C.textLabel_getPosition(tl.handle)

	return Vector3{
		X: float32(pos.x),
		Y: float32(pos.y),
		Z: float32(pos.z),
	}
}

func (tl *TextLabel) SetVirtualWorld(vw int) {
	C.textLabel_setVirtualWorld(tl.handle, C.int(vw))
}

func (tl *TextLabel) VirtualWorld() int {
	return int(C.textLabel_getVirtualWorld(tl.handle))
}
