package gomp

// #include "pickup.h"
import "C"
import (
	"unsafe"
)

type PickupType uint8

type Pickup struct {
	handle       unsafe.Pointer
	playerHandle unsafe.Pointer
}

func CreatePickup(modelID int, _type PickupType, x, y, z float32, virtualWorld uint) *Pickup {
	handle := C.pickup_create(C.int(modelID), C.uchar(_type), C.float(x), C.float(y), C.float(z), C.uint(virtualWorld), C.int(0), nil)

	pick := &Pickup{
		handle: handle,
	}

	return pick
}

func CreatePlayerPickup(player *Player, modelID int, _type PickupType, position Vector3, virtualWorld uint) *Pickup {
	handle := C.pickup_create(C.int(modelID), C.uchar(_type), C.float(position.X), C.float(position.Y), C.float(position.Z), C.uint(virtualWorld), C.int(0), player.handle)

	pick := &Pickup{
		handle:       handle,
		playerHandle: player.handle,
	}

	return pick
}

func DestroyPickup(pickup *Pickup) {

}

func (p *Pickup) ModelID() int {
	return 0
}

func (p *Pickup) SetModelID(modelID int) {

}

func (p *Pickup) Position() *Vector3 {
	return nil
}

func (p *Pickup) SetPosition(position *Vector3) {

}

func (p *Pickup) Type() int {
	return 0
}

func (p *Pickup) SetType(_type int) {

}

func (p *Pickup) VirtualWorld() int {
	return 0
}

func (p *Pickup) SetVirtualWorld(virtualWorld int) {

}

func (p *Pickup) IsStreamedIn(player *Player) bool {
	return false
}

func (p *Pickup) HideForPlayer(player *Player) {

}

func (p *Pickup) ShowForPlayer(player *Player) {

}

func (p *Pickup) IsHiddenForPlayer(player *Player) bool {
	return false
}
