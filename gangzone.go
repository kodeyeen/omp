package gomp

import "unsafe"

type GangZonePosition struct {
	Min Vector2
	Max Vector2
}

type GangZone struct {
	handle unsafe.Pointer
}

func NewGangZone() *GangZone {
	return &GangZone{}
}

func NewPlayerGangZone() *GangZone {
	return &GangZone{}
}

func DestroyGangZone(gangZone *GangZone) {
	panic("not implemented")
}

func (gz *GangZone) Flash(color uint) {
	panic("not implemented")
}

func (gz *GangZone) FlashForPlayer(player *Player, color uint) {
	panic("not implemented")
}

func (gz *GangZone) ColorForPlayer(player *Player) int {
	panic("not implemented")
}

func (gz *GangZone) FlashColorForPlayer(player *Player) int {
	panic("not implemented")
}

func (gz *GangZone) Position() *GangZonePosition {
	panic("not implemented")
}
