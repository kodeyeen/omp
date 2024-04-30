package gomp

// #include "include/pickup.h"
import "C"
import (
	"errors"
	"unsafe"
)

type PickupType uint8

type Pickup struct {
	handle unsafe.Pointer
}

func NewPickup(modelID int, _type PickupType, virtualWorld int, pos Vector3) (*Pickup, error) {
	cPickup := C.pickup_create(C.int(modelID), C.uchar(_type), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.uint(virtualWorld), 0)
	if cPickup == nil {
		return nil, errors.New("pickup limit reached")
	}

	return &Pickup{handle: cPickup}, nil
}

func FreePickup(pickup *Pickup) {
	C.pickup_release(pickup.handle)
}

func (p *Pickup) ID() int {
	return int(C.pickup_getID(p.handle))
}

func (p *Pickup) SetType(_type PickupType) {
	C.pickup_setType(p.handle, C.uchar(_type))
}

func (p *Pickup) Type() PickupType {
	return PickupType(C.pickup_getType(p.handle))
}

func (p *Pickup) SetModel(model int) {
	C.pickup_setModel(p.handle, C.int(model))
}

func (p *Pickup) Model() int {
	return int(C.pickup_getModel(p.handle))
}

func (p *Pickup) IsStreamedInFor(plr *Player) bool {
	return C.pickup_isStreamedInForPlayer(p.handle, plr.handle) != 0
}

func (p *Pickup) ShowFor(plr *Player) {
	C.pickup_setPickupHiddenForPlayer(p.handle, plr.handle, 0)
}

func (p *Pickup) HideFor(plr *Player) {
	C.pickup_setPickupHiddenForPlayer(p.handle, plr.handle, 1)
}

func (p *Pickup) IsHiddenFor(plr *Player) bool {
	return C.pickup_isPickupHiddenForPlayer(p.handle, plr.handle) != 0
}

func (p *Pickup) SetPosition(pos Vector3) {
	C.pickup_setPosition(p.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (p *Pickup) Position() Vector3 {
	pos := C.pickup_getPosition(p.handle)

	return Vector3{
		X: float32(pos.x),
		Y: float32(pos.y),
		Z: float32(pos.z),
	}
}

func (p *Pickup) SetVirtualWorld(vw int) {
	C.pickup_setVirtualWorld(p.handle, C.int(vw))
}

func (p *Pickup) VirtualWorld() int {
	return int(C.pickup_getVirtualWorld(p.handle))
}
