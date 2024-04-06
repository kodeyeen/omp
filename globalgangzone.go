package gomp

import "unsafe"

type GlobalGangZone struct {
	handle unsafe.Pointer
}

func NewGlobalGangZone() *GlobalGangZone {
	return &GlobalGangZone{}
}

func DestroyGangZone(gangZone *GlobalGangZone) {
	panic("not implemented")
}

func (gz *GlobalGangZone) Flash(color int) {
	panic("not implemented")
}

func (gz *GlobalGangZone) FlashForPlayer(player Player, color int) {
	panic("not implemented")
}

func (gz *GlobalGangZone) ColorForPlayer(player Player) int {
	panic("not implemented")
}

func (gz *GlobalGangZone) FlashColorForPlayer(player Player) int {
	panic("not implemented")
}

func (gz *GlobalGangZone) Position() *GangZonePosition {
	panic("not implemented")
}
