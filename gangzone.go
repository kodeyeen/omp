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

}

func (gz *GangZone) Flash(color int) {

}

func (gz *GangZone) FlashForPlayer(player *Player, color int) {

}

func (gz *GangZone) ColorForPlayer(player *Player) int {
	return 0
}

func (gz *GangZone) FlashColorForPlayer(player *Player) int {
	return 0
}

func (gz *GangZone) Position() *GangZonePosition {
	return nil
}
