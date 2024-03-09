package main

import (
	"unsafe"
)

type Player struct {
	handle unsafe.Pointer
	comp   *component
}

func newPlayer(handle unsafe.Pointer, comp *component) *Player {
	return &Player{
		handle: handle,
		comp:   comp,
	}
}

func (p *Player) ID() int {
	// return int(C.player_getID(p.handle))
	return 0
}

func (p *Player) Name() string {
	return p.comp.player_getName(p.handle)
}

func (p *Player) SendMessage(color int, msg string) {
	p.comp.player_sendClientMessage(p.handle, color, msg)
}
