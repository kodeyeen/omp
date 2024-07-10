package omp

// #include "include/wrappers.h"
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
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	var cID C.int

	cTl := C.TextLabel_Create(cText, C.uint(clr), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(drawDist), C.int(vw), C.bool(los), &cID)
	if cTl == nil {
		return nil, errors.New("text label limit was reached")
	}

	return &TextLabel{handle: cTl}, nil
}

func FreeTextLabel(tl *TextLabel) {
	C.TextLabel_Destroy(tl.handle)
}

func (tl *TextLabel) ID() int {
	return int(C.TextLabel_GetID(tl.handle))
}

func TextLabelAttachedData[T TextLabelAttachmentTarget](tl *TextLabel) (TextLabelAttachmentData[T], error) {
	var cPlrID, cVehID C.int

	C.TextLabel_GetAttachedData(tl.handle, &cPlrID, &cVehID)

	var t T
	var result TextLabelAttachmentData[T]
	var target any

	switch any(t).(type) {
	case Vehicle:
		veh := C.Vehicle_FromID(cVehID)
		if veh == nil {
			return result, errors.New("text label is not attached to a vehicle")
		}

		target = Vehicle{handle: veh}
	case Player:
		plr := C.Player_FromID(cPlrID)
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

// func (tl *TextLabel) SetText(text string) {
// 	cText := newCString(text)
// 	defer freeCString(cText)

// 	C.textLabel_setText(tl.handle, cText)
// }

func (tl *TextLabel) Text() string {
	var cText C.struct_CAPIStringView

	C.TextLabel_GetText(tl.handle, &cText)

	return C.GoStringN(cText.data, C.int(cText.len))
}

// func (tl *TextLabel) SetColor(clr Color) {
// 	C.textLabel_setColour(tl.handle, C.uint(clr))
// }

func (tl *TextLabel) Color() Color {
	return Color(C.TextLabel_GetColor(tl.handle))
}

func (tl *TextLabel) SetDrawDistance(drawDist float32) {
	C.TextLabel_SetDrawDistance(tl.handle, C.float(drawDist))
}

func (tl *TextLabel) DrawDistance() float32 {
	return float32(C.TextLabel_GetDrawDistance(tl.handle))
}

func (tl *TextLabel) AttachToPlayer(plr *Player, offset Vector3) {
	C.TextLabel_AttachToPlayer(tl.handle, plr.handle, C.float(offset.X), C.float(offset.Y), C.float(offset.Z))
}

func (tl *TextLabel) AttachToVehicle(veh *Vehicle, offset Vector3) {
	C.TextLabel_AttachToVehicle(tl.handle, veh.handle, C.float(offset.X), C.float(offset.Y), C.float(offset.Z))
}

// func (tl *TextLabel) DetachFromPlayer(pos Vector3) {
// 	C.textLabel_detachFromPlayer(tl.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
// }

// func (tl *TextLabel) DetachFromVehicle(pos Vector3) {
// 	C.textLabel_detachFromVehicle(tl.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
// }

func (tl *TextLabel) EnableLOSTest() {
	C.TextLabel_SetLOS(tl.handle, true)
}

func (tl *TextLabel) DisableLOSTest() {
	C.TextLabel_SetLOS(tl.handle, false)
}

func (tl *TextLabel) IsLOSTestEnabled() bool {
	return bool(C.TextLabel_GetLOS(tl.handle))
}

func (tl *TextLabel) IsStreamedInFor(plr *Player) bool {
	return bool(C.TextLabel_IsStreamedIn(plr.handle, tl.handle))
}

// func (tl *TextLabel) SetPosition(pos Vector3) {
// 	C.textLabel_setPosition(tl.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
// }

func (tl *TextLabel) Position() Vector3 {
	var cPosX, cPosY, cPosZ C.float

	C.TextLabel_GetPos(tl.handle, &cPosX, &cPosY, &cPosZ)

	return Vector3{
		X: float32(cPosX),
		Y: float32(cPosY),
		Z: float32(cPosZ),
	}
}

func (tl *TextLabel) SetVirtualWorld(vw int) {
	C.TextLabel_SetVirtualWorld(tl.handle, C.int(vw))
}

func (tl *TextLabel) VirtualWorld() int {
	return int(C.TextLabel_GetVirtualWorld(tl.handle))
}

type PlayerTextLabel struct {
	handle unsafe.Pointer
	player *Player
}

func NewPlayerTextLabel(plr *Player, text string, clr Color, pos Vector3, drawDist float32, vw int, los bool) (*PlayerTextLabel, error) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	var cID C.int

	cTl := C.PlayerTextLabel_Create(plr.handle, cText, C.uint(clr), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(drawDist), nil, nil, C.bool(los), &cID)
	if cTl == nil {
		return nil, errors.New("player text label limit was reached")
	}

	tl := &PlayerTextLabel{
		handle: cTl,
		player: plr,
	}

	return tl, nil
}

func FreePlayerTextLabel(tl *PlayerTextLabel) {
	C.PlayerTextLabel_Destroy(tl.player.handle, tl.handle)
}

func (tl *PlayerTextLabel) Player() *Player {
	return tl.player
}

func (tl *PlayerTextLabel) ID() int {
	return int(C.PlayerTextLabel_GetID(tl.player.handle, tl.handle))
}

func PlayerTextLabelAttachedData[T TextLabelAttachmentTarget](tl *PlayerTextLabel) (TextLabelAttachmentData[T], error) {
	var cPlrID, cVehID C.int

	C.PlayerTextLabel_GetAttachedData(tl.player.handle, tl.handle, &cPlrID, &cVehID)

	var t T
	var result TextLabelAttachmentData[T]
	var target any

	switch any(t).(type) {
	case Vehicle:
		veh := C.Vehicle_FromID(cVehID)
		if veh == nil {
			return result, errors.New("text label is not attached to a vehicle")
		}

		target = Vehicle{handle: veh}
	case Player:
		plr := C.Player_FromID(cPlrID)
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

// func (tl *PlayerTextLabel) SetText(text string) {
// 	cText := newCString(text)
// 	defer freeCString(cText)

// 	C.playerTextLabel_setText(tl.handle, cText)
// }

func (tl *PlayerTextLabel) Text() string {
	var cText C.struct_CAPIStringView

	C.PlayerTextLabel_GetText(tl.player.handle, tl.handle, &cText)

	return C.GoStringN(cText.data, C.int(cText.len))
}

// func (tl *PlayerTextLabel) SetColor(clr Color) {
// 	C.playerTextLabel_setColour(tl.handle, C.uint(clr))
// }

func (tl *PlayerTextLabel) Color() Color {
	var cColor C.uint
	C.PlayerTextLabel_GetColor(tl.player.handle, tl.handle, &cColor)

	return Color(cColor)
}

func (tl *PlayerTextLabel) SetDrawDistance(drawDist float32) {
	C.PlayerTextLabel_SetDrawDistance(tl.player.handle, tl.handle, C.float(drawDist))
}

func (tl *PlayerTextLabel) DrawDistance() float32 {
	return float32(C.PlayerTextLabel_GetDrawDistance(tl.player.handle, tl.handle))
}

// func (tl *PlayerTextLabel) AttachToPlayer(plr *Player, offset Vector3) {
// 	C.playerTextLabel_attachToPlayer(tl.handle, plr.handle, C.float(offset.X), C.float(offset.Y), C.float(offset.Z))
// }

// func (tl *PlayerTextLabel) AttachToVehicle(veh *Vehicle, offset Vector3) {
// 	C.playerTextLabel_attachToVehicle(tl.handle, veh.handle, C.float(offset.X), C.float(offset.Y), C.float(offset.Z))
// }

// func (tl *PlayerTextLabel) DetachFromPlayer(pos Vector3) {
// 	C.playerTextLabel_detachFromPlayer(tl.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
// }

// func (tl *PlayerTextLabel) DetachFromVehicle(pos Vector3) {
// 	C.playerTextLabel_detachFromVehicle(tl.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
// }

func (tl *PlayerTextLabel) EnableLOSTest() {
	C.PlayerTextLabel_SetLOS(tl.player.handle, tl.handle, true)
}

func (tl *PlayerTextLabel) DisableLOSTest() {
	C.PlayerTextLabel_SetLOS(tl.player.handle, tl.handle, false)
}

func (tl *PlayerTextLabel) IsLOSTestEnabled() bool {
	return bool(C.PlayerTextLabel_GetLOS(tl.player.handle, tl.handle))
}

// func (tl *PlayerTextLabel) SetPosition(pos Vector3) {
// 	C.playerTextLabel_setPosition(tl.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
// }

func (tl *PlayerTextLabel) Position() Vector3 {
	var cPosX, cPosY, cPosZ C.float

	C.PlayerTextLabel_GetPos(tl.player.handle, tl.handle, &cPosX, &cPosY, &cPosZ)

	return Vector3{
		X: float32(cPosX),
		Y: float32(cPosY),
		Z: float32(cPosZ),
	}
}

// func (tl *PlayerTextLabel) SetVirtualWorld(vw int) {
// 	C.playerTextLabel_setVirtualWorld(tl.handle, C.int(vw))
// }

func (tl *PlayerTextLabel) VirtualWorld() int {
	return int(C.PlayerTextLabel_GetVirtualWorld(tl.player.handle))
}
