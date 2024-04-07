package gomp

// #include "include/pickup.h"
import "C"
import (
	"unsafe"
)

type PickupType uint8

type Pickup struct {
	handle unsafe.Pointer
}

func NewPickup(modelID int, _type PickupType, pos Vector3, virtualWorld uint) *Pickup {
	handle := C.pickup_create(C.int(modelID), C.uchar(_type), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.uint(virtualWorld), C.int(0), nil)

	pick := &Pickup{
		handle: handle,
	}

	return pick
}

func FreePickup(pickup *Pickup) {
	panic("not implemented")
}

func (p *Pickup) ModelID() int {
	panic("not implemented")
}

func (p *Pickup) SetModelID(modelID int) {
	panic("not implemented")
}

func (p *Pickup) Position() *Vector3 {
	panic("not implemented")
}

func (p *Pickup) SetPosition(position *Vector3) {
	panic("not implemented")
}

func (p *Pickup) Type() int {
	panic("not implemented")
}

func (p *Pickup) SetType(_type int) {
	panic("not implemented")
}

func (p *Pickup) VirtualWorld() int {
	panic("not implemented")
}

func (p *Pickup) SetVirtualWorld(virtualWorld int) {
	panic("not implemented")
}

func (p *Pickup) IsStreamedIn(player *Player) bool {
	panic("not implemented")
}

func (p *Pickup) HideForPlayer(player *Player) {
	panic("not implemented")
}

func (p *Pickup) ShowForPlayer(player *Player) {
	panic("not implemented")
}

func (p *Pickup) IsHiddenForPlayer(player *Player) bool {
	panic("not implemented")
}
