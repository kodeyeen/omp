package main

// #include "component.h"
import "C"
import (
	"unsafe"
)

type PickupType uint8

type Pickup struct {
	handle       unsafe.Pointer
	playerHandle unsafe.Pointer
}

func NewPickup(modelID int, _type PickupType, position Vector3, virtualWorld uint) *Pickup {
	handle := C.pickup_create(C.int(modelID), C.uchar(_type), C.float(position.X), C.float(position.Y), C.float(position.Z), C.uint(virtualWorld), C.int(0), nil)

	pick := &Pickup{
		handle: handle,
	}

	return pick
}

func NewPlayerPickup(player *Player, modelID int, _type PickupType, position Vector3, virtualWorld uint) *Pickup {
	handle := C.pickup_create(C.int(modelID), C.uchar(_type), C.float(position.X), C.float(position.Y), C.float(position.Z), C.uint(virtualWorld), C.int(0), player.handle)

	pick := &Pickup{
		handle:       handle,
		playerHandle: player.handle,
	}

	return pick
}
