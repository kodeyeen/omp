package main

import "unsafe"

type Pickup struct {
	handle       unsafe.Pointer
	playerHandle unsafe.Pointer
}

func NewPickup() *Pickup {
	return &Pickup{}
}

func NewPlayerPickup(player *Player) *Pickup {
	return &Pickup{}
}
