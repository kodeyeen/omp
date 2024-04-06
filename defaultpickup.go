package gomp

// #include "include/pickup.h"
import "C"
import (
	"unsafe"
)

type PickupType uint8

type DefaultPickup struct {
	handle unsafe.Pointer
}

func NewPickup(modelID int, _type PickupType, pos Vector3, virtualWorld uint) *DefaultPickup {
	handle := C.pickup_create(C.int(modelID), C.uchar(_type), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.uint(virtualWorld), C.int(0), nil)

	pick := &DefaultPickup{
		handle: handle,
	}

	return pick
}

func FreePickup(pickup *DefaultPickup) {
	panic("not implemented")
}

func (p *DefaultPickup) ModelID() int {
	panic("not implemented")
}

func (p *DefaultPickup) SetModelID(modelID int) {
	panic("not implemented")
}

func (p *DefaultPickup) Position() *Vector3 {
	panic("not implemented")
}

func (p *DefaultPickup) SetPosition(position *Vector3) {
	panic("not implemented")
}

func (p *DefaultPickup) Type() int {
	panic("not implemented")
}

func (p *DefaultPickup) SetType(_type int) {
	panic("not implemented")
}

func (p *DefaultPickup) VirtualWorld() int {
	panic("not implemented")
}

func (p *DefaultPickup) SetVirtualWorld(virtualWorld int) {
	panic("not implemented")
}

func (p *DefaultPickup) IsStreamedIn(player *Player) bool {
	panic("not implemented")
}

func (p *DefaultPickup) HideForPlayer(player *Player) {
	panic("not implemented")
}

func (p *DefaultPickup) ShowForPlayer(player *Player) {
	panic("not implemented")
}

func (p *DefaultPickup) IsHiddenForPlayer(player *Player) bool {
	panic("not implemented")
}
